package celeriac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBroker(t *testing.T) {
	tests := map[string]struct {
		connectionString string
		broker           Broker
		expectedErr      error
	}{
		"Returns error when connection string is malformed": {
			"%^&*",
			nil,
			fmt.Errorf("uri is not valid"),
		},
		"Returns error when connection string does not have a scheme": {
			"foobar",
			nil,
			fmt.Errorf("uri is not valid - missing scheme"),
		},
		"Returns error when driver cannot be found for connection scheme": {
			"foobar://localhost",
			nil,
			fmt.Errorf("driver for foobar has not been registered, are you missing an import?"),
		},
		"Returns broker instance when all is well": {
			"test://localhost",
			&testBroker{},
			nil,
		},
	}

	driverRegistry = make(map[string]BrokerFactory) // clear out the registry with each pass

	for name, test := range tests {
		RegisterDriver("test", func(string) (Broker, error) { return &testBroker{}, nil })

		broker, err := NewBroker(test.connectionString)

		assert.Equal(t, test.expectedErr, err, name)
		assert.Equal(t, test.broker, broker, name)

		driverRegistry = make(map[string]BrokerFactory) // clear out the registry with each pass
	}
}
