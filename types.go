package celeriac

import "time"

// Message defines the properties of the payload which transported
// through the queue
type Message struct {
	Expiration string
	MessageID  string
	Timestamp  time.Time
	Body       []byte
}
