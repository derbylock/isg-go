package isg

import (
	"context"
	"time"
)

// inboundContextKey is a type that represents a key for storing InboundContext in a context.
type inboundContextKey int

const (
	// inboundContextKeyDefault is the default key for storing InboundContext in a context.
	inboundContextKeyDefault inboundContextKey = iota + 1
)

// ContextKeeper is an interface that abstracts keeping and extraction of InboundContext.
type ContextKeeper interface {
	// KeepInboundContext keeps the given InboundContext in the given context.
	KeepInboundContext(ctx context.Context, c *InboundContext) context.Context

	// ExtractInboundContext extracts the InboundContext from the given context.
	ExtractInboundContext(ctx context.Context) *InboundContext
}

// ContextValueContextKeeper keeps InboundContext as a value in context.
type ContextValueContextKeeper struct {
}

// NewContextValueContextKeeper creates ContextKeeper that keeps InboundContext as a value in context.
func NewContextValueContextKeeper() *ContextValueContextKeeper {
	return &ContextValueContextKeeper{}
}

// KeepInboundContext keeps the given InboundContext in the given context.
func (k *ContextValueContextKeeper) KeepInboundContext(ctx context.Context, c *InboundContext) context.Context {
	return context.WithValue(ctx, inboundContextKeyDefault, c)
}

// ExtractInboundContext extracts the InboundContext from the given context.
func (k *ContextValueContextKeeper) ExtractInboundContext(ctx context.Context) *InboundContext {
	v := ctx.Value(inboundContextKeyDefault)
	if v == nil {
		return nil
	}

	return v.(*InboundContext)
}

// InboundContext is a struct that represents the context of an inbound request.
type InboundContext struct {
	service       string
	component     string
	interfaceType string
	interfaceID   string
	startTime     time.Time
}

func (i *InboundContext) Service() string {
	return i.service
}

func (i *InboundContext) Component() string {
	return i.component
}

func (i *InboundContext) InterfaceType() string {
	return i.interfaceType
}

func (i *InboundContext) InterfaceID() string {
	return i.interfaceID
}

func (i *InboundContext) StartTime() time.Time {
	return i.startTime
}

func NewInboundContext(
	service string,
	component string,
	interfaceType string,
	interfaceID string,
	startTime time.Time,
) *InboundContext {
	return &InboundContext{
		service:       service,
		component:     component,
		interfaceType: interfaceType,
		interfaceID:   interfaceID,
		startTime:     startTime,
	}
}
