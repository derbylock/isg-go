package promisg

import (
	"context"
	"github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
	"time"
)

var emptyInboundContext = isg.NewInboundContext("-", "-", "-", "-", time.Time{})

func (p *PrometheusReporter) Outbound(
	ctx context.Context,
	service string,
	component string,
	interfaceType isg.InterfaceType,
	interfaceID string,
) (context.Context, isg.StartedContext) {
	now := p.now()
	inCtx := p.ctxKeeper.ExtractInboundContext(ctx)
	outCtx := isg.NewInboundContext(service, component, interfaceType, interfaceID, now)
	return ctx, NewOutboundStartedContext(p, inCtx, outCtx)
}

// OutboundStartedContext is a struct that represents a context for tracking the start of outbound processing.
type OutboundStartedContext struct {
	reporter *PrometheusReporter
	inCtx    startedContext
	outCtx   startedContext
}

func NewOutboundStartedContext(reporter *PrometheusReporter, inCtx startedContext, outCtx startedContext) *OutboundStartedContext {
	return &OutboundStartedContext{reporter: reporter, inCtx: inCtx, outCtx: outCtx}
}

// Finished is a method that marks the outbound processing as finished. It reports the processing status to Prometheus.
func (c *OutboundStartedContext) Finished(status isg.ProcessingStatus) {
	inCtx := c.inCtx
	if inCtx == nil {
		inCtx = emptyInboundContext
	}

	outCtx := c.outCtx
	s := string(status)
	c.reporter.outboundCounterVec.WithLabelValues(
		inCtx.Service(),
		inCtx.Component(),
		inCtx.InterfaceType().String(),
		inCtx.InterfaceID(),
		outCtx.Service(),
		outCtx.Component(),
		outCtx.InterfaceType().String(),
		outCtx.InterfaceID(),
		s,
	)

	now := c.reporter.now()
	dt := now.Sub(inCtx.StartTime())

	c.reporter.outboundHistogramVec.WithLabelValues(
		inCtx.Service(),
		inCtx.Component(),
		inCtx.InterfaceType().String(),
		inCtx.InterfaceID(),
		outCtx.Service(),
		outCtx.Component(),
		outCtx.InterfaceType().String(),
		outCtx.InterfaceID(),
		s,
	).Observe(dt.Seconds())

	c.reporter.outboundHistogramVecMinutes.WithLabelValues(
		inCtx.Service(),
		inCtx.Component(),
		inCtx.InterfaceType().String(),
		inCtx.InterfaceID(),
		outCtx.Service(),
		outCtx.Component(),
		outCtx.InterfaceType().String(),
		outCtx.InterfaceID(),
		s,
	).Observe(dt.Minutes())
}
