package promisg

import (
	"github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var _ isg.Reporter = (*PrometheusReporter)(nil)

// PrometheusReporter is a struct that represents a reporter for tracking interface processing and reporting
// results to Prometheus.
type PrometheusReporter struct {
	ctxKeeper                   isg.ContextKeeper
	now                         TimeTracker
	registerer                  prometheus.Registerer
	inboundCounterVec           *prometheus.CounterVec
	inboundHistogramVec         *prometheus.HistogramVec
	inboundHistogramVecMinutes  *prometheus.HistogramVec
	outboundCounterVec          *prometheus.CounterVec
	outboundHistogramVec        *prometheus.HistogramVec
	outboundHistogramVecMinutes *prometheus.HistogramVec
}

// TimeTracker is a function that returns the current time
type TimeTracker func() time.Time

// NewPrometheusReporter creates a new PrometheusReporter.
// Use Init() method to initialize metrics
func NewPrometheusReporter(
	ctxKeeper isg.ContextKeeper,
	getNow TimeTracker,
	registerer prometheus.Registerer,
) *PrometheusReporter {
	return &PrometheusReporter{ctxKeeper: ctxKeeper, now: getNow, registerer: registerer}
}

// Init initializes the PrometheusReporter. It registers the metrics with the Prometheus registry.
func (p *PrometheusReporter) Init() {
	p.inboundCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "iif_count",
			Help: "Inbound interface processing count.",
		},
		[]string{"in_service", "in_component", "in_if_type", "in_if_id", "status"},
	)

	p.inboundHistogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "iif_duration",
			Help: "Inbound interface processing duration histogram.",
		},
		[]string{"in_service", "in_component", "in_if_type", "in_if_id", "status"},
	)

	p.inboundHistogramVecMinutes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "iif_duration_minutes",
			Help: "Inbound interface processing duration histogram in minutes.",
		},
		[]string{"in_service", "in_component", "in_if_type", "in_if_id", "status"},
	)

	p.outboundCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "oif_count",
			Help: "Outbound interface processing count.",
		},
		[]string{"in_service", "in_component", "in_if_type", "in_if_id", "out_service", "out_component", "out_if_type", "out_if_id", "status"},
	)

	p.outboundHistogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "oif_duration",
			Help: "Outbound interface processing duration histogram.",
		},
		[]string{"in_service", "in_component", "in_if_type", "in_if_id", "out_service", "out_component", "out_if_type", "out_if_id", "status"},
	)

	p.outboundHistogramVecMinutes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "oif_duration_minutes",
			Help: "Inbound interface processing duration histogram in minutes.",
		},
		[]string{"in_service", "in_component", "in_if_type", "in_if_id", "out_service", "out_component", "out_if_type", "out_if_id", "status"},
	)

	p.registerer.MustRegister(p.inboundCounterVec)
	p.registerer.MustRegister(p.inboundHistogramVec)
	p.registerer.MustRegister(p.inboundHistogramVecMinutes)
	p.registerer.MustRegister(p.outboundCounterVec)
	p.registerer.MustRegister(p.outboundHistogramVec)
	p.registerer.MustRegister(p.outboundHistogramVecMinutes)
}
