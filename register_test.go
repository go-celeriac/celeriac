package celeriac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testBroker struct{}

func (tb *testBroker) Close() error {
	return nil
}

func (tb *testBroker) GetQueue(string) (Queue, error) {
	return nil, nil
}

func TestRegisterDriver(t *testing.T) {
	tests := map[string]struct {
		seed              func()
		registeredDrivers int
	}{
		"Registry is empty by default": {
			func() {},
			0,
		},
		"Registry contains one driver when one is registered": {
			func() {
				RegisterDriver("test", func(string) (Broker, error) { return &testBroker{}, nil })
			},
			1,
		},
	}

	for name, test := range tests {
		driverRegistry = make(map[string]BrokerFactory) // clear out the registry with each pass

		test.seed()
		assert.Equal(t, test.registeredDrivers, len(driverRegistry), name)
	}
}
