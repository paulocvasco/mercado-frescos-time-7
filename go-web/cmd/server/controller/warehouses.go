package controller

import (
	"io/ioutil"
	"mercado-frescos-time-7/go-web/internal/warehouse"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WarehousesController interface {
	GetAllWarehouse(*gin.Context)
	GetByIDWarehouse(*gin.Context)
	CreateWarehouse(*gin.Context)
	UpdateWarehouse(*gin.Context)
	DeleteWarehouse(*gin.Context)
}

type warehousesController struct {
	service warehouse.Service
}

func NewControllerWarehouse(s warehouse.Service) WarehousesController {
	newController := &warehousesController{
		service: s,
	}
	return newController
}

func (control *warehousesController) GetAllWarehouse(c *gin.Context) {
	response := control.service.GetAll()
	c.JSON(http.StatusOK, response)
}

func (control *warehousesController) GetByIDWarehouse(c *gin.Context) {
	param := c.Param("id")
	response, err := control.service.GetByID(param)
	if err != nil {
		switch err {
		case customerrors.ErrorInvalidIDParameter:
			c.JSON(http.StatusBadRequest, err)
		case customerrors.ErrorInvalidID:
			c.JSON(http.StatusNotFound, err)
		default:
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, response)
}

func (control *warehousesController) CreateWarehouse(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := control.service.Create(data)
	if err != nil {
		switch err {
		case customerrors.ErrorMissingAddres:
			c.JSON(http.StatusUnprocessableEntity, err)
		case customerrors.ErrorMissingTelephone:
			c.JSON(http.StatusUnprocessableEntity, err)
		case customerrors.ErrorMissingCapacity:
			c.JSON(http.StatusUnprocessableEntity, err)
		case customerrors.ErrorMissingTemperature:
			c.JSON(http.StatusUnprocessableEntity, err)
		default:
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, resposne)
}

func (control *warehousesController) UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	body := c.Request.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = control.service.Update(id, data)
	if err != nil {
		switch err {
		case customerrors.ErrorInvalidIDParameter:
			c.JSON(http.StatusNotFound, err)
		case customerrors.ErrorInvalidID:
			c.JSON(http.StatusNotFound, err)
		default:
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (control *warehousesController) DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")

	err := control.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusNoContent, err)
}
