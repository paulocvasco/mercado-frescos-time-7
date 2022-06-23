package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/models"
	mockService "mercado-frescos-time-7/go-web/internal/products/mock/mockService"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_Find_All(t *testing.T) {
	type responseServiceMock struct {
		data models.Products
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
	type responseWeb struct {
		Code  string          `json:"code"`
		Data  models.Products `json:"data,omitempty"`
		Error string          `json:"error,omitempty"`
	}

	prd1 := models.Product{
		Id:                             1,
		ProductCode:                    "ssd1",
		Description:                    "test 2",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 2,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	prd2 := models.Product{
		Id:                             2,
		ProductCode:                    "ssd2",
		Description:                    "test 2",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 2,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}

	testes := []testData{
		{
			testName: "get all products -  http code 200",
			responseServiceMock: responseServiceMock{
				data: models.Products{
					Products: []models.Product{
						prd1, prd2,
					},
				},
				err: nil,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code: "200",
					Data: models.Products{
						Products: []models.Product{
							prd1, prd2,
						},
					},
				},
				statusCode: 200,
			},
		},
		{
			testName: "return products[] - http code 500",
			responseServiceMock: responseServiceMock{
				data: models.Products{},
				err:  customerrors.ErrorInvalidDB,
			},
			expectResult: expectResult{
				data:       responseWeb{Code: "500", Error: customerrors.ErrorInvalidDB.Error()},
				statusCode: 500,
			},
		},
	}

	for _, test := range testes {
		gin.SetMode(gin.TestMode)

		mockP := mockService.NewService(t)
		ctrl := controller.NewProductHandler(mockP)
		mockP.On("GetAll").Return(test.responseServiceMock.data, test.responseServiceMock.err)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		router.GET("/", ctrl.GetAllProducts())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		responseCode := w.Result().StatusCode

		res := responseWeb{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.expectResult.statusCode, responseCode, test.testName)
		assert.Equal(t, test.expectResult.data, res, test.testName)

	}
}

func TestService_Create_Ok(t *testing.T) {
	type responseServiceMock struct {
		data models.Product
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
		Code  string         `json:"code"`
		Data  models.Product `json:"data,omitempty"`
		Error string         `json:"error,omitempty"`
	}

	prd1 := models.Product{
		Id:                             3,
		ProductCode:                    "ssd3",
		Description:                    "test 2",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 2,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	prd2 := models.Product{
		Id:                             3,
		ExpirationRate:                 2,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	testes := []testData{
		{
			testName: "create product - http code 201",
			responseServiceMock: responseServiceMock{
				data: prd1,
				err:  nil,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code: "201",
					Data: prd1,
				},
				statusCode: 201,
			},
			postData: prd1,
		},
		{
			testName: "product fields error - http code 422",
			responseServiceMock: responseServiceMock{
				data: models.Product{},
			},
			expectResult: expectResult{
				data: responseWeb{
					Code:  "422",
					Data:  models.Product{},
					Error: "validation error in the field(s): product_code, description, width, height, length, net_weight",
				},
				statusCode: 422,
			},
			postData: prd2,
		},
		{
			testName: "create product - http code 409",
			responseServiceMock: responseServiceMock{
				data: models.Product{},
				err:  customerrors.ErrorConflict,
			},
			expectResult: expectResult{
				data: responseWeb{
					Code:  "409",
					Data:  models.Product{},
					Error: "conflict error detected",
				},
				statusCode: 409,
			},
			postData: prd1,
		},
	}

	for _, test := range testes {
		gin.SetMode(gin.TestMode)

		mockP := mockService.NewService(t)
		ctrl := controller.NewProductHandler(mockP)
		mockP.On("Insert", mock.Anything).Return(test.responseServiceMock.data, test.responseServiceMock.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		postData, _ := json.Marshal(test.postData)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(postData))
		router.POST("/", ctrl.SaveProducts())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		responseCode := w.Result().StatusCode

		res := responseWeb{}
		json.Unmarshal(body, &res)

		assert.Equal(t, test.expectResult.statusCode, responseCode, test.testName)
		assert.Equal(t, test.expectResult.data, res, test.testName)

	}
}
