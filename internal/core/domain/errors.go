package domain

import "errors"

// ErrShipmentNotFound indicates that the shipment could not be found.
var ErrShipmentNotFound = errors.New("shipment not found")
