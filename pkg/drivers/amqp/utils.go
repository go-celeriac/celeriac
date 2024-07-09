package amqp

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const timeFormat = time.RFC3339

// NewMessageID generates a UUIDv4 as a string
// if generation fails for any reason, the function panics
func NewMessageID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return id.String()
}

func TimeToString(t time.Time) string {
	return t.Format(timeFormat)
}

func StringToTime(v string) (time.Time, error) {
	if v == "" {
		return time.Time{}, fmt.Errorf("time value is empty")
	}
	return time.Parse(timeFormat, v)
}
