package controller_test

import (
	"encoding/json"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/models"
	mockWarehouse "mercado-frescos-time-7/go-web/internal/warehouse/mock"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllWarehouse(t *testing.T) {
	type responseServiceMock struct {
		data models.Warehouses
		err  error
	}
	type expectResult struct {
		data       interface{}
		statusCode int
	}
	type testData struct {
		testName string
		responseServiceMock
		expectResult
	}
	testCases := []testData{
		{
			testName: "should return status 200 and list with 2 warehouses",
			responseServiceMock: responseServiceMock{
				data: models.Warehouses{
					Warehouses: []models.Warehouse{
						{Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20},
						{Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20},
					}},
				err: nil,
			},
			expectResult: expectResult{
				data: models.Warehouses{
					Warehouses: []models.Warehouse{
						{Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20},
						{Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20},
					}},
				statusCode: 200,
			},
		},
		{
			testName: "should return status 500",
			responseServiceMock: responseServiceMock{
				data: models.Warehouses{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectResult: expectResult{
				data:       web.Response{Code: "500", Error: customerrors.ErrorInvalidDB.Error()},
				statusCode: 500,
			},
		},
	}
	for _, test := range testCases {
		gin.SetMode(gin.TestMode)

		mockServ := mockWarehouse.NewService(t)
		ctrl := controller.NewControllerWarehouse(mockServ)
		mockServ.On("GetAll").Return(test.responseServiceMock.data, test.responseServiceMock.err)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		router.GET("/", ctrl.GetAllWarehouse)
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		if w.Result().StatusCode < 300 {
			res := models.Warehouses{}
			json.Unmarshal(body, &res)

			assert.Equal(t, test.expectResult.statusCode, w.Result().StatusCode, test.testName)
			assert.Equal(t, test.expectResult.data, res, test.testName)
		} else {
			res := web.Response{}
			json.Unmarshal(body, &res)

			assert.Equal(t, test.expectResult.statusCode, w.Result().StatusCode, test.testName)
			assert.Equal(t, test.expectResult.data, res, test.testName)
		}
	}
}
