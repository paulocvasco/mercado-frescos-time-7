package warehouse_test

import (
	"encoding/json"
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/warehouse"
	mockRepository "mercado-frescos-time-7/go-web/internal/warehouse/mock"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	type mockGetAllResponse struct {
		data models.Warehouses
		err  error
	}
	type getAllExpected struct {
		data models.Warehouses
		err  error
	}
	type testData struct {
		testName     string
		response     mockGetAllResponse
		expectResult getAllExpected
	}

	testsCases := []testData{
		{
			testName: "should return all warehouses",
			response: mockGetAllResponse{
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
			response: mockGetAllResponse{
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
	type mockGetByIdResponse struct {
		data models.Warehouse
		err  error
	}
	type getByIdExpected struct {
		data models.Warehouse
		err  error
	}
	type testData struct {
		testName string
		mockGetByIdResponse
		expectedResult getByIdExpected
		argRequestId   int
	}

	testsCases := []testData{
		{
			testName: "should return warehouse by id",
			mockGetByIdResponse: mockGetByIdResponse{
				data: models.Warehouse{ID: 1, Address: "foo"},
				err:  nil,
			},
			expectedResult: getByIdExpected{
				data: models.Warehouse{ID: 1, Address: "foo"},
				err:  nil,
			},
			argRequestId: 1,
		},
		{
			testName: "should return invalid id error",
			mockGetByIdResponse: mockGetByIdResponse{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			expectedResult: getByIdExpected{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			argRequestId: 1,
		},
		{
			testName: "should return invalid db error",
			mockGetByIdResponse: mockGetByIdResponse{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectedResult: getByIdExpected{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			argRequestId: 1,
		},
	}
	for _, test := range testsCases {
		mockRepo := mockRepository.NewRepository(t)
		serv := warehouse.NewService(mockRepo)
		mockRepo.On("GetByID", mock.Anything).Return(test.mockGetByIdResponse.data, test.mockGetByIdResponse.err)

		response, err := serv.GetByID(test.argRequestId)

		assert.Equal(t, test.expectedResult.data, response, test.testName)
		assert.Equal(t, test.expectedResult.err, err, test.testName)
	}
}

func TestCreate(t *testing.T) {
	type mockCreateResponse struct {
		data models.Warehouse
		err  error
	}
	type expectedResult struct {
		data models.Warehouse
		err  error
	}
	type testData struct {
		testName string
		mockCreateResponse
		expectedResult
		argModelCreate models.Warehouse
	}

	testsCases := []testData{
		{
			testName: "should return warehouse by id",
			mockCreateResponse: mockCreateResponse{
				data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
				err:  nil,
			},
			argModelCreate: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return invalid id error",
			mockCreateResponse: mockCreateResponse{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			argModelCreate: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return invalid db error",
			mockCreateResponse: mockCreateResponse{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			argModelCreate: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return missing address error",
			mockCreateResponse: mockCreateResponse{
				data: models.Warehouse{},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorMissingAddres,
			},
			argModelCreate: models.Warehouse{Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return missing telephone error ",
			mockCreateResponse: mockCreateResponse{
				data: models.Warehouse{},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorMissingTelephone,
			},
			argModelCreate: models.Warehouse{Address: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return missing capacity error ",
			mockCreateResponse: mockCreateResponse{
				data: models.Warehouse{},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorMissingCapacity,
			},
			argModelCreate: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: -10, MinimunTemperature: 10},
		},
		{
			testName: "should return missing temperature error ",
			mockCreateResponse: mockCreateResponse{
				data: models.Warehouse{},
				err:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorMissingTemperature,
			},
			argModelCreate: models.Warehouse{Address: "foo", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 0},
		},
	}
	for _, test := range testsCases {
		mockRepo := mockRepository.NewRepository(t)
		serv := warehouse.NewService(mockRepo)

		mockRepo.On("Create", mock.Anything).Return(test.mockCreateResponse.data, test.mockCreateResponse.err).Maybe()

		response, err := serv.Create(test.argModelCreate)

		assert.Equal(t, test.expectedResult.data, response, test.testName)
		assert.Equal(t, test.expectedResult.err, err, test.testName)
	}
}

func TestDelete(t *testing.T) {
	type mockDeleteResponse struct {
		err error
	}
	type expectedResult struct {
		err error
	}
	type testData struct {
		testName string
		mockDeleteResponse
		expectedResult
		serviceArg string
	}

	testsCases := []testData{
		{
			testName: "should return nil",
			mockDeleteResponse: mockDeleteResponse{
				err: nil,
			},
			expectedResult: expectedResult{
				err: nil,
			},
			serviceArg: "1",
		},
		{
			testName: "should return invalid id error",
			mockDeleteResponse: mockDeleteResponse{
				err: customerrors.ErrorInvalidID,
			},
			expectedResult: expectedResult{
				err: customerrors.ErrorInvalidID,
			},
			serviceArg: "1",
		},
		{
			testName: "should return invalid db error",
			mockDeleteResponse: mockDeleteResponse{
				err: customerrors.ErrorInvalidDB,
			},
			expectedResult: expectedResult{
				err: customerrors.ErrorInvalidDB,
			},
			serviceArg: "1",
		},
		{
			testName: "should return syntax error",
			mockDeleteResponse: mockDeleteResponse{
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
		mockRepo.On("Delete", mock.Anything).Return(test.mockDeleteResponse.err).Maybe()

		err := serv.Delete(test.serviceArg)
		var conversionError *strconv.NumError
		if errors.As(err, &conversionError) {
			err = conversionError.Err
		}

		assert.Equal(t, test.expectedResult.err, err, test.testName)
	}
}

func TestUpdate(t *testing.T) {
	type mockResponse struct {
		dataGetById models.Warehouse
		dataUpdate  models.Warehouse
		errGetById  error
		errUpdate   error
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
			testName: "should return updated warehouse",
			mockResponse: mockResponse{
				dataGetById: models.Warehouse{ID: 1, Address: "foo", WarehouseCode: "CODE", Telephone: "foo", MinimunCapacity: 10, MinimunTemperature: 10},
				dataUpdate:  models.Warehouse{ID: 1, Address: "foobar", Telephone: "foobar", MinimunCapacity: 10, MinimunTemperature: 10},
				errGetById:  nil,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{ID: 1, Address: "foobar", WarehouseCode: "CODE", Telephone: "foobar", MinimunCapacity: 10, MinimunTemperature: 10},
				err:  nil,
			},
			serviceArg: models.Warehouse{ID: 1, Address: "foobar", WarehouseCode: "CODE", Telephone: "foobar", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return invalid error id",
			mockResponse: mockResponse{
				dataGetById: models.Warehouse{},
				dataUpdate:  models.Warehouse{},
				errGetById:  customerrors.ErrorInvalidID,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			serviceArg: models.Warehouse{ID: 1, Address: "foobar", WarehouseCode: "CODE", Telephone: "foobar", MinimunCapacity: 10, MinimunTemperature: 10},
		},
		{
			testName: "should return invalid error db",
			mockResponse: mockResponse{
				dataGetById: models.Warehouse{},
				dataUpdate:  models.Warehouse{},
				errGetById:  nil,
				errUpdate:   customerrors.ErrorInvalidDB,
			},
			expectedResult: expectedResult{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			serviceArg: models.Warehouse{ID: 1, Address: "foobar", WarehouseCode: "CODE", Telephone: "foobar", MinimunCapacity: 10, MinimunTemperature: 10},
		},
	}
	for _, test := range testsCases {
		mockRepo := mockRepository.NewRepository(t)
		serv := warehouse.NewService(mockRepo)

		mockRepo.On("GetByID", mock.Anything).Return(test.mockResponse.dataGetById, test.mockResponse.errGetById).Maybe()
		mockRepo.On("Update", mock.Anything, mock.Anything).Return(test.mockResponse.errUpdate).Maybe()

		whBytes, _ := json.Marshal(test.serviceArg)
		response, err := serv.Update(test.serviceArg.ID, whBytes)

		assert.Equal(t, test.expectedResult.data, response, test.testName)
		assert.Equal(t, test.expectedResult.err, err, test.testName)
	}
}
