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

### Prerequisites

- Go 1.25+
- [golangci-lint](https://golangci-lint.run/welcome/install/) v2+

### Local Development

```bash
# Install dependencies
go mod tidy
go mod download

# Build
go build ./...

# Test
go test ./...

# Lint (must pass before PR merge)
golangci-lint run --timeout=10m

# Run locally (requires .env or env vars)
go run cmd/main.go
```

### Pre-push Checklist

```bash
go build ./... && go test ./... && golangci-lint run --timeout=10m
```

## License

MIT
