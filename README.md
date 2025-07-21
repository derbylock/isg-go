# isg-go
Interface service graph library for go

## isg-go-prometheus

This is a Prometheus reporter for the isg-go library. It provides functionality for tracking interface processing, and reporting the results to Prometheus in [specific format](docs/metrics.md).

## Installation

Use go get to install the package:

```bash
go get github.com/derbylock/isg-go/isg-go-prometheus
```

## Usage
First, initialize the PrometheusReporter:

```go
import (
    "github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
    "github.com/derbylock/isg-go/isg-go-prometheus/pkg/promisg"
)

func main() {
    ctxKeeper := isg.GetContextKeeper()
	registry := prometheus.NewRegistry()
	
    reporter := promisg.NewPrometheusReporter(ctxKeeper, time.Now, registry)
    reporter.Init()
}
```

Then, use the reporter to track inbound and outbound interface processing:

```go
func process(reporter *promisg.PrometheusReporter, ctx context.Context, service string, component string, ifType string, ifId string) context.Context {
    ctx, startedCtx := reporter.Inbound(ctx, service, component, ifType, ifId)
    defer startedCtx.Finished(isg.Success)
	
    // process...
    return ctx
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.