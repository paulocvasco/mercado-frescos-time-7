package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	"github.com/stretchr/testify/mock"
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

func TestGetByIDWarehouse(t *testing.T) {
	type responseServiceMock struct {
		data models.Warehouse
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
		requestedId string
	}
	testCases := []testData{
		{
			testName: "should return status 200 and a warehouse with correct id",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{
					ID: 1, Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20,
				},
				err: nil,
			},
			expectResult: expectResult{
				data: models.Warehouse{
					ID: 1, Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20,
				},
				statusCode: 200,
			},
			requestedId: "1",
		},
		{
			testName: "should return status 404 - invalid id",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			expectResult: expectResult{
				data:       web.Response{Code: "404", Error: customerrors.ErrorInvalidID.Error()},
				statusCode: 404,
			},
			requestedId: "1",
		},
		{
			testName: "should return status 500 - invalid id - ALTERAR ESSE ERRO DPS",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			expectResult: expectResult{
				data:       web.Response{Code: "500", Error: "internal error"},
				statusCode: 500,
			},
			requestedId: "A",
		},
	}
	for _, test := range testCases {
		gin.SetMode(gin.TestMode)

		mockServ := mockWarehouse.NewService(t)
		ctrl := controller.NewControllerWarehouse(mockServ)
		mockServ.On("GetByID", mock.Anything).Return(test.responseServiceMock.data, test.responseServiceMock.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%v", test.requestedId), nil)
		router.GET("/:id", ctrl.GetByIDWarehouse)
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		if w.Result().StatusCode < 300 {
			res := models.Warehouse{}
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

func TestCreateWarehouse(t *testing.T) {
	type responseServiceMock struct {
		data models.Warehouse
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
		postData interface{}
	}
	testCases := []testData{
		{
			testName: "should return status 201 and a new warehouse",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{
					ID: 1, Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20,
				},
				err: nil,
			},
			expectResult: expectResult{
				data: models.Warehouse{
					ID: 1, Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20,
				},
				statusCode: 201,
			},
			postData: models.Warehouse{
				Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20,
			},
		},
		{
			testName: "should return status 409",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{},
				err:  customerrors.ErrorConflict,
			},
			expectResult: expectResult{
				data:       web.Response{Code: "409", Error: customerrors.ErrorConflict.Error()},
				statusCode: 409,
			},
			postData: models.Warehouse{
				Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20,
			},
		},
		{
			testName: "should return status 422 and a validation fields error",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{},
				err:  nil,
			},
			expectResult: expectResult{
				data:       web.Response{Code: "422", Error: "validation error in the field(s): address, telephone, minimuncapacity, minimuntemperature"},
				statusCode: 422,
			},
			postData: models.Warehouse{},
		},
	}
	for _, test := range testCases {
		gin.SetMode(gin.TestMode)

		mockServ := mockWarehouse.NewService(t)
		ctrl := controller.NewControllerWarehouse(mockServ)
		mockServ.On("Create", mock.Anything).Return(test.responseServiceMock.data, test.responseServiceMock.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		postData, _ := json.Marshal(test.postData)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(postData))
		router.POST("/", ctrl.CreateWarehouse)
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		if w.Result().StatusCode < 300 {
			res := models.Warehouse{}
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
