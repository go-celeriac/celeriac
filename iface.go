package celeriac


type TaskId string
type TaskDefinition string

// Worker is what processes the tasks which come from the message broker
type Worker interface {
	Process(id string, body string)
}

// Broker defines the features which all drivers must implement
type Broker interface {
	Close() error

	Enqueue(task TaskDefinition, queueName string) (TaskId, error)
	Dequeue(worker Worker, queueName string) error
}
