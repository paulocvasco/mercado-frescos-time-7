package controller

import (
	"mercado-frescos-time-7/go-web/internal/sections/domain"
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
	GetAllReportProducts(ctx *gin.Context)
	GetReportProductsById(ctx *gin.Context)
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, sections)
}

func (controller *sectionsController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	section, err := controller.service.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, section)
}

func (controller *sectionsController) Store(ctx *gin.Context) {
	var req storeSection
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, section)
}

func (controller *sectionsController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req updateSection
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, section)
}

func (controller *sectionsController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = controller.service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *sectionsController) GetAllReportProducts(ctx *gin.Context) {

	section, err := controller.service.GetAllProductReports(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, section)
	return

}

func (controller *sectionsController) GetReportProductsById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if id == 0 {
		section, err := controller.service.GetAllProductReports(ctx)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, section)
		return
	} else {
		section, err := controller.service.GetById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, section)
		return
	}
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
