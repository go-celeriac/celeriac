package celeriac

// Broker defines the features which all drivers must implement
type Broker interface {
	Close() error
}
