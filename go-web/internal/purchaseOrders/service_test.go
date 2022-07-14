package purchaseOrders_test

import (
	"errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/purchaseOrders"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/purchaseOrders/mocks"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {

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
		name        string
		mockReponse repository.ResultPost
		sendCreate  models.PurchaseOrders
		expectError error
	}

	testsCases := []tests{
		{"Created Ok", response, sendCreate, nil},
		{"Error Created", repository.ResultPost{}, models.PurchaseOrders{}, errors.New("Error")},
	}

	for _, value := range testsCases {
		mockRepository := mocks.NewRepositoryMysql(t)
		service := purchaseOrders.NewService(mockRepository)

		mockRepository.On("Create", mock.Anything).Return(value.mockReponse, value.expectError)
		resp, err := service.Create(sendCreate)
		assert.Equal(t, value.mockReponse, resp)
		assert.Equal(t, value.expectError, err)
	}

}

func TestGetPurchaseOrder(t *testing.T) {

	responseAll := []models.ResponsePurchaseByBuyer{
		{ID: 1, CardNumberID: "Card#1", FirstName: "Pedro", LastName: "Avelar", PurchaseOrdersCount: 2},
		{ID: 2, CardNumberID: "Card#2", FirstName: "Mario", LastName: "Campos", PurchaseOrdersCount: 3},
	}

	responseOnly := []models.ResponsePurchaseByBuyer{
		{ID: 1, CardNumberID: "Card#1", FirstName: "Pedro", LastName: "Avelar", PurchaseOrdersCount: 2},
	}

	type tests struct {
		name        string
		mockReponse []models.ResponsePurchaseByBuyer
		sendId      int
		expectError error
	}

	testsCases := []tests{
		{"Get Id=0", responseAll, 0, nil},
		{"Get Id=1", responseOnly, 1, nil},
		{"Get Id=0 Error", []models.ResponsePurchaseByBuyer{}, 0, customerrors.ErrorInvalidDB},
		{"Get Id=1 Error", []models.ResponsePurchaseByBuyer{}, 1, customerrors.ErrorInvalidID},
	}

	for _, value := range testsCases {
		mockRepository := mocks.NewRepositoryMysql(t)
		service := purchaseOrders.NewService(mockRepository)
		mockRepository.On("GetAllPurchaseOrder").Return(value.mockReponse, value.expectError).Maybe()
		mockRepository.On("GetIdPurchaseOrder", value.sendId).Return(value.mockReponse, value.expectError).Maybe()
		resp, err := service.GetPurchaseOrder(value.sendId)
		assert.Equal(t, value.mockReponse, resp)
		assert.Equal(t, value.expectError, err)
	}

}
