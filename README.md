# shipping-service

Shipping microservice for tracking and cost estimation.

## Features

- Shipment tracking
- Cost estimation
- Get shipment by order

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/shipping/track` | Track shipment |
| `GET` | `/api/v1/shipping/estimate` | Estimate cost |
| `GET` | `/api/v1/shipping/orders/:id` | Get by order ID |

## Tech Stack

- Go + Gin framework
- PostgreSQL 16 (supporting-db cluster, cross-namespace)
- PgBouncer connection pooling
- OpenTelemetry tracing

## Development

```bash
go mod download
go test ./...
go run cmd/main.go
```

## License

MIT
