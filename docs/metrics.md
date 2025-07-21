# Service Dependency Monitoring in Prometheus Format

## Monitoring Format

### Dimensions
This section lists standard dimension types referenced by the standard

| Dimension     | Description                                                                                                    | Example                               |
|---------------|----------------------------------------------------------------------------------------------------------------|---------------------------------------|
| in_service    | Inbound Service Name                                                                                           | odin                                  |
| in_component  | Inbound Component Name                                                                                         | app                                   |
| out_service   | Outbound Service Name                                                                                          | odin                                  |
| out_component | Outbound Component Name                                                                                        | app                                   |
| in_if_type    | Input Interface Type: http, grpc, topic, job                                                                   | http                                  |
| in_if_id      | Input Interface ID (OpenAPI's operationID, gRPC method, topic name, worker name)                               | getStore, perf.PerfTest.PerfTestScale |
| out_if_type   | Output Interface Type: http, grpc, topic, job, db                                                              | http                                  |
| out_if_id     | Output Interface ID (OpenAPI's operationID, gRPC method, topic name, worker name, stored procedure / query ID) | stores.StoresStats/GetStoresBaseStats |
| status        | Status (HTTP/gRPC request status)                                                                              | 200, ok, failed, timeout              |


### Input Interfaces

Generally, services should provide the following metrics for input interfaces:

| Metric Type | Metric Name | Description             | Dimensions                                             |
|-------------|-------------|-------------------------|--------------------------------------------------------|
| Counter     | if_count    | Request call count      | in_service, in_component, in_if_type, in_if_id, status |
| Histogram   | if_duration | Request processing time | in_service, in_component, in_if_type, in_if_id, status |

Types of input interfaces:

- HTTP API
- gRPC API
- Kafka consumer
- Periodic worker
- Other interfaces

### Output Interfaces

| Metric Type | Metric Name | Description             | Dimensions                                                                                                 |
|-------------|-------------|-------------------------|------------------------------------------------------------------------------------------------------------|
| Counter     | if_count    | Request call count      | in_service, in_component, in_if_type, in_if_id, out_service, out_component, out_if_type, out_if_id, status |
| Histogram   | if_duration | Request processing time | in_service, in_component, in_if_type, in_if_id, out_service, out_component, out_if_type, out_if_id, status |


Types of output interfaces:
- HTTP API
- gRPC API
- Kafka producer
- Job
- Database
- Other external integrations