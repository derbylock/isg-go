package promisg

import (
	"context"
	"github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestPrometheusReporter_Inbound(t *testing.T) {
	now := time.Now()
	getNow := func() time.Time { return now }

	registry := prometheus.NewRegistry()
	reporter := NewPrometheusReporter(isg.DefaultContextKeeper(), getNow, registry)
	reporter.Init()

	ctx := context.Background()

	testFunc := func() {
		_, startedCtx := reporter.Inbound(ctx, "service1", "component1", "type1", "id1")
		defer startedCtx.Finished(isg.ProcessingStatusOK)

		now = now.Add(time.Second)
	}

	testFunc()

	assert.NoError(t, testutil.GatherAndCompare(registry,
		strings.NewReader(
			`
			# HELP iif_count Inbound interface processing count.
			# TYPE iif_count counter
			iif_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 0
			# HELP iif_duration Inbound interface processing duration histogram.
			# TYPE iif_duration histogram
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.005"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.01"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.025"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.05"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.1"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.25"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.5"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="1"} 1
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="2.5"} 1
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="5"} 1
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="10"} 1
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="+Inf"} 1
			iif_duration_sum{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 1
			iif_duration_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 1
			# HELP iif_duration_minutes Inbound interface processing duration histogram in minutes.
			# TYPE iif_duration_minutes histogram
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.005"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.01"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.025"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.05"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.1"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.25"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.5"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="1"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="2.5"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="5"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="10"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="+Inf"} 1
			iif_duration_minutes_sum{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 0.016666666666666666
			iif_duration_minutes_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 1
			`,
		),
	))
}

func TestPrometheusReporter_Outbound(t *testing.T) {
	now := time.Now()
	getNow := func() time.Time { return now }

	registry := prometheus.NewRegistry()
	reporter := NewPrometheusReporter(isg.DefaultContextKeeper(), getNow, registry)
	reporter.Init()

	ctx := context.Background()

	testFunc := func() {
		inCtx, inboundStartedCtx := reporter.Inbound(
			ctx,
			"service1",
			"component1",
			"type1",
			"id1",
		)
		defer inboundStartedCtx.Finished(isg.ProcessingStatusOK)

		testOutBoundFunc := func(ctx context.Context) {
			_, outboundStartedCtx := reporter.Outbound(
				ctx,
				"service2",
				"component2",
				"type2",
				"id2",
			)
			defer outboundStartedCtx.Finished(isg.ProcessingStatusOK)

			now = now.Add(time.Second)
		}

		testOutBoundFunc(inCtx)
		now = now.Add(time.Minute)
	}

	testFunc()

	assert.NoError(t, testutil.GatherAndCompare(registry,
		strings.NewReader(
			`
			# HELP iif_count Inbound interface processing count.
			# TYPE iif_count counter
			iif_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 0
			# HELP iif_duration Inbound interface processing duration histogram.
			# TYPE iif_duration histogram
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.005"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.01"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.025"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.05"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.1"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.25"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.5"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="1"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="2.5"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="5"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="10"} 0
			iif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="+Inf"} 1
			iif_duration_sum{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 61
			iif_duration_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 1
			# HELP iif_duration_minutes Inbound interface processing duration histogram in minutes.
			# TYPE iif_duration_minutes histogram
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.005"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.01"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.025"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.05"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.1"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.25"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="0.5"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="1"} 0
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="2.5"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="5"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="10"} 1
			iif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok",le="+Inf"} 1
			iif_duration_minutes_sum{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 1.0166666666666666
			iif_duration_minutes_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",status="ok"} 1
			# HELP oif_count Outbound interface processing count.
			# TYPE oif_count counter
			oif_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok"} 0
			# HELP oif_duration Outbound interface processing duration histogram.
			# TYPE oif_duration histogram
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.005"} 0
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.01"} 0
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.025"} 0
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.05"} 0
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.1"} 0
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.25"} 0
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.5"} 0
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="1"} 1
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="2.5"} 1
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="5"} 1
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="10"} 1
			oif_duration_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="+Inf"} 1
			oif_duration_sum{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok"} 1
			oif_duration_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok"} 1
			# HELP oif_duration_minutes Inbound interface processing duration histogram in minutes.
			# TYPE oif_duration_minutes histogram
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.005"} 0
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.01"} 0
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.025"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.05"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.1"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.25"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="0.5"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="1"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="2.5"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="5"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="10"} 1
			oif_duration_minutes_bucket{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok",le="+Inf"} 1
			oif_duration_minutes_sum{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok"} 0.016666666666666666
			oif_duration_minutes_count{in_component="component1",in_if_id="id1",in_if_type="type1",in_service="service1",out_component="component2",out_if_id="id2",out_if_type="type2",out_service="service2",status="ok"} 1
			`,
		),
	))
}
