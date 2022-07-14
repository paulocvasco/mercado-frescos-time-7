package controller_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/models"
	mockService "mercado-frescos-time-7/go-web/internal/product_records/mock/mockService"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreateProductRecord(t *testing.T) {
	type responseServiceMock struct {
		data models.ProductRecord
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
	type responseWeb struct {
		Code  string               `json:"code"`
		Data  models.ProductRecord `json:"data,omitempty"`
		Error string               `json:"error,omitempty"`
	}

	record1 := controller.ProductRecordInsert{
		LastUpdateDate: "2022-07-13",
		PurchasePrice:  10.10,
		SalePrice:      20.20,
		ProductId:      1,
	}
	record2 := controller.ProductRecordInsert{
		ProductId:      1,
		LastUpdateDate: "2006-01-40",
		PurchasePrice:  10.10,
		SalePrice:      20.20,
	}
	time, _ := time.Parse("2006-01-02", "2022-07-13")
	recordResponse := models.ProductRecord{
		Id:             1,
		ProductId:      1,
		LastUpdateDate: time,
		PurchasePrice:  10.10,
		SalePrice:      20.20,
	}
	testes := []testData{
		{
			testName: "create record - http code 201",
			responseServiceMock: responseServiceMock{
				data: recordResponse,
				err:  nil,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code: "201",
					Data: recordResponse,
				},
				statusCode: 201,
			},
			postData: record1,
		},
		{
			testName: "create record service error date- http code 422",
			responseServiceMock: responseServiceMock{
				data: models.ProductRecord{},
				err:  customerrors.ErrorInvalidDate,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code:  "422",
					Data:  models.ProductRecord{},
					Error: customerrors.ErrorInvalidDate.Error(),
				},
				statusCode: 422,
			},
			postData: record1,
		},
		{
			testName: "create record error date - http code 400",
			responseServiceMock: responseServiceMock{
				data: recordResponse,
				err:  nil,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code:  "422",
					Data:  models.ProductRecord{},
					Error: customerrors.ErrorInvalidDate.Error(),
				},
				statusCode: 422,
			},
			postData: record2,
		},
		{
			testName: "product fields error - http code 422",
			responseServiceMock: responseServiceMock{
				data: models.ProductRecord{},
				err:  nil,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code:  "422",
					Data:  models.ProductRecord{},
					Error: "validation error in the field(s): lastupdatedate, purchaseprice, saleprice, productid",
				},
				statusCode: 422,
			},
			postData: controller.ProductRecordInsert{},
		},
	}

	for _, test := range testes {
		gin.SetMode(gin.TestMode)

		mockP := mockService.NewService(t)
		ctrl := controller.NewProductRecordsController(mockP)
		mockP.On("Insert", mock.Anything).Return(test.responseServiceMock.data, test.responseServiceMock.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		postData, _ := json.Marshal(test.postData)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(postData))
		router.POST("/", ctrl.InsertProductRecord())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		responseCode := w.Result().StatusCode

		res := responseWeb{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.expectResult.statusCode, responseCode, test.testName)
		assert.Equal(t, test.expectResult.data, res, test.testName)

	}
}

func TestServiceGetProductRecord(t *testing.T) {
	type responseServiceMock struct {
		data models.ProductsRecordsResponse
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
		idUrlRequest string
	}
	type responseWeb struct {
		Code  string                         `json:"code"`
		Data  models.ProductsRecordsResponse `json:"data,omitempty"`
		Error string                         `json:"error,omitempty"`
	}
	recordsResponse := models.ProductsRecordsResponse{Records: []models.ProductRecordsResponse{
		{
			ProductId:    1,
			Description:  "test",
			RecordsCount: 2,
		},
	}}
	testes := []testData{
		{
			testName: "return records - http code 200",
			responseServiceMock: responseServiceMock{
				data: recordsResponse,
				err:  nil,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code: "200",
					Data: recordsResponse,
				},
				statusCode: 200,
			},
			idUrlRequest: "1",
		},
		{
			testName: "return error invalid param - http code 400",
			responseServiceMock: responseServiceMock{
				data: models.ProductsRecordsResponse{},
				err:  customerrors.ErrorInvalidDate,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code:  "400",
					Data:  models.ProductsRecordsResponse{},
					Error: "input param: a must be an integer",
				},
				statusCode: 400,
			},
			idUrlRequest: "a",
		},
		{
			testName: "return error invalid db - http code 500",
			responseServiceMock: responseServiceMock{
				data: models.ProductsRecordsResponse{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code:  "500",
					Data:  models.ProductsRecordsResponse{},
					Error: customerrors.ErrorInvalidDB.Error(),
				},
				statusCode: 500,
			},
			idUrlRequest: "0",
		},
	}

	for _, test := range testes {
		gin.SetMode(gin.TestMode)

		mockP := mockService.NewService(t)
		ctrl := controller.NewProductRecordsController(mockP)
		mockP.On("GetProductRecords", mock.Anything).Return(test.responseServiceMock.data, test.responseServiceMock.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/?id="+test.idUrlRequest, nil)
		router.GET("/", ctrl.GetProductRecordsById())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		responseCode := w.Result().StatusCode

		res := responseWeb{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.expectResult.statusCode, responseCode, test.testName)
		assert.Equal(t, test.expectResult.data, res, test.testName)

	}
}
