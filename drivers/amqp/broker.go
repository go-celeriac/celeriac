package amqp

import (
	"fmt"

	"github.com/go-celeriac/celeriac"
)

// AMQP is an instance of a celeriac Broker which uses amqp
type AMQP struct{}

// Close closes the connection to the broker
func (a *AMQP) Close() error {
	return nil
}

func newAMQP(uri string) (celeriac.Broker, error) {
	fmt.Println("Loading AMQP driver...")

	return &AMQP{}, nil
}

func init() {
	celeriac.RegisterDriver("amqp", newAMQP)
}
