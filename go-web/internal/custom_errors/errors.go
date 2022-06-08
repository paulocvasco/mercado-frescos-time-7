package customerrors

import "errors"

var (
	ErrorSectionNotFound    = errors.New("section not found")
	ErrorStoreFailed        = errors.New("failed to store")
	ErrorEmptySection       = errors.New("empty section")
	ErrorSectionNumber      = errors.New("invalid parameter")
	ErrorCurrentCapacity    = errors.New("invalid current capacity")
	ErrorMinimumCapacity    = errors.New("invalid minimum capacity")
	ErrorMaximumCapacity    = errors.New("invalid maximum capacity")
	ErrorWarehouseID        = errors.New("invalid warehouse id")
	ErrorProductTypeID      = errors.New("invalid product type id")
	ErrorInvalidID          = errors.New("invalid id")
	ErrorInvalidIDParameter = errors.New("invalid parameter recieved as id")
	ErrorMissingAddres      = errors.New("address parameter is required")
	ErrorMissingTelephone   = errors.New("telephone parameter is required")
	ErrorMissingCapacity    = errors.New("capacity parameter is required")
	ErrorMissingTemperature = errors.New("temperature parameter is required")
	ErrorItemNotFound       = errors.New("item not found")
)
