package amqp

import (
	"fmt"

	"github.com/go-celeriac/celeriac"
	mq "github.com/streadway/amqp"
)

// Broker is an instance of a celeriac Broker which uses amqp
type Broker struct {
	channel    *mq.Channel
	connection *mq.Connection

	queues map[string]*Queue
}

// Close closes the connection to the broker
func (a *Broker) Close() error {
	if err := a.channel.Close(); err != nil {
		return err
	}

	return a.connection.Close()
}

func (a *Broker) Enqueue(t celeriac.Task, args ...interface{}) (string, error) {
	key := celeriac.QueueNameForTask(t)

	queue, err := a.GetQueue(key)
	if err != nil {
		return "", err
	}

	data, err := (&celeriac.MessageBody{
		Args: args,
		Task: key,
	}).ToBytes()

	if err != nil {
		return "", err
	}

	return queue.Publish(data)
}

func (a *Broker) GetQueue(name string) (celeriac.Queue, error) {
	if a.channel == nil {
		ch, err := a.connection.Channel()
		if err != nil {
			return nil, err
		}

		a.channel = ch
	}

	var queue *Queue

	if q, exists := a.queues[name]; exists {
		queue = q
	} else {

		q, err := a.channel.QueueDeclare(
			name,
			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)

		if err != nil {
			return nil, err
		}

		queue = &Queue{
			channel: a.channel,
			q:       &q,
		}

		a.queues[name] = queue
	}

	return queue, nil
}

func newBroker(uri string) (celeriac.Broker, error) {
	fmt.Println("Loading AMQP driver...")

	connection, err := mq.Dial(uri)
	if err != nil {
		return nil, err
	}

	return &Broker{
		connection: connection,
		queues:     make(map[string]*Queue),
	}, nil
}

func init() {
	celeriac.RegisterDriver("amqp", newBroker)
}
