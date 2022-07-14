package controller

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/carriers"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"

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
	var id int
	qp := c.Query("id")
	if qp != "" {
		var err error
		id, err = strconv.Atoi(qp)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
	}
	rep, err := ctrl.service.Get(id)
	if err != nil {
		status, msg := customerrors.ErrorHandleResponse(err)
		res := web.NewResponse(status, nil, msg)
		c.JSON(status, res)
		return
	}
	res := web.NewResponse(http.StatusOK, rep, "")
	c.JSON(http.StatusOK, res)
}
