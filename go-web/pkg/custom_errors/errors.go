package customerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrorSectionNotFound       = errors.New("section not found")
	ErrorStoreFailed           = errors.New("failed to store")
	ErrorEmptySection          = errors.New("empty section")
	ErrorSectionNumber         = errors.New("invalid parameter")
	ErrorCurrentCapacity       = errors.New("invalid current capacity")
	ErrorMinimumCapacity       = errors.New("invalid minimum capacity")
	ErrorMaximumCapacity       = errors.New("invalid maximum capacity")
	ErrorWarehouseID           = errors.New("invalid warehouse id")
	ErrorProductTypeID         = errors.New("invalid product type id")
	ErrorInvalidID             = errors.New("invalid id")
	ErrorInvalidIDParameter    = errors.New("invalid parameter recieved as id")
	ErrorMissingAddres         = errors.New("address parameter is required")
	ErrorMissingTelephone      = errors.New("telephone parameter is required")
	ErrorMissingCapacity       = errors.New("capacity parameter is required")
	ErrorMissingTemperature    = errors.New("temperature parameter is required")
	ErrorItemNotFound          = errors.New("item not found")
	ErrorConflict              = errors.New("conflict error detected")
	ErrorCardIdAlreadyExists   = errors.New("card Number Id already exist")
	ErrorInvalidDB             = errors.New("invalid database")
	ErrorSectionAlreadyExists  = errors.New("section number already exists")
	ErrorWarehouseCodeConflict = errors.New("warehouse code already exist")
	ErrorMarshallJson          = errors.New("malformed json")
)

func ErrorHandleResponse(err error) (int, string) {
	{ // custom errors
		if errors.Is(err, ErrorInvalidID) {
			return http.StatusNotFound, err.Error()
		}
		if errors.Is(err, ErrorInvalidIDParameter) {
			return http.StatusNotFound, err.Error()
		}
		if errors.Is(err, ErrorMissingAddres) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorMissingTelephone) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorMissingTemperature) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorItemNotFound) {
			return http.StatusNotFound, err.Error()
		}
		if errors.Is(err, ErrorMinimumCapacity) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorSectionNotFound) {
			return http.StatusNotFound, err.Error()
		}
		if errors.Is(err, ErrorStoreFailed) {
			return http.StatusInternalServerError, err.Error()
		}
		if errors.Is(err, ErrorEmptySection) {
			return http.StatusNotFound, err.Error()
		}
		if errors.Is(err, ErrorSectionNumber) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorCurrentCapacity) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorMinimumCapacity) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorMaximumCapacity) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorWarehouseID) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorProductTypeID) {
			return http.StatusUnprocessableEntity, err.Error()
		}
		if errors.Is(err, ErrorConflict) {
			return http.StatusConflict, err.Error()
		}
		if errors.Is(err, ErrorInvalidDB) {
			return http.StatusInternalServerError, err.Error()
		}
		if errors.Is(err, ErrorCardIdAlreadyExists) {
			return http.StatusConflict, err.Error()
		}
		if errors.Is(err, ErrorSectionAlreadyExists) {
			return http.StatusConflict, err.Error()
		}
		if errors.Is(err, ErrorWarehouseCodeConflict) {
			return http.StatusConflict, err.Error()
		}
	}
	{ // validate errors
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fields := []string{}
			for _, v := range ve {
				fields = append(fields, v.Field())
			}
			return http.StatusUnprocessableEntity, fmt.Sprintf("validation error in the field(s): %v", strings.ToLower(strings.Join(fields, ", ")))
		}
		var js *json.SyntaxError
		if errors.As(err, &js) {
			return http.StatusBadRequest, "bad JSON"
		}
		var jt *json.UnmarshalTypeError
		if errors.As(err, &jt) {
			return http.StatusBadRequest, fmt.Sprintf("type error in %v", jt.Field)
		}
	}
	return http.StatusInternalServerError, "internal error"
}

type ErrorFormat struct {
	ErrorCode    int
	ErrorMessage string
}
