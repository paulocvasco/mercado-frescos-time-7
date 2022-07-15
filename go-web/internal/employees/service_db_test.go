package employees_test

import (
	"github.com/go-playground/assert/v2"
	"mercado-frescos-time-7/go-web/internal/employees"
	"mercado-frescos-time-7/go-web/internal/employees/mocks"
	model "mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
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

func TestServiceReport_GetReportInboundOrders(t *testing.T) {
	repository := mocks.NewReportInterface(t)
	t.Run("should return report inbound order", func(t *testing.T) {

		repository.
			On("GetReportInboundOrders", 2).Return(inboundList, nil).Maybe()

		service := employees.NewServiceReport(repository)

		inboundList, _ := service.GetReportInboundOrders(2)

		assert.Equal(t, inboundList, inboundList)
	})

	t.Run("should return an error", func(t *testing.T) {

		repository.On("GetReportInboundOrders", 8).
			Return([]model.ReportInboundOrders{}, customErrors.ErrorInvalidID).Once()

		service := employees.NewServiceReport(repository)

		_, err := service.GetReportInboundOrders(8)
		assert.Equal(t, err, err)

	})
}
