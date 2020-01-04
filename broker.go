package celeriac

import (
	"fmt"
	"net/url"
)

// NewBroker creates a new Broker based on the given URI
func NewBroker(uri string) (Broker, error) {
	url, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("uri is not valid")
	}

	if url.Scheme == "" {
		return nil, fmt.Errorf("uri is not valid - missing scheme")
	}

	factory, ok := driverRegistry[url.Scheme]
	if !ok {
		return nil, fmt.Errorf("driver for %s has not been registered, are you missing an import?", url.Scheme)
	}

	return factory(uri)
}
