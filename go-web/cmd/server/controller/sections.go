package controller

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/sections"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SectionsController interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type sectionsController struct {
	service sections.Service
}

func NewController(s sections.Service) SectionsController {
	newController := &sectionsController{
		service: s,
	}
	return newController
}

func (controller *sectionsController) GetAll(ctx *gin.Context) {
	response, err := controller.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (controller *sectionsController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := controller.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, customerrors.ErrorSectionNotFound)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (controller *sectionsController) Store(ctx *gin.Context) {
	newSection := storeSection{}
	err := ctx.ShouldBindJSON(&newSection)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, customerrors.ErrorSectionNumber)
		return
	}
	sectionToJson, err := json.Marshal(newSection)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, customerrors.ErrorSectionNumber)
		return
	}
	section, err := controller.service.Store(sectionToJson)
	if err != nil {
		ctx.JSON(http.StatusConflict, customerrors.ErrorCardIdAlreadyExists)
		return
	}
	ctx.JSON(http.StatusCreated, section)

}

func (controller *sectionsController) Update(ctx *gin.Context) {
	newSection := updateSection{}
	err := ctx.ShouldBindJSON(&newSection)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, customerrors.ErrorSectionNumber)
		return
	}
	section, err := json.Marshal(newSection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	id := ctx.Param("id")
	data, err := controller.service.Update(id, section)
	if err != nil {
		ctx.JSON(http.StatusNotFound, customerrors.ErrorSectionNotFound)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (controller *sectionsController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := controller.service.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, customerrors.ErrorSectionNotFound)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
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
