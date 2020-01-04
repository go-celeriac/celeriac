package celeriac

var (
	driverRegistry = make(map[string]BrokerFactory)
)

// BrokerFactory should create a new Broker with the given connection string
type BrokerFactory func(connection string) (Broker, error)

// RegisterDriver registers a new BrokerFactory to be used by NewBroker, this function
// should be used in an init function in the implementation packages
func RegisterDriver(name string, factory BrokerFactory) {
	driverRegistry[name] = factory
}
