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

func Inbound(ctx context.Context, service string, component string, interfaceType InterfaceType, interfaceID string) (context.Context, StartedContext) {
	if defaultReporter == nil {
		return ctx, nilStartedContext
	}

	return defaultReporter.Inbound(ctx, service, component, interfaceType, interfaceID)
}

func Outbound(ctx context.Context, service string, component string, interfaceType InterfaceType, interfaceID string) (context.Context, StartedContext) {
	if defaultReporter == nil {
		return ctx, nilStartedContext
	}

	return defaultReporter.Outbound(ctx, service, component, interfaceType, interfaceID)
}
