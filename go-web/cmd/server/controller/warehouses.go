package controller

import (
	"errors"
	"io/ioutil"
	"mercado-frescos-time-7/go-web/internal/models"
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
	response, err := control.service.GetAll()
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
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
	var newWarehouse models.Warehouse
	err := c.ShouldBindJSON(&newWarehouse)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}

	response, err := control.service.Create(newWarehouse)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (control *warehousesController) UpdateWarehouse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}

	body := c.Request.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := control.service.Update(id, data)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (control *warehousesController) DeleteWarehouse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}

	err = control.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
