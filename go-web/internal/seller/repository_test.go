package seller_test

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/seller"
	"mercado-frescos-time-7/go-web/internal/seller/mocks"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllRepository(t *testing.T) {
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

	seller.Populate()

	testCases := []tests{
		{"Get all Sellers", response, response, nil, "Values Differents"},
		{"GetAll return Error", nil, []models.Seller{}, errors.New("Error"), "Value Error Different"},
	}

	for _, value := range testCases {
		mockDB := mocks.NewDB(t)
		mockRepository := seller.NewRepository(mockDB)
		mockDB.On("Load", mock.Anything).Return(value.expectError)
		resp, err := mockRepository.GetAll()
		assert.Equal(t, value.expectResponse, resp, value.name, value.message)
		assert.Equal(t, value.expectError, err, value.name, value.message)

	}
	seller.Clean()

}

func TestGetIdRepositiory(t *testing.T) {
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

	seller.Populate()

	testCases := []tests{
		{"Get Id = 1", response[0], response[0], nil, "Id doesn`t exist", 1},
		{"Get Id = 2", response[1], response[1], nil, "Id doesn`t exist", 2},
		{"Get Id = 3", response[2], response[2], nil, "Id doesn`t exist", 3},
		{"Get Id = 4", models.Seller{}, models.Seller{}, customerrors.ErrorInvalidID, "Id doesn`t exist", 4},
	}
	for _, value := range testCases {
		mockDB := mocks.NewDB(t)
		mockRepository := seller.NewRepository(mockDB)
		mockDB.On("Load", mock.Anything).Return(value.expectError)
		resp, err := mockRepository.GetId(value.params)
		assert.Equal(t, value.expectResponse, resp, value.message)
		assert.Equal(t, value.expectError, err, value.message)
	}
	seller.Clean()
}

func TestDeleteRepository(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   error
		expectResponse error
		expectError    error
		message        string
		params         int
	}

	seller.Populate()

	testCases := []tests{
		{"Get Id = 1", nil, nil, nil, "Id doesn`t exist", 1},
		{"Get Id = 3", errors.New("Error"), nil, errors.New("Error"), "Id doesn`t exist", 3},
	}
	for _, value := range testCases {
		mockDB := mocks.NewDB(t)
		mockRepository := seller.NewRepository(mockDB)
		mockDB.On("Load", mock.Anything).Return(nil)
		mockDB.On("Save", mock.Anything).Return(value.expectError)
		err := mockRepository.Delete(value.params)
		assert.Equal(t, value.expectError, err, value.message)
	}
	seller.Clean()
}

func TestStoreRepository(t *testing.T) {
	type tests struct {
		name              string
		mockResponse      models.Seller
		expectResponse    models.Seller
		expectMockError   error
		message           string
		expectResultError error
	}
	response := models.Seller{
		ID:           1,
		Cid:          1234567,
		Company_name: "Meli1",
		Address:      "Rua 1",
		Telephone:    "(11) 33387767",
	}

	seller.Populate()

	testCases := []tests{
		{"Store", response, response, nil, "Errro created", nil},
		{"Error on Created", models.Seller{}, models.Seller{}, errors.New("Error"), "Errro created", errors.New("Error")},
	}
	for _, value := range testCases {
		mockDB := mocks.NewDB(t)
		mockRepository := seller.NewRepository(mockDB)
		mockDB.On("Load", mock.Anything).Return(nil)
		mockDB.On("Save", mock.Anything).Return(value.expectMockError)
		resp, err := mockRepository.Store(value.mockResponse)
		assert.Equal(t, value.expectResponse, resp, value.message)
		assert.Equal(t, value.expectResultError, err, value.message)

	}

	seller.Clean()
}

func TestStoreRepository2(t *testing.T) {
	type tests struct {
		name              string
		mockResponse      models.Seller
		expectResponse    models.Seller
		expectMockError   error
		message           string
		expectResultError error
	}
	response := models.Seller{
		ID:           1,
		Cid:          1234567,
		Company_name: "Meli1",
		Address:      "Rua 1",
		Telephone:    "(11) 33387767",
	}

	seller.Populate()

	testCases := []tests{
		{"Error in Load", response, models.Seller{}, errors.New("Error"), "load fail", errors.New("Error")},
		{"Conflict in Cid", models.Seller{Cid: 123}, models.Seller{}, nil, "Conflict", customerrors.ErrorConflict},
	}
	for _, value := range testCases {
		mockDB := mocks.NewDB(t)
		mockRepository := seller.NewRepository(mockDB)
		mockDB.On("Load", mock.Anything).Return(value.expectMockError)
		resp, err := mockRepository.Store(value.mockResponse)
		assert.Equal(t, value.expectResponse, resp, value.message)
		assert.Equal(t, value.expectResultError, err, value.message)

	}

	seller.Clean()
}

func TestUpdateRepository2(t *testing.T) {
	type tests struct {
		name            string
		mockResponse    models.Seller
		expectResponse  models.Seller
		valueUpdate     models.Seller
		expectMockError error
		message         string
		getIdError      error
		cidError        error
		expectResultError error
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

	seller.Populate()

	testCases := []tests{
		{"Update fail Load", response[0], models.Seller{}, valueUpdate[0], errors.New("Error"), "Errro Update", nil, nil, errors.New("Error")},
	}
	for _, value := range testCases {
		mockDB := mocks.NewDB(t)
		mockRepository := seller.NewRepository(mockDB)
		mockDB.On("Load", mock.Anything).Return(value.expectMockError)


		resp, err := mockRepository.Update(seller.Seller(value.valueUpdate), value.mockResponse.ID)

		assert.Equal(t, value.expectResponse, resp, value.message)
		assert.Equal(t, value.expectResultError, err, value.message)

	}
	seller.Clean()
}

func TestUpdateRepository(t *testing.T) {
	type tests struct {
		name            string
		mockResponse    models.Seller
		expectResponse  models.Seller
		valueUpdate     models.Seller
		expectMockError error
		message         string
		getIdError      error
		cidError        error
		expectResultError error
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

	seller.Populate()

	testCases := []tests{
		{"Update", response[0], response[0], valueUpdate[0], nil, "Errro Update", nil, nil, nil},
		{"Update fail Save", response[0], models.Seller{}, valueUpdate[0], errors.New("Error"), "Errro Update", nil, nil, errors.New("Error")},

	}
	for _, value := range testCases {
		mockDB := mocks.NewDB(t)
		mockRepository := seller.NewRepository(mockDB)
		mockDB.On("Load", mock.Anything).Return(nil)
		mockDB.On("Save", mock.Anything).Return(value.expectMockError)

		resp, err := mockRepository.Update(seller.Seller(value.valueUpdate), value.mockResponse.ID)

		assert.Equal(t, value.expectResponse, resp, value.message)
		assert.Equal(t, value.expectResultError, err, value.message)

	}
	seller.Clean()
}

func TestLastIDRepository(t *testing.T){
	mockDB := mocks.NewDB(t)
	mockRepository := seller.NewRepository(mockDB)
	mockDB.On("Load", mock.Anything).Return(nil)

	seller.Populate()

	lastid, _ := mockRepository.LastID()
	assert.Equal(t, 3, lastid, "igual")
	seller.Clean()
}

func TestLastIDRepositoryError(t *testing.T){
	mockDB := mocks.NewDB(t)
	mockRepository := seller.NewRepository(mockDB)

	seller.Populate()

	mockDB.On("Load", mock.Anything).Return(errors.New("Error"))
	_, err := mockRepository.LastID()
	assert.Equal(t, errors.New("Error"), err, "igual")
	seller.Clean()
}