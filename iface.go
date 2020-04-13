package celeriac

// Broker defines the features which all drivers must implement
type Broker interface {
	Close() error
	GetQueue(string) (Queue, error)
	Enqueue(Task, ...interface{}) (string, error)
}

// Queue defines the features of a Queue which all drivers must implement
type Queue interface {
	Consume() (<-chan Message, error)
	Publish([]byte) (string, error)
}

// Task defines the functionality required for a type to be executed via a queue
type Task interface {
	Init() error
	Run(...interface{}) error
	Exit() error
}
