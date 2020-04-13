package celeriac

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	driverRegistry = make(map[string]BrokerFactory)
	taskRegistry   = make(map[string]Task)
)

// BrokerFactory should create a new Broker with the given connection string
type BrokerFactory func(connection string) (Broker, error)

// RegisterDriver registers a new BrokerFactory to be used by NewBroker, this function
// should be used in an init function in the implementation packages
func RegisterDriver(name string, factory BrokerFactory) {
	driverRegistry[name] = factory
}

// RegisterTask registers a Task instance for later retrieval
func RegisterTask(t Task) {
	key := QueueNameForTask(t)
	fmt.Println("Registering", key)

	taskRegistry[key] = t
}

func GetTask(name string) Task {
	t, ok := taskRegistry[name]

	if !ok {
		return nil
	}

	elem := reflect.TypeOf(t).Elem()
	newThingValue := reflect.New(elem)
	return newThingValue.Interface().(Task)
}

func QueueNameForTask(t Task) string {
	value := reflect.Indirect(reflect.ValueOf(t))

	taskType := value.Type()

	parts := strings.Split(taskType.PkgPath(), "/")

	path := parts[len(parts)-1]

	return strings.ToLower(fmt.Sprintf("%s.%s", path, taskType.Name()))
}
