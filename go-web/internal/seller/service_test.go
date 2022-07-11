package seller_test

import (
	"encoding/json"
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/seller"
	"mercado-frescos-time-7/go-web/internal/seller/mocks"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   []models.Seller
		expectResponse []models.Seller
		expectError    error
		message        string
	}
	response := []models.Seller{
		{
			ID:           1,
			Cid:          123,
			Company_name: "Meli1",
			Address:      "Rua 1",
			Telephone:    "(11) 33387767",
		},
		{
			ID:           2,
			Cid:          1234,
			Company_name: "Meli2",
			Address:      "Rua 3",
			Telephone:    "(11) 33387768",
		},
		{
			ID:           3,
			Cid:          12356,
			Company_name: "Meli3",
			Address:      "Rua 3",
			Telephone:    "(11) 33387769",
		},
	}

	testCases := []tests{
		{"Get all Sellers", response, response, nil, "Values Differents"},
		{"GetAll return Error", nil, nil, errors.New("Error"), "Value Error Different"},
	}

	for _, value := range testCases {
		mockRepository := mocks.NewRepository(t)
		service := seller.NewService(mockRepository)
		mockRepository.On("GetAll").Return(value.mockResponse, value.expectError)
		resp, err := service.GetAll()
		assert.Equal(t, value.expectResponse, resp, value.name, value.message)
		assert.Equal(t, value.expectError, err, value.name, value.message)

	}

}

func TestGetId(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   models.Seller
		expectResponse models.Seller
		expectError    error
		message        string
		params         int
	}
	response := []models.Seller{
		{
			ID:           1,
			Cid:          123,
			Company_name: "Meli1",
			Address:      "Rua 1",
			Telephone:    "(11) 33387767",
		},
		{
			ID:           2,
			Cid:          1234,
			Company_name: "Meli2",
			Address:      "Rua 3",
			Telephone:    "(11) 33387768",
		},
		{
			ID:           3,
			Cid:          12356,
			Company_name: "Meli3",
			Address:      "Rua 3",
			Telephone:    "(11) 33387769",
		},
	}

	testCases := []tests{
		{"Get Id = 1", response[0], response[0], nil, "Id doesn`t exist", 1},
		{"Get Id = 2", response[1], response[1], nil, "Id doesn`t exist", 2},
		{"Get Id = 3", response[2], response[2], nil, "Id doesn`t exist", 3},
		{"Get Id = 4", models.Seller{}, models.Seller{}, customerrors.ErrorInvalidID, "Id doesn`t exist", 4},
	}
	for _, value := range testCases {
		mockRepository := mocks.NewRepository(t)
		service := seller.NewService(mockRepository)
		mockRepository.On("GetId", value.params).Return(value.mockResponse, value.expectError)
		resp, err := service.GetId(value.params)
		assert.Equal(t, value.expectResponse, resp, value.message)
		assert.Equal(t, value.expectError, err, value.message)
	}
}

func TestDelete(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   error
		expectResponse error
		expectError    error
		message        string
		params         int
	}
	testCases := []tests{
		{"Get Id = 1", nil, nil, nil, "Id doesn`t exist", 1},
		{"Get Id = 4", errors.New("Error"), nil, errors.New("Error"), "Id doesn`t exist", 4},
	}
	for _, value := range testCases {
		mockRepository := mocks.NewRepository(t)
		service := seller.NewService(mockRepository)
		mockRepository.On("Delete", value.params).Return(value.mockResponse, value.expectError)
		err := service.Delete(value.params)
		assert.Equal(t, value.expectError, err, value.message)
	}
}

func TestStore(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   models.Seller
		expectResponse models.Seller
		expectError    error
		message        string
		errorLastID    error
	}
	response := models.Seller{
		ID:           1,
		Cid:          123,
		Company_name: "Meli1",
		Address:      "Rua 1",
		Telephone:    "(11) 33387767",
	}
	testCases := []tests{
		{"Store", response, response, nil, "Errro created", nil},
		{"Error", models.Seller{}, models.Seller{}, errors.New("Error"), "Errro created", nil},
		{"Store Error LastID", models.Seller{}, models.Seller{}, errors.New("Error"), "Id doesn`t exist", errors.New("Error")},
		{"Error LastID", models.Seller{}, models.Seller{}, errors.New("Error"), "Errro created", nil},
	}
	for _, value := range testCases {
		mockRepository := mocks.NewRepository(t)
		service := seller.NewService(mockRepository)

		mockRepository.On("LastID").Return((value.mockResponse.ID - 1), value.errorLastID).Maybe()

		mockRepository.On("Store",
			value.mockResponse).
			Return(value.mockResponse, value.expectError).Maybe()

		resp, err := service.
			Store(value.mockResponse)
		assert.Equal(t, value.expectResponse, resp, value.message)
		assert.Equal(t, value.expectError, err, value.message)

	}

}

func TestUpdate(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   models.Seller
		expectResponse models.Seller
		valueUpdate    models.Seller
		expectError    error
		message        string
		getIdError     error
		cidError       error
	}
	response := []models.Seller{{
		ID:           1,
		Cid:          123,
		Company_name: "Mercado Livre1",
		Address:      "Rua 1",
		Telephone:    "(11) 3333-3333",
	},
	}
	valueUpdate := []models.Seller{{
		ID:           1,
		Cid:          123,
		Company_name: "Mercado Livre1",
		Address:      "Rua 1",
		Telephone:    "(11) 3333-3333",
	},
	}

	testCases := []tests{
		{"Update", response[0], response[0], valueUpdate[0], nil, "Errro Update", nil, nil},
		{"Update Error", models.Seller{}, models.Seller{}, models.Seller{}, errors.New("Error"), "Errro Update", nil, nil},
		{"Update Error Cid", models.Seller{}, models.Seller{}, models.Seller{}, customerrors.ErrorConflict, "Errro Update", nil, customerrors.ErrorConflict},
		{"Update Error Get Id", models.Seller{}, models.Seller{}, models.Seller{}, customerrors.ErrorInvalidID, "Errro Update", customerrors.ErrorInvalidID, nil},
	}
	for _, value := range testCases {
		mockRepository := mocks.NewRepository(t)
		service := seller.NewService(mockRepository)

		mockRepository.On("GetId", value.expectResponse.ID).Return(value.mockResponse, value.getIdError).Maybe()
		mockRepository.On("CheckCid", value.expectResponse.Cid).Return(value.mockResponse, value.cidError).Maybe()
		mockRepository.On("Update", seller.Seller(value.mockResponse), value.mockResponse.ID).
			Return(value.expectResponse, value.expectError).Maybe()

		sellerByte, _ := json.Marshal(value.valueUpdate)
		resp, err := service.Update(sellerByte, value.mockResponse.ID)

		assert.Equal(t, value.expectResponse, resp, value.message)
		assert.Equal(t, value.expectError, err, value.message)

	}
}
