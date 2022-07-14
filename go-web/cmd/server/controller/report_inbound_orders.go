package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/employees"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"
)

type ReportController struct {
	service employees.ServiceReport
}

func (c *ReportController) GetReportInboundOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idEmployee := ctx.Query("id")
		if idEmployee == "" {
			idEmployee = "0"
		}

		id, err := strconv.Atoi(idEmployee)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		report, err := c.service.GetReportInboundOrders(id)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, report, ""))
	}
}

func NewReport(r employees.ServiceReport) *ReportController {
	return &ReportController{
		service: r,
	}

}
