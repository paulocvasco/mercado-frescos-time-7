package controller

import (
	"github.com/gin-gonic/gin"
	"mercado-frescos-time-7/go-web/internal/productBatch/domain"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
)

type ProductBatchController interface {
	Store(*gin.Context)
}

type productBatchController struct {
	service domain.ProductBatchService
}

func NewControllerProductBatch(s domain.ProductBatchService) ProductBatchController {
	newController := &productBatchController{
		service: s,
	}

	return newController
}

func (c productBatchController) Store(ctx *gin.Context) {
	var req storeProductBatch

	if err := ctx.ShouldBindJSON(&req); err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	productBatch, err := c.service.Store(ctx, &domain.ProductBatch{
		BatchNumber:        *req.BatchNumber,
		CurrentQuantity:    *req.CurrentQuantity,
		CurrentTemperature: *req.CurrentTemperature,
		DueDate:            *req.DueDate,
		InitialQuantity:    *req.InitialQuantity,
		ManufacturingDate:  *req.ManufacturingDate,
		ManufacturingHour:  *req.ManufacturingHour,
		MinimumTemperature: *req.MinimumTemperature,
		ProductId:          *req.ProductId,
		SectionId:          *req.SectionId,
	})

	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	response := web.NewResponse(http.StatusCreated, productBatch, "")
	ctx.JSON(http.StatusCreated, response)
}

type storeProductBatch struct {
	BatchNumber        *int    `json:"batch_number" binding:"required"`
	CurrentQuantity    *int    `json:"current_quantity" binding:"required"`
	CurrentTemperature *int    `json:"current_temperature" binding:"required"`
	DueDate            *string `json:"due_date" binding:"required"`
	InitialQuantity    *int    `json:"initial_quantity" binding:"required"`
	ManufacturingDate  *string `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  *int    `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature *int    `json:"minimum_temperature" binding:"required"`
	ProductId          *int    `json:"product_id" binding:"required"`
	SectionId          *int    `json:"section_id" binding:"required"`
}
