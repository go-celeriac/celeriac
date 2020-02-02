package amqp

import (
	"log"
	"time"

	"github.com/go-celeriac/celeriac"
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
			var expires *time.Time

			t, err := StringToTime(msg.Expiration)
			if err != nil {
				log.Printf("[ERROR] unable to parse expiration %s due to: %s", msg.Expiration, err.Error())
			} else {
				expires = &t
			}

			output <- celeriac.Message{
				MessageID:  msg.MessageId,
				Expiration: expires,
				Timestamp:  msg.Timestamp,
				Body:       msg.Body,
			}
		}
	}()

	return output, nil
}

func (q *Queue) Publish(body []byte) (string, error) {
	messageID := NewMessageID()

	err := q.channel.Publish(
		"",
		q.q.Name,
		false,
		false,
		mq.Publishing{
			ContentType: "text/plain",
			MessageId:   messageID,
			Body:        body,
		},
	)

	return messageID, err
}
