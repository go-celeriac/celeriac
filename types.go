package celeriac

import "time"

type Message struct {
	Expiration string    // implementation use - message expiration spec
	MessageID  string    // application use - message identifier
	Timestamp  time.Time // application use - message timestamp
	Type       string    // application use - message type name

	DeliveryTag uint64
	Redelivered bool
	Exchange    string // basic.publish exchange
	RoutingKey  string // basic.publish routing key

	Body []byte
}
