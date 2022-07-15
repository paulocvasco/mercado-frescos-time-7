package locality_test

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/locality"
	"mercado-frescos-time-7/go-web/internal/locality/mocks"
	"mercado-frescos-time-7/go-web/internal/models"

	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStore(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   models.Locality
		expectResponse models.Locality
		expectError    error
		message        string
		errorLastID    error
	}
	response := models.Locality{
	Id: "1", 
	Locality_name: "Juan", 
	Province_name: "Minas",
	Country_name: "Argentina",
	}
	testCases := []tests{
		{"Store", response, response, nil, "Errro created", nil},
		{"Error", models.Locality{}, models.Locality{}, errors.New("Error"), "Errro created", nil},
		{"Store Error LastID", models.Locality{}, models.Locality{}, errors.New("Error"), "Id doesn`t exist", errors.New("Error")},
		{"Error LastID", models.Locality{}, models.Locality{}, errors.New("Error"), "Errro created", nil},
	}
	for _, value := range testCases {
		mockRepository := mocks.NewRepository(t)
		service := locality.NewService(mockRepository)

		mockRepository.On("Store",
			value.mockResponse).
			Return(value.mockResponse, value.expectError).Maybe()

		resp, err := service.
			Store(value.mockResponse)
		assert.Equal(t, value.expectResponse, resp)
		assert.Equal(t, value.expectError, err)

	}

}