# isg-go
Interface service graph library for go

## isg-go-prometheus

This is a Prometheus reporter for the isg-go library. It provides functionality for tracking interface processing, and reporting the results to Prometheus in [specific format](docs/metrics.md).

## Installation

Use go get to install the package:

```bash
go get github.com/derbylock/isg-go/isg-go-prometheus@latest
```

## Usage
First, initialize the PrometheusReporter:

```go
package main

import (
    "github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
    "github.com/derbylock/isg-go/isg-go-prometheus/pkg/promisg"
)

func main() {
    ctxKeeper := isg.GetContextKeeper()
    registry := prometheus.NewRegistry()

    reporter := promisg.NewPrometheusReporter(ctxKeeper, time.Now, registry)
    reporter.Init()

    isg.SetDefaultContextKeeper(ctxKeeper)
    isg.SetDefaultReporter(reporter)
}
```

Then, use the reporter to track inbound and outbound interface processing:

```go
package httpserverimpl

import (
    "context"
    "github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
)

func getProductByID(ctx context.Context, id string) error {
    ctx, processingCtx := isg.Inbound(ctx, "products", "app", isg.HTTP, "getProductByID")
    defer processingCtx.Finish()

    // process...

    if err != nil {
        processingCtx.Fail()
        return fmt.Errorf("something wrong: %w", err)
    }

    return nil
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.