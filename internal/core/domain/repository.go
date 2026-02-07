package domain

import "context"

// ShipmentRepository defines the interface for shipment data access.
type ShipmentRepository interface {
	GetByTrackingNumber(ctx context.Context, trackingNumber string) (*Shipment, error)
	GetByOrderID(ctx context.Context, orderID string) (*Shipment, error)
}
