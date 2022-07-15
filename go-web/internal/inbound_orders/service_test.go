package inbound_orders_test

import (
	"github.com/stretchr/testify/assert"
	"mercado-frescos-time-7/go-web/internal/inbound_orders"
	"mercado-frescos-time-7/go-web/internal/inbound_orders/repository"
	mocks2 "mercado-frescos-time-7/go-web/internal/inbound_orders/repository/mocks"
	"mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"
)

var expectInboundOrders = repository.InboundOrders{
	ID:             1,
	OrderDate:      "2022-09-09",
	OrderNumber:    "2536",
	EmployeeId:     1,
	ProductBatchId: 1,
	WareHouseId:    1,
}

var expectInboundOrdersModel = models.InboundOrders{
	ID:             1,
	OrderDate:      "2022-09-09",
	OrderNumber:    "2536",
	EmployeeId:     1,
	ProductBatchId: 1,
	WareHouseId:    1,
}

func TestService_Create(t *testing.T) {

	t.Run("should create a inbound orders", func(t *testing.T) {
		repository := mocks2.NewRepository(t)
		// repository.On("GetCardNumberId", expectBuyer.CardNumberID).Return(nil).Once()
		repository.On("Create", expectInboundOrders.OrderDate, expectInboundOrders.OrderNumber, expectInboundOrders.EmployeeId,
			expectInboundOrders.ProductBatchId, expectInboundOrders.WareHouseId).
			Return(expectInboundOrders, nil).Once()

		service := inbound_orders.NewService(repository)

		result, _ := service.Create(expectInboundOrders.OrderDate, expectInboundOrders.OrderNumber, expectInboundOrders.EmployeeId,
			expectInboundOrders.ProductBatchId, expectInboundOrders.WareHouseId)

		assert.Equal(t, expectInboundOrdersModel, result)
	})

	t.Run("shouldn`t create a inbound order", func(t *testing.T) {
		repository := mocks2.NewRepository(t)
		// repository.On("GetCardNumberId", expectBuyer.CardNumberID).Return(nil)
		repository.On("Create", expectInboundOrders.OrderDate, expectInboundOrders.OrderNumber, expectInboundOrders.EmployeeId,
			expectInboundOrders.ProductBatchId, expectInboundOrders.WareHouseId).
			Return(expectInboundOrders, customErrors.ErrorInvalidOrderNumber).Once()

		service := inbound_orders.NewService(repository)

		_, err := service.Create(expectInboundOrders.OrderDate, expectInboundOrders.OrderNumber, expectInboundOrders.EmployeeId,
			expectInboundOrders.ProductBatchId, expectInboundOrders.WareHouseId)

		assert.Equal(t, err, customErrors.ErrorInvalidOrderNumber)

	})

}
