package controller

import (
	"mercado-frescos-time-7/go-web/internal/carriers"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CarriersController interface {
	CreateCarrier(*gin.Context)
	GetCarriers(*gin.Context)
}

type carriersController struct {
	service carriers.Service
}

func NewControllerCarriers(s carriers.Service) CarriersController {
	newController := &carriersController{
		service: s,
	}
	return newController
}

func (ctrl *carriersController) CreateCarrier(c *gin.Context) {
	var newCarrier models.CarrierRequest
	err := c.ShouldBindJSON(&newCarrier)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	nc, err := ctrl.service.Create(newCarrier)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	res := web.NewResponse(http.StatusCreated, nc, "")
	c.JSON(http.StatusCreated, res)
}

func (ctrl *carriersController) GetCarriers(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
