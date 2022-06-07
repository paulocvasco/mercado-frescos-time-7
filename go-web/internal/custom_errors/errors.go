package customerrors

import "errors"

var (
	ErrorInvalidID       = errors.New("invalid id")
	ErrorSectionNotFound = errors.New("section not found")
	ErrorStoreFailed     = errors.New("failed to store")
	ErrorEmptySection    = errors.New("empty section")
	ErrorSectionNumber   = errors.New("invalid parameter")
	ErrorCurrentCapacity = errors.New("invalid current capacity")
	ErrorMinimumCapacity = errors.New("invalid minimum capacity")
	ErrorMaximumCapacity = errors.New("invalid maximum capacity")
	ErrorWarehouseID     = errors.New("invalid warehouse id")
	ErrorProductTypeID   = errors.New("invalid product type id")
)
