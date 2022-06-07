package controller

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/sections"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type controller struct {
	service sections.Service
}

func NewController(s sections.Service) Controller {
	newController := &controller{
		service: s,
	}
	return newController
}

func (controller *controller) GetAll(ctx *gin.Context) {
	response, err := controller.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, web.NewResponse(200, response, ""))
}

func (controller *controller) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := controller.service.GetById(id)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, response, "Section não encontrada"))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (controller *controller) Store(ctx *gin.Context) {
	newSection := storeSection{}
	err := ctx.ShouldBindJSON(&newSection)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	section := models.Section{
		SectionNumber:      newSection.SectionNumber,
		CurrentTemperature: newSection.CurrentTemperature,
		MinimumTemperature: newSection.MinimumTemperature,
		CurrentCapacity:    newSection.CurrentCapacity,
		MinimumCapacity:    newSection.MinimumCapacity,
		MaximumCapacity:    newSection.MaximumCapacity,
		WarehouseId:        newSection.WarehouseId,
		ProductTypeId:      newSection.ProductTypeId,
	}

	err = controller.service.Store(section)

	if err != nil {
		ctx.JSON(422, web.NewResponse(422, nil, "Não contém os campos necessários"))
		return
	}

	ctx.JSON(201, web.NewResponse(201, section, "Section criada com sucesso!"))
}

func (controller *controller) Update(ctx *gin.Context) {
	newSection := updateSection{}
	err := ctx.ShouldBindJSON(&newSection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	section, err := json.Marshal(newSection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	id := ctx.Param("id")
	err = controller.service.Update(id, section)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, "Section não encontrada!"))
		return
	}
	ctx.JSON(200, web.NewResponse(20, section, "Section atualizada com sucesso!"))
}

func (controller *controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := controller.service.Delete(id)

	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, "Section não encontrada!"))
		return
	}
	ctx.JSON(204, web.NewResponse(204, nil, "Section deletada com sucesso!"))
}

type storeSection struct {
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"MaximumCapacity" binding:"required"`
	WarehouseId        int `json:"warehouse_id" binding:"required"`
	ProductTypeId      int `json:"product_type_id" binding:"required"`
}

type updateSection struct {
	SectionNumber      int `json:"section_number,omitempty"`
	CurrentTemperature int `json:"current_temperature,omitempty"`
	MinimumTemperature int `json:"minimum_temperature,omitempty"`
	CurrentCapacity    int `json:"current_capacity,omitempty"`
	MinimumCapacity    int `json:"minimum_capacity,omitempty"`
	MaximumCapacity    int `json:"MaximumCapacity,omitempty"`
	WarehouseId        int `json:"warehouse_id,omitempty"`
	ProductTypeId      int `json:"product_type_id,omitempty"`
}
