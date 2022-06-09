package customerrors

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

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
	ErrorConflict           = errors.New("conflict error detected")
)

func ErrorHandleResponse(err error) (int, string) {
	{ // custom errors
		if errors.Is(err, ErrorInvalidID) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorInvalidIDParameter) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorMissingAddres) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorMissingTelephone) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorMissingTemperature) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorItemNotFound) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorMinimumCapacity) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorSectionNotFound) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorStoreFailed) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorEmptySection) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorSectionNumber) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorCurrentCapacity) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorMinimumCapacity) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorMaximumCapacity) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorWarehouseID) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorProductTypeID) {
			return 404, err.Error()
		}
		if errors.Is(err, ErrorConflict) {
			return 409, err.Error()
		}
	}
	{ // validate errors
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, v := range ve {
				return 404, fmt.Sprintf("%v field validation failed", v.Field())
			}
		}
		var js *json.SyntaxError
		if errors.As(err, &js) {
			return 404, "bad JSON"
		}
		var jt *json.UnmarshalTypeError
		if errors.As(err, &jt) {
			return 404, fmt.Sprintf("type error in %v", jt.Field)
		}
	}
	return 500, "internal error"
}
