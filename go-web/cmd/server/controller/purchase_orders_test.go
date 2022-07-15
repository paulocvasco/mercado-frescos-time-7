package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/purchaseOrders/mocks"
	"mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePurschaseOrders(t *testing.T) {
	response := repository.ResultPost{
		Id:            1,
		OrderNumber:   "order#2",
		OrderDate:     "2021-04-04",
		TrackingCode:  "abscf124",
		BuyerId:       1,
		CarrierID:     1,
		OrderStatusId: 1,
		WareHouseID:   1,
	}

	sendCreate := models.PurchaseOrders{
		OrderNumber:   "order#2",
		OrderDate:     "2021-04-04",
		TrackingCode:  "abscf124",
		BuyerId:       1,
		CarrierID:     1,
		OrderStatusId: 1,
		WareHouseID:   1,
	}

	type tests struct {
		name               string
		mockReponse        repository.ResultPost
		sendCreate         models.PurchaseOrders
		expectError        error
		expectedstatuscode int
	}
	type webResponse struct {
		Code  string                `json:"code"`
		Data  repository.ResultPost `json:"data"`
		Error string                `json:"error"`
	}

	testsCases := []tests{
		{"Created Ok", response, sendCreate, nil, 201},
		{"Error Created 40", repository.ResultPost{}, models.PurchaseOrders{}, errors.New("Erro"), 422},
		{"Error Created 40", repository.ResultPost{}, sendCreate, customerrors.ErrorInvalidDB, 500},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.PurchaseOrdersNewController(mockService)

		mockService.On("Create", value.sendCreate).Return(value.mockReponse, value.expectError).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)
		post, _ := json.Marshal(value.sendCreate)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(post))
		router.POST("/", control.CreatePurchaseOrders())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse
		json.Unmarshal(body, &res)

		assert.Equal(t, value.mockReponse, res.Data)
		assert.Equal(t, value.expectedstatuscode, w.Result().StatusCode)

	}

}

func TestGetPurchaseOrder(t *testing.T) {
	type webResponse struct {
		Code  string                           `json:"code"`
		Data  []models.ResponsePurchaseByBuyer `json:"data"`
		Error string                           `json:"error"`
	}

	responseAll := []models.ResponsePurchaseByBuyer{
		{ID: 1, CardNumberID: "Card#1", FirstName: "Pedro", LastName: "Avelar", PurchaseOrdersCount: 2},
		{ID: 2, CardNumberID: "Card#2", FirstName: "Mario", LastName: "Campos", PurchaseOrdersCount: 3},
	}

	responseOnly := []models.ResponsePurchaseByBuyer{
		{ID: 1, CardNumberID: "Card#1", FirstName: "Pedro", LastName: "Avelar", PurchaseOrdersCount: 2},
	}
	type tests struct {
		name               string
		sendId             string
		mockResponse       []models.ResponsePurchaseByBuyer
		ExpectedError      error
		expectedstatuscode int
	}

	testsCases := []tests{
		{"Get Id=0", "", responseAll, nil, 200},
		{"Get Id=1", "1", responseOnly, nil, 200},
		{"Get Id=Casa Error", "casa", []models.ResponsePurchaseByBuyer(nil), customerrors.ErrorInvalidIDParameter, 400},
		{"Get Id=1 Error", "1", []models.ResponsePurchaseByBuyer(nil), customerrors.ErrorInvalidID, 404},
	}

	for _, value := range testsCases {
		gin.SetMode(gin.TestMode)
		mockService := mocks.NewService(t)
		control := controller.PurchaseOrdersNewController(mockService)

		mockService.On("GetPurchaseOrder", mock.Anything).Return(value.mockResponse, value.ExpectedError).Maybe()

		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?id=%v", value.sendId), nil)
		router.GET("/", control.GetPurchaseOrder())
		router.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)

		var res webResponse
		json.Unmarshal(body, &res)

		log.Println(res.Data)
		log.Println(value.mockResponse)

		assert.Equal(t, value.mockResponse, res.Data, value.name)
		assert.Equal(t, value.expectedstatuscode, w.Result().StatusCode, value.name)
	}

}
