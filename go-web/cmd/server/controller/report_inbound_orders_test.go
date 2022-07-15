package controller_test

import (
	json2 "encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/employees/mocks"
	model "mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

var inboundList = []model.ReportInboundOrders{
	{ID: 1,
		CardNumberId:       "40543",
		FirstName:          "Alice",
		LastName:           "Souza",
		WareHouseId:        1,
		InboundOrdersCount: 1,
	},

	{ID: 2,
		CardNumberId:       "405443",
		FirstName:          "Alice",
		LastName:           "Souza",
		WareHouseId:        2,
		InboundOrdersCount: 1,
	},
}

func TestReportController_GetReportInboundOrders(t *testing.T) {

	t.Run("should return all report and code 200", func(t *testing.T) {
		service := mocks.NewServiceReport(t)
		controller := controller.NewReport(service)
		service.On("GetReportInboundOrders", mock.Anything).Return(inboundList, nil)

		r := gin.Default()
		r.GET("/employees/reportInboundOrders", controller.GetReportInboundOrders())

		req, _ := http.NewRequest("GET", "/employees/reportInboundOrders", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		res, _ := ioutil.ReadAll(w.Body)
		json2.Unmarshal(res, &inboundList)

		assert.Equal(t, inboundList, inboundList)
		assert.Equal(t, 200, w.Code)

	})

	t.Run("should return one report and code 200", func(t *testing.T) {
		service := mocks.NewServiceReport(t)
		controller := controller.NewReport(service)
		service.On("GetReportInboundOrders", mock.Anything).Return(inboundList, nil)

		r := gin.Default()
		r.GET("/employees/reportInboundOrders", controller.GetReportInboundOrders())

		req, _ := http.NewRequest("GET", "/employees/reportInboundOrders?id=1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		res, _ := ioutil.ReadAll(w.Body)
		json2.Unmarshal(res, &inboundList)

		assert.Equal(t, inboundList, inboundList)
		assert.Equal(t, 200, w.Code)

	})

	t.Run("should return code 404", func(t *testing.T) {
		service := mocks.NewServiceReport(t)
		controller := controller.NewReport(service)
		service.On("GetReportInboundOrders", mock.Anything).Return([]model.ReportInboundOrders{}, customErrors.ErrorItemNotFound)

		r := gin.Default()
		r.GET("/employees/reportInboundOrders", controller.GetReportInboundOrders())

		req, _ := http.NewRequest("GET", "/employees/reportInboundOrders?id=1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		res, _ := ioutil.ReadAll(w.Body)
		json2.Unmarshal(res, &inboundList)

		assert.Equal(t, 404, w.Code)

	})

	t.Run("should return code 400", func(t *testing.T) {
		service := mocks.NewServiceReport(t)
		controller := controller.NewReport(service)
		service.On("GetReportInboundOrders", mock.Anything).Return([]model.ReportInboundOrders{}, errors.New("invalid param")).Maybe()

		r := gin.Default()
		r.GET("/employees/reportInboundOrders", controller.GetReportInboundOrders())

		req, _ := http.NewRequest("GET", "/employees/reportInboundOrders?id=number", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		res, _ := ioutil.ReadAll(w.Body)
		json2.Unmarshal(res, &inboundList)

		assert.Equal(t, 400, w.Code)

	})
}
