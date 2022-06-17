package warehouse_test

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/warehouse"
	"mercado-frescos-time-7/go-web/internal/warehouse/mock/mockRepository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
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
				err:  nil,
			},
			expectResult: getAllExpected{
				data: models.Warehouses{},
				err:  nil,
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
		serviceArg int
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
