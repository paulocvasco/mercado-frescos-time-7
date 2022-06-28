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
	warehouses, err := control.service.GetAll()
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	response := web.NewResponse(http.StatusOK, warehouses, "")
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
	warehouse, err := control.service.GetByID(id)
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
	response := web.NewResponse(http.StatusOK, warehouse, "")
	c.JSON(http.StatusOK, response)
}

func (control *warehousesController) CreateWarehouse(c *gin.Context) {
	var newWarehouse models.PostWarehouse
	err := c.ShouldBindJSON(&newWarehouse)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}

	nw, err := control.service.Create(newWarehouse)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	response := web.NewResponse(http.StatusCreated, nw, "")
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

	updatedWarehouse, err := control.service.Update(id, data)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	response := web.NewResponse(http.StatusOK, updatedWarehouse, "")
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
	response := web.NewResponse(http.StatusNoContent, nil, "")
	c.JSON(http.StatusNoContent, response)
}
