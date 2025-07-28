package promisg

import (
	"github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
	"time"
)

// startedContext is an interface that represents a context for tracking the start of processing.
// It provides methods for getting the service, component, interface type, interface ID, and start time.
type startedContext interface {
	Service() string
	Component() string
	InterfaceType() isg.InterfaceType
	InterfaceID() string
	StartTime() time.Time
}
