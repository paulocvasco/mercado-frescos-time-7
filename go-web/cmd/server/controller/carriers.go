package controller

import (
	"mercado-frescos-time-7/go-web/internal/carriers"
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
	c.JSON(http.StatusOK, nil)
}

func (ctrl *carriersController) GetCarriers(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
