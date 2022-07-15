package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mercado-frescos-time-7/go-web/internal/inbound_orders"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
)

type InboundOrdersController struct {
	service inbound_orders.Service
}

type requestInboundOrders struct {
	OrderDate      string `json:"order_date" binding:"required"`
	OrderNumber    string `json:"order_number" binding:"required"`
	EmployeeId     int    `json:"employee_id" binding:"required"`
	ProductBatchId int    `json:"product_batch_id" binding:"required"`
	WarehouseId    int    `json:"warehouse_id" binding:"required"`
}

func (c *InboundOrdersController) Create() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req requestInboundOrders
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		inbound, err := c.service.Create(req.OrderDate, req.OrderNumber, req.EmployeeId, req.ProductBatchId, req.WarehouseId)

		if err != nil {
			log.Println(err)
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, inbound, ""))
	}

}

func NewInboundOrders(i inbound_orders.Service) *InboundOrdersController {
	return &InboundOrdersController{
		service: i,
	}
}
