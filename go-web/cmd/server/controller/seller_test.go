package controller_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/Seller/mocks"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	type responseController struct {
		data       []models.Seller
		statusCode int
	}

	type tests struct {
		name           string
		mockResponse   responseController
		expectResponse responseController
		expectError    error
		message        string
	}

	response := responseController{
		[]models.Seller{
			{ID: 1, Cid: 123, Company_name: "Meli1", Address: "Rua 1", Telephone: "(11) 33387767"},
			{ID: 2, Cid: 1234, Company_name: "Meli2", Address: "Rua 2", Telephone: "(11) 33387768"},
			{ID: 3, Cid: 12345, Company_name: "Meli3", Address: "Rua 3", Telephone: "(11) 33387768"},
		}, 200}

	testsCases := []tests{
		{"GetAll", response, response, nil, "Error GetAll"},
		{"GetAll Error", responseController{statusCode: 400}, responseController{statusCode: 400}, errors.New("Error"), "Error GetAll"},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("GetAll").Return(value.mockResponse.data, value.expectError)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		router.GET("/", control.SellersGetAll())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := value.expectResponse.data

		json.Unmarshal(body, &res)
		assert.Equal(t, value.expectResponse.data, res, value.message)
		assert.Equal(t, value.expectResponse.statusCode, w.Result().StatusCode, value.message)

	}

}
func TestGetID(t *testing.T) {
	type responseController struct {
		data       models.Seller
		statusCode int
		idRequest  string
	}

	type tests struct {
		name           string
		mockResponse   responseController
		expectResponse responseController
		expectError    error
		message        string
	}

	response := responseController{
		models.Seller{ID: 1, Cid: 123, Company_name: "Meli1", Address: "Rua 1", Telephone: "(11) 33387767"},
		200,
		"1",
	}

	testsCases := []tests{
		{"GetId", response, response, nil, "Error GetId"},
		{"GetId Error", responseController{statusCode: 500, idRequest: "Error"}, responseController{statusCode: 500, idRequest: "Error"}, customerrors.ErrorInvalidID, "Error GetId"},
		{"GetId Error", responseController{statusCode: 500, idRequest: "Error"}, responseController{statusCode: 500, idRequest: "Error"}, errors.New("Error"), "Error GetId"},
		{"GetId Error", responseController{statusCode: 404, idRequest: "1"}, responseController{statusCode: 404, idRequest: "1"}, customerrors.ErrorInvalidID, "Error GetId"},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("GetId", mock.Anything).Return(value.mockResponse.data, value.expectError).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%v", value.mockResponse.idRequest), nil)
		router.GET("/:id", control.SellersGetId())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := value.expectResponse.data
		json.Unmarshal(body, &res)
		if value.expectError != customerrors.ErrorInvalidID && value.expectError != nil {
			status, msg := customerrors.ErrorHandleResponse(value.expectError)
			assert.Equal(t, value.expectResponse.data, res, value.message)
			assert.Equal(t, status, w.Result().StatusCode, msg)

		} else {
			assert.Equal(t, value.expectResponse.data, res, value.message)
			assert.Equal(t, value.expectResponse.statusCode, w.Result().StatusCode, value.message)
		}

	}

}

func TestDelete(t *testing.T) {
	type responseController struct {
		data       models.Seller
		statusCode int
		idRequest  string
	}

	type tests struct {
		name           string
		mockResponse   responseController
		expectResponse responseController
		expectError    error
		message        string
	}

	response := responseController{
		models.Seller{ID: 1, Cid: 123, Company_name: "Meli1", Address: "Rua 1", Telephone: "(11) 33387767"},
		200,
		"1",
	}

	testsCases := []tests{
		{"GetId", response, response, nil, "Error GetId"},
		{"GetId Error", responseController{statusCode: 500, idRequest: "Error"}, responseController{statusCode: 500, idRequest: "Error"}, customerrors.ErrorInvalidID, "Error GetId"},
		{"GetId Error", responseController{statusCode: 500, idRequest: "Error"}, responseController{statusCode: 500, idRequest: "Error"}, errors.New("Error"), "Error GetId"},
		{"GetId Error", responseController{statusCode: 404, idRequest: "1"}, responseController{statusCode: 404, idRequest: "1"}, customerrors.ErrorInvalidID, "Error GetId"},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("Delete", mock.Anything).Return(value.mockResponse.data, value.expectError).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%v", value.mockResponse.idRequest), nil)
		router.GET("/:id", control.SellersDelete())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := value.expectResponse.data
		json.Unmarshal(body, &res)
		if value.expectError != customerrors.ErrorInvalidID && value.expectError != nil {
			status, msg := customerrors.ErrorHandleResponse(value.expectError)
			assert.Equal(t, value.expectResponse.data, res, value.message)
			assert.Equal(t, status, w.Result().StatusCode, msg)

		} else {
			assert.Equal(t, value.expectResponse.data, res, value.message)
			assert.Equal(t, value.expectResponse.statusCode, w.Result().StatusCode, value.message)
		}

	}

}
