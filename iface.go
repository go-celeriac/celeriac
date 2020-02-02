package celeriac

// Broker defines the features which all drivers must implement
type Broker interface {
	Close() error
	GetQueue(string) (Queue, error)
}

// Queue defines the features of a Queue which all drivers must implement
type Queue interface {
	Consume() (<-chan Message, error)
	Publish([]byte) (string, error)
}
