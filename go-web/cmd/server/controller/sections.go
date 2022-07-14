package controller

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/sections/domain"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SectionsController interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetReportProducts(ctx *gin.Context)
}

type sectionsController struct {
	service domain.SectionService
}

func NewController(s domain.SectionService) SectionsController {
	newController := &sectionsController{
		service: s,
	}
	return newController
}

func (controller *sectionsController) GetAll(ctx *gin.Context) {
	sections, err := controller.service.GetAll(ctx.Request.Context())

	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}
	response := web.NewResponse(http.StatusOK, sections, "")
	ctx.JSON(http.StatusOK, response)
}

func (controller *sectionsController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	section, err := controller.service.GetById(ctx, id)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}
	response := web.NewResponse(http.StatusOK, section, "")
	ctx.JSON(http.StatusOK, response)
}

func (controller *sectionsController) Store(ctx *gin.Context) {
	var req storeSection
	if err := ctx.ShouldBindJSON(&req); err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	section, err := controller.service.Store(ctx, &domain.Section{
		SectionNumber:      req.SectionNumber,
		CurrentTemperature: req.CurrentTemperature,
		MinimumTemperature: req.MinimumTemperature,
		CurrentCapacity:    req.CurrentCapacity,
		MinimumCapacity:    req.MinimumCapacity,
		MaximumCapacity:    req.MaximumCapacity,
		WarehouseId:        req.WarehouseId,
		ProductTypeId:      req.ProductTypeId,
	})
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	response := web.NewResponse(http.StatusCreated, section, "")
	ctx.JSON(http.StatusCreated, response)
}

func (controller *sectionsController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	var req updateSection
	if err := ctx.ShouldBindJSON(&req); err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	section, err := controller.service.Update(ctx, &domain.Section{
		Id:                 id,
		SectionNumber:      req.SectionNumber,
		CurrentTemperature: req.CurrentTemperature,
		MinimumTemperature: req.MinimumTemperature,
		CurrentCapacity:    req.CurrentCapacity,
		MinimumCapacity:    req.MinimumCapacity,
		MaximumCapacity:    req.MaximumCapacity,
		WarehouseId:        req.WarehouseId,
		ProductTypeId:      req.ProductTypeId,
	})
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	response := web.NewResponse(http.StatusOK, section, "")
	ctx.JSON(http.StatusOK, response)
}

func (controller *sectionsController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	err = controller.service.Delete(ctx, id)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}

	response := web.NewResponse(http.StatusNoContent, nil, "")
	ctx.JSON(http.StatusNoContent, response)
}

func (controller *sectionsController) GetReportProducts(ctx *gin.Context) {
	queryId := ctx.Query("id")
	if queryId == "" {
		queryId = "0"
	}
	id, err := strconv.Atoi(queryId)

	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}
	reports, err := controller.service.GetReportProducts(ctx, id)

	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		ctx.JSON(status, res)
		return
	}
	response := web.NewResponse(http.StatusOK, reports, "")
	ctx.JSON(http.StatusOK, response)
}

type storeSection struct {
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseId        int `json:"warehouse_id" binding:"required"`
	ProductTypeId      int `json:"product_type_id" binding:"required"`
}

type updateSection struct {
	SectionNumber      int `json:"section_number,omitempty"`
	CurrentTemperature int `json:"current_temperature,omitempty"`
	MinimumTemperature int `json:"minimum_temperature,omitempty"`
	CurrentCapacity    int `json:"current_capacity,omitempty"`
	MinimumCapacity    int `json:"minimum_capacity,omitempty"`
	MaximumCapacity    int `json:"maximum_capacity,omitempty"`
	WarehouseId        int `json:"warehouse_id,omitempty"`
	ProductTypeId      int `json:"product_type_id,omitempty"`
}
