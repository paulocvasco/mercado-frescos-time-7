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
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (controller *controller) Store(ctx *gin.Context) {
	newSection := storeSection{}
	err := ctx.ShouldBindJSON(&newSection)
	// tratar erros + especificamente
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
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, section)
}

func (controller *controller) Update(ctx *gin.Context) {
	newSection := models.Section{}
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
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (controller *controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := controller.service.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
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
