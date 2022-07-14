package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/seller/mocks"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"io/ioutil"
	"log"
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
		id                 string
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
		{"GetId", "1", resp, webResponse{Code: "200", Data: resp.data}, http.StatusOK, "Ok GetId"},
		{"GetId Error", "2", serviceResponse{resp.data, customerrors.ErrorInvalidDB}, webResponse{Code: "500", Data: models.Seller{}, Error: customerrors.ErrorInvalidDB.Error()}, http.StatusInternalServerError, "InternalError GetId"},
		{"GetId Error", "3", serviceResponse{resp.data, customerrors.ErrorInvalidID}, webResponse{Code: "404", Data: models.Seller{}, Error: customerrors.ErrorInvalidID.Error()}, http.StatusNotFound, "InternalError GetId"},
		{"GetId Error", "A", serviceResponse{resp.data, customerrors.ErrorInvalidIDParameter}, webResponse{Code: "400", Data: models.Seller{}, Error: "input param: A must be an integer"}, http.StatusBadRequest, "InternalError GetId"},
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
		Code  string `json:"code"`
		Error string `json:"error"`
	}

	type serviceResponse struct {
		err error
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
		{"Delete", "1", resp, webResponse{Code: "204", Error: ""}, http.StatusNoContent, "Error Delete"},
		{"Delete Error", "1", serviceResponse{customerrors.ErrorInvalidDB}, webResponse{Code: "500", Error: customerrors.ErrorInvalidDB.Error()}, http.StatusInternalServerError, "Error Delete"},
		{"Delete Error", "1", serviceResponse{customerrors.ErrorInvalidID}, webResponse{Code: "404", Error: customerrors.ErrorInvalidID.Error()}, http.StatusNotFound, "Error Delete Not Found"},
		{"Delete Error", "A", serviceResponse{customerrors.ErrorInvalidIDParameter}, webResponse{Code: "400", Error: "input param: A must be an integer"}, http.StatusBadRequest, "InternalError GetId"},
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
	type webResponse struct {
		Code  string        `json:"code"`
		Data  models.Seller `json:"data"`
		Error string        `json:"error"`
	}

	type mockResponse struct {
		data models.Seller
		err  error
	}

	type update struct {
		Company_name string
		Address      string
		Telephone    string
	}
	sendUpdate := update{"Mercado Livre", "Rua1", "(11) 3334-5564"}

	type tests struct {
		name string
		mockResponse
		expectResponse     webResponse
		message            string
		sendPatch          update
		idRequest          string
		expectedstatuscode int
	}

	response := models.Seller{
		ID: 1, Cid: 123, Company_name: "Mercalo Livre", Address: "Rua1", Telephone: "(11) 3334-5564",
	}

	testsCases := []tests{
		{"Update", mockResponse{response, nil}, webResponse{"200", response, ""}, "Error GetId", sendUpdate, "1", 200},
		{"Update Error", mockResponse{models.Seller{}, customerrors.ErrorInvalidIDParameter}, webResponse{"400", models.Seller{}, "input param: A must be an integer"}, "Error GetId", sendUpdate, "A", 400},
		{"Update Error", mockResponse{models.Seller{}, customerrors.ErrorInvalidID}, webResponse{"404", models.Seller{}, customerrors.ErrorInvalidID.Error()}, "Error GetId", sendUpdate, "1", 404},
		{"Update Error", mockResponse{models.Seller{}, customerrors.ErrorConflict}, webResponse{"409", models.Seller{}, customerrors.ErrorConflict.Error()}, "Error GetId", sendUpdate, "1", 409},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("Update", mock.Anything, mock.Anything).Return(value.mockResponse.data, value.mockResponse.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		patchData, _ := json.Marshal(value.sendPatch)

		req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/%v", value.idRequest), bytes.NewReader(patchData))
		router.PATCH("/:id", control.SellersUpdate())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse
		json.Unmarshal(body, &res)

		assert.Equal(t, value.expectResponse, res, value.message)
		assert.Equal(t, value.expectedstatuscode, w.Result().StatusCode, value.message)

	}

}

func TestStore(t *testing.T) {
	type webResponse struct {
		Code  string        `json:"code"`
		Data  models.Seller `json:"data"`
		Error string        `json:"error"`
	}

	type mockResponse struct {
		data models.Seller
		err  error
	}

	type create struct {
		Cid          int    `json:"cid"`
		Company_name string `json:"company_name"`
		Address      string `json:"address"`
		Telephone    string `json:"telephone"`
		LocalityID   string `json:"locality_id"`
	}

	type tests struct {
		name string
		mockResponse
		expectResponse     webResponse
		message            string
		sendCreate         create
		idRequest          string
		expectedstatuscode int
	}
	sendCreate := create{123, "Mercado Livre", "Rua1", "(11) 3334-5564", "1"}

	response := models.Seller{ID: 1, Cid: 123, Company_name: "Mercalo Livre", Address: "Rua1", Telephone: "(11) 3334-5564", LocalityID: "1"}

	testsCases := []tests{
		{"Store", mockResponse{response, nil}, webResponse{"201", response, ""}, "Store", sendCreate, "1", 201},
		{"Store Error", mockResponse{models.Seller{}, customerrors.ErrorInvalidIDParameter}, webResponse{"422", models.Seller{}, "validation error in the field(s): cid, companyname, address, telephone, localityid"}, "Error GetId", create{}, "1", 422},
		{"Store Error", mockResponse{models.Seller{}, customerrors.ErrorInvalidID}, webResponse{"404", models.Seller{}, customerrors.ErrorInvalidID.Error()}, "Error GetId", sendCreate, "1", 404},
		{"Store Error", mockResponse{models.Seller{}, customerrors.ErrorConflict}, webResponse{"409", models.Seller{}, customerrors.ErrorConflict.Error()}, "Error GetId", sendCreate, "1", 409},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.NewSellers(mockService)

		mockService.On("Store", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(value.mockResponse.data, value.mockResponse.err).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)
		post, _ := json.Marshal(value.sendCreate)
		log.Println(sendCreate)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(post))
		router.POST("/", control.SellersStore())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse
		json.Unmarshal(body, &res)

		assert.Equal(t, value.expectResponse, res, value.message)
		assert.Equal(t, value.expectedstatuscode, w.Result().StatusCode, value.message)

	}

}
