// Package promisg provides a Prometheus reporter for the isg-go library.
package promisg

import (
	"context"
	"github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
)

func (p *PrometheusReporter) Inbound(
	ctx context.Context,
	service string,
	component string,
	interfaceType isg.InterfaceType,
	interfaceID string,
) (context.Context, isg.StartedContext) {
	t := p.now()
	inboundStartedContext := isg.NewInboundContext(service, component, interfaceType, interfaceID, t)
	return p.ctxKeeper.KeepInboundContext(ctx, inboundStartedContext),
		NewInboundStartedContext(p, inboundStartedContext)
}

// InboundStartedContext is a struct that represents a context for tracking the start of inbound processing.
type InboundStartedContext struct {
	reporter *PrometheusReporter
	inCtx    startedContext
}

func NewInboundStartedContext(reporter *PrometheusReporter, startedCtx *isg.InboundContext) *InboundStartedContext {
	return &InboundStartedContext{reporter: reporter, inCtx: startedCtx}
}

// Finished is a method that marks the inbound processing as finished. It reports the processing status to Prometheus.
func (c *InboundStartedContext) Finished(status isg.ProcessingStatus) {
	startedCtx := c.inCtx
	s := string(status)

	now := c.reporter.now()
	dt := now.Sub(startedCtx.StartTime())

	c.reporter.inboundHistogramVec.WithLabelValues(
		startedCtx.Service(),
		startedCtx.Component(),
		startedCtx.InterfaceType().String(),
		startedCtx.InterfaceID(),
		s,
	).Observe(dt.Seconds())

	c.reporter.inboundHistogramVecMinutes.WithLabelValues(
		startedCtx.Service(),
		startedCtx.Component(),
		startedCtx.InterfaceType().String(),
		startedCtx.InterfaceID(),
		s,
	).Observe(dt.Minutes())
}
