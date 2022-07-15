package controller

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/purchaseOrders"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PurchaseOrdersController struct {
	service purchaseOrders.Service
}

func PurchaseOrdersNewController(p purchaseOrders.Service) *PurchaseOrdersController {
	return &PurchaseOrdersController{
		service: p,
	}
}

func (p *PurchaseOrdersController) CreatePurchaseOrders() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input models.PurchaseOrders
		if err := context.ShouldBindJSON(&input); err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}

		inboundOrders, err := p.service.Create(input)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		context.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, inboundOrders, ""))
	}

}
func (p *PurchaseOrdersController) GetPurchaseOrder() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Query("id")
		var intId int
		var err error
		if id == "" {
			intId = 0
		} else {
			intId, err = strconv.Atoi(id)

		}
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		all, err := p.service.GetPurchaseOrder(intId)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			context.JSON(status, res)
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, all, ""))
	}
}
