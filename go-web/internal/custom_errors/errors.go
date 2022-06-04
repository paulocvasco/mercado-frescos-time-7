package customerrors

import "errors"

var (
	ErrorInvalidID          = errors.New("invalid id")
	ErrorInvalidIDParameter = errors.New("invalid parameter recieved as id")
	ErrorMissingAddres      = errors.New("address parameter is required")
	ErrorMissingTelephone   = errors.New("telephone parameter is required")
	ErrorMissingCapacity    = errors.New("capacity parameter is required")
	ErrorMissingTemperature = errors.New("temperature parameter is required")
	ErrorItemNotFound       = errors.New("item not found")
)
