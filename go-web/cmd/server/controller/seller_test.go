package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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
		models.Seller{ID: 1, Cid: 0, Company_name: "", Address: "", Telephone: ""},
		204,
		"1",
	}

	testsCases := []tests{
		{"Delete", response, response, nil, "Error Delete"},
		{"Delete Error", responseController{statusCode: 500, idRequest: "Error"}, responseController{statusCode: 500, idRequest: "Error"}, customerrors.ErrorInvalidID, "Error Delete"},
		{"Delete Error", responseController{statusCode: 404, idRequest: "1"}, responseController{statusCode: 404, idRequest: "1"}, customerrors.ErrorInvalidID, "Error Delete"},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("Delete", mock.Anything).Return(value.expectError).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/%v", value.mockResponse.idRequest), nil)
		router.DELETE("/:id", control.SellersDelete())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := value.expectResponse.data
		json.Unmarshal(body, &res)
		log.Println(res)
		// log.Println(w.Result().StatusCode, "value: ", value.expectResponse.statusCode)

		assert.Equal(t, value.expectResponse.data, res, value.message)
		assert.Equal(t, value.expectResponse.statusCode, w.Result().StatusCode, value.message)

	}

}

func TestUpdate(t *testing.T) {
	type responseController struct {
		data       models.Seller
		statusCode int
		idRequest  string
	}
	type update struct {
		Company_name string
		Address      string
		Telephone    string
	}
	sendUpdate := update{"Mercado Livre", "Rua1", "(11) 3334-5564"}

	type tests struct {
		name           string
		mockResponse   responseController
		expectResponse responseController
		expectError    error
		message        string
		sendPatch      update
	}

	response := responseController{
		models.Seller{ID: 1, Cid: 123, Company_name: "Mercalo Livre", Address: "Rua1", Telephone: "(11) 3334-5564"},
		200,
		"1",
	}

	testsCases := []tests{
		{"Update", response, response, nil, "Error GetId", sendUpdate},
		{"Update Error", responseController{statusCode: 500, idRequest: "Error"}, responseController{statusCode: 500, idRequest: "Error"}, customerrors.ErrorInvalidID, "Error Delete", sendUpdate},
		{"Update Error", responseController{statusCode: 404, idRequest: "1"}, responseController{statusCode: 404, idRequest: "1"}, customerrors.ErrorInvalidID, "Error Delete", update{}},
		{"Update Error", responseController{statusCode: 409, idRequest: "1"}, responseController{statusCode: 409, idRequest: "1"}, customerrors.ErrorConflict, "Error Delete", update{}},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("Update", mock.Anything, mock.Anything).Return(value.mockResponse.data, value.expectError).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		patchData, _ := json.Marshal(value.sendPatch)

		req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/%v", value.mockResponse.idRequest), bytes.NewReader(patchData))
		router.PATCH("/:id", control.SellersUpdate())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := value.expectResponse.data
		json.Unmarshal(body, &res)

		assert.Equal(t, value.expectResponse.data, res, value.message)
		assert.Equal(t, value.expectResponse.statusCode, w.Result().StatusCode, value.message)

	}

}

func TestStore(t *testing.T) {
	type responseController struct {
		data       models.Seller
		statusCode int
	}

	type create struct {
		Cid          int
		Company_name string
		Address      string
		Telephone    string
	}
	type tests struct {
		name           string
		mockResponse   responseController
		expectResponse responseController
		expectError    error
		message        string
		sendPost       create
	}

	sendCreate := create{123, "Mercado Livre", "Rua1", "(11) 3334-5564"}

	response := responseController{
		models.Seller{ID: 1, Cid: 123, Company_name: "Mercalo Livre", Address: "Rua1", Telephone: "(11) 3334-5564"},
		201,
	}

	testsCases := []tests{
		{"Store", response, response, nil, "Error Store", sendCreate},
		{"Store Error", responseController{statusCode: 422}, responseController{statusCode: 422}, customerrors.ErrorInvalidID, "Error Delete status 422", create{}},
		{"Store Error", responseController{statusCode: 409}, responseController{statusCode: 409}, customerrors.ErrorConflict, "Error Delete status 409", sendCreate},
		{"Store Error", responseController{statusCode: 404}, responseController{statusCode: 404}, customerrors.ErrorInvalidID, "Error Delete status 404", sendCreate},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("Store", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(value.mockResponse.data, value.expectError).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)
		post, _ := json.Marshal(value.sendPost)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(post))
		router.POST("/", control.SellersStore())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		res := value.expectResponse.data
		json.Unmarshal(body, &res)
		log.Println(res, value.expectResponse.data)
		assert.Equal(t, value.expectResponse.data, res, value.message)
		assert.Equal(t, value.expectResponse.statusCode, w.Result().StatusCode, value.message)

	}

}
