package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/seller/mocks"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	type mockResponse struct {
		data []models.Seller
		err  error
	}

	type getAllResponse struct {
		Seller []models.Seller `json:"sellers"`
	}

	type webResponse struct {
		Code  string         `json:"code"`
		Data  getAllResponse `json:"data"`
		Error string         `json:"error"`
	}

	type tests struct {
		name string
		mockResponse
		expectResponse   webResponse
		expectStatusCode int
	}

	mr := mockResponse{
		[]models.Seller{
			{ID: 1, Cid: 123, Company_name: "Meli1", Address: "Rua 1", Telephone: "(11) 33387767"},
			{ID: 2, Cid: 1234, Company_name: "Meli2", Address: "Rua 2", Telephone: "(11) 33387768"},
			{ID: 3, Cid: 12345, Company_name: "Meli3", Address: "Rua 3", Telephone: "(11) 33387768"},
		}, nil}

	testsCases := []tests{
		{"GetAll", mr, webResponse{Code: "200", Data: getAllResponse{mr.data}, Error: ""}, http.StatusOK},
		{"GetAll Error", mockResponse{nil, customerrors.ErrorInvalidDB}, webResponse{Code: "500", Data: getAllResponse{}, Error: customerrors.ErrorInvalidDB.Error()}, http.StatusInternalServerError},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("GetAll").Return(value.mockResponse.data, value.mockResponse.err)

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		router.GET("/", control.SellersGetAll())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse

		json.Unmarshal(body, &res)
		assert.Equal(t, value.expectResponse, res, value.name)
		assert.Equal(t, value.expectStatusCode, w.Result().StatusCode, value.name)

	}

}

func TestGetID(t *testing.T) {

	type webResponse struct {
		Code  string        `json:"code"`
		Data  models.Seller `json:"data"`
		Error string        `json:"error"`
	}

	type serviceResponse struct {
		data models.Seller
		err  error
	}

	type tests struct {
		name               string
		id                 int
		mockResponse       serviceResponse
		expectResponse     webResponse
		expectedstatuscode int
		message            string
	}

	resp := serviceResponse{
		models.Seller{ID: 1, Cid: 123, Company_name: "Meli1", Address: "Rua 1", Telephone: "(11) 33387767"},
		nil,
	}

	testsCases := []tests{
		{"GetId", 1, resp, webResponse{Code: "200", Data: resp.data}, http.StatusOK, "Ok GetId"},
		{"GetId Error", 2, serviceResponse{resp.data, customerrors.ErrorInvalidDB}, webResponse{Code: "500", Data: models.Seller{}, Error: customerrors.ErrorInvalidDB.Error()}, http.StatusInternalServerError, "InternalError GetId"},
		{"GetId Error", 3, serviceResponse{resp.data, customerrors.ErrorInvalidID}, webResponse{Code: "404", Data: models.Seller{}, Error: customerrors.ErrorInvalidID.Error()}, http.StatusNotFound, "InternalError GetId"},
		//{"GetId Error", responseController{statusCode: 400, idRequest: "Error"}, responseController{statusCode: 400, idRequest: "Error"}, customerrors.ErrorInvalidID, "Error GetId status 400"},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("GetId", mock.Anything).Return(value.mockResponse.data, value.mockResponse.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%v", value.id), nil)
		router.GET("/:id", control.SellersGetId())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse
		json.Unmarshal(body, &res)
		assert.Equal(t, value.expectResponse, res, value.message)
		assert.Equal(t, value.expectedstatuscode, w.Result().StatusCode, value.message)

	}

}

func TestDelete(t *testing.T) {

	type webResponse struct {
		Code  string        `json:"code"`
		Error string        `json:"error"`
	}

	type serviceResponse struct {
		err  error
	}

	type tests struct {
		name               string
		id                 string
		mockResponse       serviceResponse
		expectResponse     webResponse
		expectedstatuscode int
		message            string
	}

	resp := serviceResponse{
		nil,
	}

	testsCases := []tests{
		{"Delete", "1", resp, webResponse{Code: "204", Error: "delete erro"}, http.StatusNoContent, "Error Delete"},
		//{"Delete Error", responseController{statusCode: 500, idRequest: "Error"}, responseController{statusCode: 500, idRequest: "Error"}, customerrors.ErrorInvalidID, "Error Delete"},
		//{"Delete Error", responseController{statusCode: 404, idRequest: "1"}, responseController{statusCode: 404, idRequest: "1"}, customerrors.ErrorInvalidID, "Error Delete"},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("Delete", mock.Anything).Return(value.mockResponse.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)
		router.DELETE("/:id", control.SellersDelete())
		req := httptest.NewRequest(http.MethodDelete, "/"+value.id, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, value.expectedstatuscode, w.Result().StatusCode, value.message)

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

		assert.Equal(t, value.expectResponse.data, res, value.message)
		assert.Equal(t, value.expectResponse.statusCode, w.Result().StatusCode, value.message)

	}

}
