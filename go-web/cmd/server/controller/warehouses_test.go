package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/models"
	mockWarehouse "mercado-frescos-time-7/go-web/internal/warehouse/mock"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllWarehouse(t *testing.T) {
	type webResponse struct {
		Code  string            `json:"code"`
		Data  models.Warehouses `json:"data"`
		Error string            `json:"error"`
	}

	type responseServiceMock struct {
		data models.Warehouses
		err  error
	}
	type expectResult struct {
		response   webResponse
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
				response: webResponse{
					Code: "200",
					Data: models.Warehouses{
						Warehouses: []models.Warehouse{
							{Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20},
							{Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20},
						}},
					Error: "",
				},
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
				response: webResponse{
					Code:  "500",
					Data:  models.Warehouses{},
					Error: customerrors.ErrorInvalidDB.Error(),
				},
				statusCode: 500,
			},
		},
	}
	for _, test := range testCases {
		gin.SetMode(gin.TestMode)

		mockServ := mockWarehouse.NewService(t)
		ctrl := controller.NewControllerWarehouse(mockServ)
		mockServ.On("GetAll").Return(test.data, test.err)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		router.GET("/", ctrl.GetAllWarehouse)
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse
		json.Unmarshal(body, &res)

		assert.Equal(t, test.statusCode, w.Result().StatusCode, test.testName)
		assert.Equal(t, test.response, res, test.testName)

	}
}

func TestGetByIDWarehouse(t *testing.T) {
	type webResponse struct {
		Code  string           `json:"code"`
		Data  models.Warehouse `json:"data"`
		Error string           `json:"error"`
	}

	type responseServiceMock struct {
		data models.Warehouse
		err  error
	}
	type expectResult struct {
		data       webResponse
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
				data: webResponse{
					Code: "200",
					Data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20},
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
				data:       webResponse{Code: "404", Error: customerrors.ErrorInvalidID.Error()},
				statusCode: 404,
			},
			requestedId: "1",
		},
		{
			testName: "should return status 404 - invalid id",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectResult: expectResult{
				data:       webResponse{Code: "500", Error: customerrors.ErrorInvalidDB.Error()},
				statusCode: 500,
			},
			requestedId: "1",
		},
		{
			testName: "should return status 500 - invalid id",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{},
				err:  customerrors.ErrorInvalidID,
			},
			expectResult: expectResult{
				data:       webResponse{Code: "500", Error: "internal error"},
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

		res := webResponse{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.statusCode, w.Result().StatusCode, test.testName)
		assert.Equal(t, test.expectResult.data, res, test.testName)

	}
}

func TestCreateWarehouse(t *testing.T) {
	type webResponse struct {
		Code  string           `json:"code"`
		Data  models.Warehouse `json:"data"`
		Error string           `json:"error"`
	}

	type responseServiceMock struct {
		data models.Warehouse
		err  error
	}
	type expectResult struct {
		data       webResponse
		statusCode int
	}
	type testData struct {
		testName string
		responseServiceMock
		expectResult
		postData interface{}
	}

	dummyTmp := 20
	dummyCap := 20

	testCases := []testData{
		{
			testName:            "should return status 201 and a new warehouse",
			responseServiceMock: responseServiceMock{data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20}, err: nil},
			expectResult: expectResult{
				data: webResponse{
					Code: "201",
					Data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20}},
				statusCode: 201,
			},
			postData: models.PostWarehouse{Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: &dummyTmp, MinimunTemperature: &dummyTmp},
		},
		{
			testName:            "should return status 409",
			responseServiceMock: responseServiceMock{data: models.Warehouse{}, err: customerrors.ErrorConflict},
			expectResult: expectResult{
				data: webResponse{
					Code:  "409",
					Error: customerrors.ErrorConflict.Error()},
				statusCode: 409,
			},
			postData: models.PostWarehouse{Address: "foo", Telephone: "foo", WarehouseCode: "foo", MinimunCapacity: &dummyTmp, MinimunTemperature: &dummyCap},
		},
		{
			testName: "should return status 422 and a validation fields error",
			responseServiceMock: responseServiceMock{
				data: models.Warehouse{},
				err:  nil,
			},
			expectResult: expectResult{
				data:       webResponse{Code: "422", Error: "validation error in the field(s): address, telephone, warehousecode, minimuncapacity, minimuntemperature"},
				statusCode: 422,
			},
			postData: models.PostWarehouse{},
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

		res := webResponse{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.statusCode, w.Result().StatusCode, test.testName)
		assert.Equal(t, test.expectResult.data, res, test.testName)
	}
}

func TestDeleteWarehouse(t *testing.T) {
	type responseServiceMock struct {
		err error
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
			testName:            "should return status 204",
			responseServiceMock: responseServiceMock{err: nil},
			expectResult:        expectResult{statusCode: 204, data: web.Response{}},
			requestedId:         "1",
		},
		{
			testName:            "should return status 404 - invalid id",
			responseServiceMock: responseServiceMock{err: customerrors.ErrorInvalidID},
			expectResult:        expectResult{data: web.Response{Code: "404", Error: customerrors.ErrorInvalidID.Error()}, statusCode: 404},
			requestedId:         "1",
		},
		{
			testName:            "should return status 500 - invalid id",
			responseServiceMock: responseServiceMock{err: strconv.ErrSyntax},
			expectResult:        expectResult{data: web.Response{Code: "500", Error: "internal error"}, statusCode: 500},
			requestedId:         "A",
		},
	}
	for _, test := range testCases {
		gin.SetMode(gin.TestMode)

		mockServ := mockWarehouse.NewService(t)
		ctrl := controller.NewControllerWarehouse(mockServ)
		mockServ.On("Delete", mock.Anything).Return(test.responseServiceMock.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/%v", test.requestedId), nil)
		router.DELETE("/:id", ctrl.DeleteWarehouse)
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := web.Response{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.statusCode, w.Result().StatusCode, test.testName)
		assert.Equal(t, test.data, res, test.testName)

	}
}

func TestUpdateWarehouse(t *testing.T) {
	type webResponse struct {
		Code  string           `json:"code"`
		Data  models.Warehouse `json:"data"`
		Error string           `json:"error"`
	}
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
		patchData   string
		requestedId string
	}
	testCases := []testData{
		{
			testName:            "should return status 200 and a edited warehouse",
			responseServiceMock: responseServiceMock{data: models.Warehouse{ID: 1, Address: "foo-edited", Telephone: "foo-edited", WarehouseCode: "foo-edited", MinimunCapacity: 200, MinimunTemperature: 200}, err: nil},
			expectResult: expectResult{
				data: webResponse{
					Code: "200",
					Data: models.Warehouse{ID: 1, Address: "foo-edited", Telephone: "foo-edited", WarehouseCode: "foo-edited", MinimunCapacity: 200, MinimunTemperature: 200}},
				statusCode: 200},
			patchData:   `{"address": "foo-edited", "telephone": "foo-edited", "warehouse_code": "foo-edited", "minimun_capacity": 200, "minimun_temperature": 200}`,
			requestedId: "/1",
		},
		{
			testName:            "should return status 200 and a edited warehouse",
			responseServiceMock: responseServiceMock{data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo-edited", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20}},
			expectResult: expectResult{
				data: webResponse{
					Code: "200",
					Data: models.Warehouse{ID: 1, Address: "foo", Telephone: "foo-edited", WarehouseCode: "foo", MinimunCapacity: 20, MinimunTemperature: 20}},
				statusCode: 200,
			},
			patchData:   `{"telephone": "foo-edited", "warehouse_code": "foo", "minimun_capacity": 20, "minimun_temperature": 20}`,
			requestedId: "/1",
		},
		{
			testName:            "should return status 409",
			responseServiceMock: responseServiceMock{data: models.Warehouse{}, err: customerrors.ErrorConflict},
			expectResult: expectResult{
				data:       webResponse{Code: "409", Error: customerrors.ErrorConflict.Error()},
				statusCode: 409,
			},
			patchData:   `{"warehouse_code": "foo"}`,
			requestedId: "/1",
		},
		{
			testName:            "should return status 500",
			responseServiceMock: responseServiceMock{data: models.Warehouse{}, err: nil},
			expectResult: expectResult{
				data:       webResponse{Code: "500", Error: "internal error"},
				statusCode: 500,
			},
			patchData:   `{"warehouse_code": "foo"}`,
			requestedId: "/A",
		},
	}
	for _, test := range testCases {
		gin.SetMode(gin.TestMode)

		mockServ := mockWarehouse.NewService(t)
		ctrl := controller.NewControllerWarehouse(mockServ)
		mockServ.On("Update", mock.Anything, mock.Anything).Return(test.responseServiceMock.data, test.responseServiceMock.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodPatch, test.requestedId, strings.NewReader(test.patchData))
		router.PATCH("/:id", ctrl.UpdateWarehouse)

		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := webResponse{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.statusCode, w.Result().StatusCode, test.testName)
		assert.Equal(t, test.expectResult.data, res, test.testName)

	}
}

type mockedReader string

func (m mockedReader) Read(p []byte) (int, error) {
	return 0, errors.New("dummy error")
}

func TestUpdateWarehouseFail(t *testing.T) {
	type webResponse struct {
		Code  string           `json:"code"`
		Data  models.Warehouse `json:"data"`
		Error string           `json:"error"`
	}
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
		patchData   io.Reader
		requestedId string
	}
	testCases := []testData{
		{
			testName:            "erro to read body",
			responseServiceMock: responseServiceMock{},
			expectResult: expectResult{
				data:       webResponse{Code: "500", Error: "internal error"},
				statusCode: 500},
			patchData:   mockedReader("foo"),
			requestedId: "/1",
		},
	}
	for _, test := range testCases {
		gin.SetMode(gin.TestMode)

		mockServ := mockWarehouse.NewService(t)
		ctrl := controller.NewControllerWarehouse(mockServ)
		mockServ.On("Update", mock.Anything, mock.Anything).Return(test.responseServiceMock.data, test.responseServiceMock.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodPatch, test.requestedId, test.patchData)
		router.PATCH("/:id", ctrl.UpdateWarehouse)

		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := webResponse{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.statusCode, w.Result().StatusCode, test.testName)
		assert.Equal(t, test.expectResult.data, res, test.testName)

	}
}
