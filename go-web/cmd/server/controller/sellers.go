package controller

import (
	"errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/seller"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/web"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Sellers struct {
	service seller.Service
}

type request struct {
	Cid         int    `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  string `json:"locality_id" binding:"required"`
}

type getAllResponse struct {
	Seller []models.Seller `json:"sellers"`
}

var gar getAllResponse

func (c *Sellers) SellersStore() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		p, err := c.service.Store(models.Seller{Cid: req.Cid, Company_name: req.CompanyName, Address: req.Address, Telephone: req.Telephone, LocalityID: req.LocalityID})
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}

func (c *Sellers) SellersGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		gar.Seller = p
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, gar, ""))
	}
}

func (c *Sellers) SellersGetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		p, err := c.service.GetId(int(id))
		if err != nil {
			if errors.Is(err, customerrors.ErrorInvalidID) {
				status, msg := customerrors.ErrorHandleResponse(err)
				res := web.NewResponse(status, nil, msg)
				ctx.JSON(status, res)
				return
			} else {
				status, msg := customerrors.ErrorHandleResponse(err)
				res := web.NewResponse(status, nil, msg)
				ctx.JSON(status, res)
				return
			}
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

func (c *Sellers) SellersUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		body := ctx.Request.Body
		defer body.Close()

		data, err := ioutil.ReadAll(body)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		p, err := c.service.Update(data, id)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

func (c *Sellers) SellersDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			ctx.JSON(status, res)
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))

	}
}

func NewSellers(p seller.Service) *Sellers {
	return &Sellers{
		service: p,
	}
}
