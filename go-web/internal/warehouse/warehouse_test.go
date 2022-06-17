package warehouse_test

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/warehouse"
	"mercado-frescos-time-7/go-web/internal/warehouse/mock/mockRepository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	type mockResponse struct {
		data models.Warehouses
		err  error
	}
	type getAllExpected struct {
		data models.Warehouses
		err  error
	}
	type testData struct {
		testName     string
		response     mockResponse
		expectResult getAllExpected
	}

	testsCases := []testData{
		{
			testName: "should return all warehouses",
			response: mockResponse{
				data: models.Warehouses{Warehouses: []models.Warehouse{
					{ID: 0, Address: "foo", Telephone: "foo"},
					{ID: 1, Address: "foo", Telephone: "foo"},
				},
				},
				err: nil,
			},
			expectResult: getAllExpected{
				data: models.Warehouses{Warehouses: []models.Warehouse{
					{ID: 0, Address: "foo", Telephone: "foo"},
					{ID: 1, Address: "foo", Telephone: "foo"},
				},
				},
				err: nil,
			},
		},
		{
			testName: "should return error",
			response: mockResponse{
				data: models.Warehouses{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectResult: getAllExpected{
				data: models.Warehouses{},
				err:  customerrors.ErrorInvalidDB,
			},
		},
	}
	for _, test := range testsCases {
		mockRepo := mockRepository.NewRepository(t)
		serv := warehouse.NewService(mockRepo)
		mockRepo.On("GetAll").Return(test.response.data, test.response.err)

		response, err := serv.GetAll()

		assert.Equal(t, test.expectResult.data, response, test.testName)
		assert.Equal(t, test.expectResult.err, err, test.testName)
	}
}

func TestGetById(t *testing.T) {
	type mockResponse struct {
		data models.Warehouse
		err  error
	}
	type getByIdExpected struct {
		data models.Warehouse
		err  error
	}
	type testData struct {
		testName string
		mockResponse
		expectedResult getByIdExpected
		serviceArg     int
	}

	testsCases := []testData{
		{
			testName: "should return warehouse by id",
			mockResponse: mockResponse{
				data: models.Warehouse{ID: 1, Address: "foo"},
				err:  nil,
			},
			expectedResult: getByIdExpected{
				data: models.Warehouse{ID: 1, Address: "foo"},
				err:  nil,
			},
			serviceArg: 1,
		},
		{
			testName: "should return invalid id error",
			mockResponse: mockResponse{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			expectedResult: getByIdExpected{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			serviceArg: 1,
		},
		{
			testName: "should return invalid db error",
			mockResponse: mockResponse{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectedResult: getByIdExpected{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			serviceArg: 1,
		},
	}
	for _, test := range testsCases {
		mockRepo := mockRepository.NewRepository(t)
		serv := warehouse.NewService(mockRepo)
		mockRepo.On("GetByID", mock.Anything).Return(test.mockResponse.data, test.mockResponse.err)

		response, err := serv.GetByID(test.serviceArg)

		assert.Equal(t, test.expectedResult.data, response, test.testName)
		assert.Equal(t, test.expectedResult.err, err, test.testName)
	}
}

func TestCreate(t *testing.T) {
	type mockResponse struct {
		data models.Warehouse
		err  error
	}
	type expectedResult struct {
		data models.Warehouse
		err  error
	}
	type testData struct {
		testName string
		mockResponse
		expectedResult
		serviceArg models.Warehouse
	}

	testsCases := []testData{
		{
			testName: "should return warehouse by id",
			mockResponse: mockResponse{
				data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
				err:  nil,
			},
			serviceArg: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return invalid id error",
			mockResponse: mockResponse{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			serviceArg: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return invalid db error",
			mockResponse: mockResponse{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			serviceArg: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return missing address error",
			mockResponse: mockResponse{
				data: models.Warehouse{},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorMissingAddres,
			},
			serviceArg: models.Warehouse{Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return missing telephone error ",
			mockResponse: mockResponse{
				data: models.Warehouse{},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorMissingTelephone,
			},
			serviceArg: models.Warehouse{Address: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return missing capacity error ",
			mockResponse: mockResponse{
				data: models.Warehouse{},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorMissingCapacity,
			},
			serviceArg: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: -10, MinimunTemperature: 10},
		},
		{
			testName: "should return missing temperature error ",
			mockResponse: mockResponse{
				data: models.Warehouse{},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorMissingTemperature,
			},
			serviceArg: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 0},
		},
	}
	for _, test := range testsCases {
		mockRepo := mockRepository.NewRepository(t)
		serv := warehouse.NewService(mockRepo)

		mockRepo.On("Create", mock.Anything).Return(test.mockResponse.data, test.mockResponse.err).Maybe()

		response, err := serv.Create(test.serviceArg)

		assert.Equal(t, test.expectedResult.data, response, test.testName)
		assert.Equal(t, test.expectedResult.err, err, test.testName)
	}
}

func TestDelete(t *testing.T) {
	type mockResponse struct {
		err error
	}
	type expectedResult struct {
		err error
	}
	type testData struct {
		testName string
		mockResponse
		expectedResult
		serviceArg string
	}

	testsCases := []testData{
		{
			testName: "should return nil",
			mockResponse: mockResponse{
				err: nil,
			},
			expectedResult: expectedResult{
				err: nil,
			},
			serviceArg: "1",
		},
		{
			testName: "should return invalid id error",
			mockResponse: mockResponse{
				err: customerrors.ErrorInvalidID,
			},
			expectedResult: expectedResult{
				err: customerrors.ErrorInvalidID,
			},
			serviceArg: "1",
		},
		{
			testName: "should return invalid db error",
			mockResponse: mockResponse{
				err: customerrors.ErrorInvalidDB,
			},
			expectedResult: expectedResult{
				err: customerrors.ErrorInvalidDB,
			},
			serviceArg: "1",
		},
		{
			testName: "should return syntax error",
			mockResponse: mockResponse{
				err: nil,
			},
			expectedResult: expectedResult{
				err: strconv.ErrSyntax,
			},
			serviceArg: "A",
		},
	}
	for _, test := range testsCases {
		mockRepo := mockRepository.NewRepository(t)
		serv := warehouse.NewService(mockRepo)
		mockRepo.On("Delete", mock.Anything).Return(test.mockResponse.err).Maybe()

		err := serv.Delete(test.serviceArg)
		var conversionError *strconv.NumError;
		if errors.As(err, &conversionError){
			err = conversionError.Err
		}
		
		assert.Equal(t, test.expectedResult.err, err, test.testName)
	}
}
