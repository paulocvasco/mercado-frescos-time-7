package controller

import (
	"errors"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/internal/warehouse"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	response, err := control.service.GetByID(id)
	if err != nil {
		if errors.Is(err, customerrors.ErrorInvalidID) {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		} else {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
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
	c.JSON(http.StatusOK, response)
}

func (control *warehousesController) UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	body := c.Request.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = control.service.Update(index, data)
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
