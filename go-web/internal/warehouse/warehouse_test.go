package warehouse_test

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/warehouse"
	"mercado-frescos-time-7/go-web/internal/warehouse/mock/mockRepository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	type getAllResponse struct {
		data models.Warehouses
		err  error
	}
	type getAllExpected struct {
		data models.Warehouses
		err  error
	}
	type testData struct {
		testName     string
		response     getAllResponse
		expectResult getAllExpected
	}

	testsCases := []testData{
		{
			testName: "should return all warehouses",
			response: getAllResponse{
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
			response: getAllResponse{
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
