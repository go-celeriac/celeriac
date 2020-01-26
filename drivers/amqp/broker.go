package amqp

import (
	"log"

	"github.com/google/uuid"
	"github.com/streadway/amqp"

	"github.com/go-celeriac/celeriac"
)

// AMQP is an instance of a celeriac Broker which uses amqp
type AMQP struct{
	Connection amqp.Connection
}

// Close closes the connection to the broker
func (a *AMQP) Close() error {
	err := a.Connection.Close()

	return err
}

func (a *AMQP) declareQueue(queueName string) (amqp.Queue, error){
	ch, err := a.Connection.Channel()
	defer ch.Close()

	if err != nil {
		log.Fatal("Unable to open channel")
	}

	queue, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	return queue, err
}

func (a *AMQP) Enqueue(definition celeriac.TaskDefinition, queueName string) (celeriac.TaskId, error) {
	ch, err := a.Connection.Channel()

	if err != nil {
		log.Fatal("Unable to open channel")
	}

	defer ch.Close()

	queue, err := a.declareQueue(queueName)
	id, err := uuid.NewRandom()
	messageId := id.String()

	if err != nil {
		log.Fatal("Unable to generate random ID")
	}

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(definition),
			MessageId:    messageId,
		})

	log.Println("Published message", id.String())
	return celeriac.TaskId(messageId), err
}

func (a *AMQP) Dequeue(worker celeriac.Worker, queueName string) error {
	ch, err := a.Connection.Channel()

	if err != nil {
		log.Fatal("Unable to connect to channel")
	}

	defer ch.Close()

	queue, err := a.declareQueue(queueName)

	if err != nil {
		log.Fatal("Unable to declare queue", err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Unable to consume from queue", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			worker.Process(d.MessageId, string(d.Body))
		}
	}()

	log.Printf("Waiting to receive task. Press CTRL-C to exit.")
	<-forever
	return err
}

func newAMQP(uri string) (celeriac.Broker, error) {
	log.Println("Loading AMQP driver...")
	log.Println("Connecting to Broker!")
	connection, err := amqp.Dial(uri)

	if err != nil {
		log.Fatal("Unable to connect to AMQP broker")
	}

	return &AMQP{Connection:*connection}, nil
}

func init() {
	celeriac.RegisterDriver("amqp", newAMQP)
}
