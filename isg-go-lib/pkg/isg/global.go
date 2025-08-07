package isg

import "context"

var (
	defaultReporter      Reporter
	defaultContextKeeper ContextKeeper = NewContextValueContextKeeper()
)

func DefaultReporter() Reporter {
	return defaultReporter
}

func SetDefaultReporter(reporter Reporter) {
	defaultReporter = reporter
}

func DefaultContextKeeper() ContextKeeper {
	return defaultContextKeeper
}

func SetDefaultContextKeeper(keeper ContextKeeper) {
	defaultContextKeeper = keeper
}

func Inbound(ctx context.Context, service string, component string, interfaceType InterfaceType, interfaceID string) (context.Context, *StatelessStartedContext) {
	if defaultReporter == nil {
		return ctx, NewStatelessStartedContext(nilStartedContext)
	}

	newCtx, startedCtx := defaultReporter.Inbound(ctx, service, component, interfaceType, interfaceID)
	return newCtx, NewStatelessStartedContext(startedCtx)
}

func Outbound(ctx context.Context, service string, component string, interfaceType InterfaceType, interfaceID string) (context.Context, *StatelessStartedContext) {
	if defaultReporter == nil {
		return ctx, NewStatelessStartedContext(nilStartedContext)
	}

	newCtx, startedCtx := defaultReporter.Outbound(ctx, service, component, interfaceType, interfaceID)
	return newCtx, NewStatelessStartedContext(startedCtx)
}

type StatelessStartedContext struct {
	parent StartedContext
	status ProcessingStatus
}

func NewStatelessStartedContext(parent StartedContext) *StatelessStartedContext {
	return &StatelessStartedContext{parent: parent, status: ProcessingStatusOK}
}

func (c *StatelessStartedContext) Finished(status ProcessingStatus) {
	c.parent.Finished(status)
}

func (c *StatelessStartedContext) SetStatus(status ProcessingStatus) {
	c.status = status
}

func (c *StatelessStartedContext) Finish() {
	c.parent.Finished(c.status)
}

func (c *StatelessStartedContext) Fail() {
	c.status = ProcessingStatusFail
}
