package amqp

import (
	"github.com/go-celeriac/celeriac"
	"github.com/google/uuid"
	mq "github.com/streadway/amqp"
)

type Queue struct {
	channel *mq.Channel
	q       *mq.Queue
}

func (q *Queue) Consume() (<-chan celeriac.Message, error) {
	messages, err := q.channel.Consume(
		q.q.Name,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		return nil, err
	}

	output := make(chan celeriac.Message)

	go func() {
		for msg := range messages {
			output <- celeriac.Message{
				MessageID:  msg.MessageId,
				Expiration: msg.Expiration,
				Timestamp:  msg.Timestamp,
				Body:       msg.Body,
			}
		}
	}()

	return output, nil
}

func (q *Queue) Publish(body []byte) error {
	messageID, _ := uuid.NewRandom()

	return q.channel.Publish(
		"",
		q.q.Name,
		false,
		false,
		mq.Publishing{
			ContentType: "text/plain",
			MessageId:   messageID.String(),
			Body:        body,
		},
	)
}
