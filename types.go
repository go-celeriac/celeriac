package celeriac

import (
	"encoding/json"
	"time"
)

// Message defines the properties of the payload which transported
// through the queue
type Message struct {
	Body       []byte
	Expiration *time.Time
	MessageID  string
	Timestamp  time.Time
}

type MessageBody struct {
	Args []interface{} `json:"args,omitempty"`
	Task string        `json:"task"`
}

func (mb *MessageBody) ToBytes() ([]byte, error) {
	return json.Marshal(mb)
}

func ParseMessageBody(data []byte) (*MessageBody, error) {
	var mb MessageBody

	if err := json.Unmarshal(data, &mb); err != nil {
		return nil, err
	}

	return &mb, nil
}
