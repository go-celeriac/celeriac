package celeriac

import "time"

// Message defines the properties of the payload which transported
// through the queue
type Message struct {
	Body       []byte
	Expiration *time.Time
	MessageID  string
	Timestamp  time.Time
}
