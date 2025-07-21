package promisg

import "time"

// startedContext is an interface that represents a context for tracking the start of processing.
// It provides methods for getting the service, component, interface type, interface ID, and start time.
type startedContext interface {
	Service() string
	Component() string
	InterfaceType() string
	InterfaceID() string
	StartTime() time.Time
}
