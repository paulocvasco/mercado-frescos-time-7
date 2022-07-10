package reportsellers_test

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/reportsellers"
	"mercado-frescos-time-7/go-web/internal/reportsellers/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetAll(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   []models.ReportSeller
		expectResponse []models.ReportSeller
		expectError    error
		message        string
	}
	response := []models.ReportSeller{
		{LocalityID: "1", Locality_name: "Buritizeiro", SellerCount: "7"},
		{LocalityID: "2", Locality_name: "Buritizeiro", SellerCount: "5"},
		{LocalityID: "3", Locality_name: "Buritizeiro", SellerCount: "2"},
	}

	testCases := []tests{
		{"ReportSeller All", response, response, nil, "Values Differents"},
		{"ReportSeller return Error", []models.ReportSeller{}, []models.ReportSeller{}, errors.New("Error"), "Value Error Different"},
	}

	for _, value := range testCases {
		mockRepository := mocks.NewRepository(t)
		service := reportsellers.NewService(mockRepository)
		mockRepository.On("ReportSellers", 0).Return(value.mockResponse, value.expectError)
		resp, err := service.ReportSellers(0)
		assert.Equal(t, value.expectResponse, resp)
		assert.Equal(t, value.expectError, err)
	}

}
