package controller

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/locality"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestLocality struct {
	Id            string `json:"id" binding:"required"`
	Locality_name string `json:"locality_name" binding:"required"`
	Province_name string `json:"province_name" binding:"required"`
	Country_name  string `json:"country_name" binding:"required"`
}

type Localities struct {
	service locality.Service
}

func (c *Localities) LocalityStore() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestLocality
		if err := ctx.ShouldBindJSON(&req); err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		p, err := c.service.Store(models.Locality{Id: req.Id, Locality_name: req.Locality_name, Province_name: req.Province_name, Country_name: req.Country_name})
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}

func NewLocality(p locality.Service) *Localities {
	return &Localities{
		service: p,
	}
}
