// Package isg provides a way to report inbound and outbound requests.
package isg

import (
	"context"
)

// StartedContext is an interface that represents a context that has been started.
type StartedContext interface {
	Finished()
}

var nilStartedContext *NilStartedContext = nil

type NilStartedContext struct {
}

func (*NilStartedContext) Finished() {
	// do nothing
}

// Reporter is an interface that represents an object that can report inbound and outbound requests.
type Reporter interface {
	// Inbound reports an inbound interface processing started.
	// It returns a StartedContext that should be used to report when the processing has finished.
	Inbound(ctx context.Context, service string, component string, interfaceType string, interfaceID string) StartedContext

	// Outbound reports an outbound interface processing.
	// It returns a StartedContext that should be used to report when the processing has finished.
	Outbound(ctx context.Context, service string, component string, interfaceType string, interfaceID string) StartedContext
}
