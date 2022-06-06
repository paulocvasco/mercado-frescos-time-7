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
		Section_number:      newSection.Section_number,
		Current_temperature: newSection.Current_temperature,
		Minimum_temperature: newSection.Minimum_temperature,
		Current_capacity:    newSection.Current_capacity,
		Minimum_capacity:    newSection.Minimum_capacity,
		Maximum_capacity:    newSection.Maximum_capacity,
		Warehouse_id:        newSection.Warehouse_id,
		Product_type_id:     newSection.Product_type_id,
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
	Section_number      int `json:"section_number" binding:"required"`
	Current_temperature int `json:"current_temperature" binding:"required"`
	Minimum_temperature int `json:"minimum_temperature" binding:"required"`
	Current_capacity    int `json:"current_capacity" binding:"required"`
	Minimum_capacity    int `json:"minimum_capacity" binding:"required"`
	Maximum_capacity    int `json:"Maximum_capacity" binding:"required"`
	Warehouse_id        int `json:"warehouse_id" binding:"required"`
	Product_type_id     int `json:"product_type_id" binding:"required"`
}

type updateSection struct {
	Section_number      int `json:"section_number,omitempty"`
	Current_temperature int `json:"current_temperature,omitempty"`
	Minimum_temperature int `json:"minimum_temperature,omitempty"`
	Current_capacity    int `json:"current_capacity,omitempty"`
	Minimum_capacity    int `json:"minimum_capacity,omitempty"`
	Maximum_capacity    int `json:"Maximum_capacity,omitempty"`
	Warehouse_id        int `json:"warehouse_id,omitempty"`
	Product_type_id     int `json:"product_type_id,omitempty"`
}
