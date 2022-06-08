package controller

import (
	"io/ioutil"
	customerrors "mercado-frescos-time-7/go-web/internal/custom_errors"
	"mercado-frescos-time-7/go-web/internal/warehouse"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetAllWarehouse(*gin.Context)
	GetByIDWarehouse(*gin.Context)
	CreateWarehouse(*gin.Context)
	UpdateWarehouse(*gin.Context)
	DeleteWarehouse(*gin.Context)
}

type controller struct {
	service warehouse.Service
}

func NewControllerWarehouse(s warehouse.Service) Controller {
	newController := &controller{
		service: s,
	}
	return newController
}

func (control *controller) GetAllWarehouse(c *gin.Context) {
	response, err := control.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (control *controller) GetByIDWarehouse(c *gin.Context) {
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

func (control *controller) CreateWarehouse(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = control.service.Create(data)
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
	c.JSON(http.StatusOK, nil)
}

func (control *controller) UpdateWarehouse(c *gin.Context) {
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

func (control *controller) DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")

	err := control.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusNoContent, err)
}
