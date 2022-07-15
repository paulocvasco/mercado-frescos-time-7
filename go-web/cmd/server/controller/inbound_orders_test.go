package controller_test

import (
	"bytes"
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	controller2 "mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/inbound_orders/mocks"
	model "mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

var expectInboundOrders = model.InboundOrders{
	ID:             1,
	OrderDate:      "2022-09-09",
	OrderNumber:    "2536",
	EmployeeId:     1,
	ProductBatchId: 1,
	WareHouseId:    1,
}

var expectInboundConflict = model.InboundOrders{
	ID:             1,
	OrderDate:      "2022-09-09",
	OrderNumber:    "2536",
	EmployeeId:     1,
	ProductBatchId: 1,
}

func TestInboundOrdersController_InboundOrdersCreate(t *testing.T) {

	t.Run("should return code 201 and create a inbound order", func(t *testing.T) {
		service := mocks.NewService(t)
		controller := controller2.NewInboundOrders(service)

		service.On("Create", expectInboundOrders.OrderDate, expectInboundOrders.OrderNumber, expectInboundOrders.EmployeeId,
			expectInboundOrders.ProductBatchId, expectInboundOrders.WareHouseId).
			Return(expectInboundOrders, nil)

		r := gin.Default()

		body, _ := json2.Marshal(expectInboundOrders)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/inboundOrders", bytes.NewBuffer(body))
		r.POST("/inboundOrders", controller.Create())
		r.ServeHTTP(w, req)
		res, _ := ioutil.ReadAll(w.Body)
		json2.Unmarshal(res, &expectInboundOrders)

		assert.Equal(t, expectInboundOrders, expectInboundOrders)
		assert.Equal(t, 201, w.Code)

	})

	t.Run("shouldn't create a inbound order and return 409", func(t *testing.T) {
		service := mocks.NewService(t)
		controller := controller2.NewInboundOrders(service)

		service.On("Create", expectInboundOrders.OrderDate, expectInboundOrders.OrderNumber, expectInboundOrders.EmployeeId,
			expectInboundOrders.ProductBatchId, expectInboundOrders.WareHouseId).
			Return(model.InboundOrders{}, customErrors.ErrorConflict)

		r := gin.Default()
		r.POST("/inboundOrders", controller.Create())

		body, _ := json2.Marshal(expectInboundOrders)

		req, _ := http.NewRequest("POST", "/inboundOrders", bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 409, w.Code)

	})

	t.Run("shouldn't create a inbound order and return 422", func(t *testing.T) {
		service := mocks.NewService(t)
		controller := controller2.NewInboundOrders(service)

		service.On("Create", expectInboundOrders.OrderDate, expectInboundOrders.OrderNumber, expectInboundOrders.EmployeeId,
			expectInboundOrders.ProductBatchId, expectInboundOrders.WareHouseId).
			Return(model.InboundOrders{}, customErrors.ErrorConflict).Maybe()

		r := gin.Default()
		r.POST("/inboundOrders", controller.Create())

		body, _ := json2.Marshal(expectInboundConflict)

		req, _ := http.NewRequest("POST", "/inboundOrders", bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 422, w.Code)

	})
}
