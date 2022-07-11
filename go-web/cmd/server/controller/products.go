package controller

import (
	"encoding/json"
	"errors"
	"mercado-frescos-time-7/go-web/internal/products"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service products.Service
}

func NewProductHandler(p products.Service) ProductHandler {
	return ProductHandler{
		service: p,
	}
}

func (ph *ProductHandler) GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		pp, err := ph.service.GetAll()
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, pp, ""))
	}
}

func (ph *ProductHandler) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		obj, err := ph.service.GetById(id)
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
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, obj, ""))
	}
}

func (ph *ProductHandler) SaveProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		newProduct := saveProduct{}
		err := c.ShouldBindJSON(&newProduct)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		pJSON, err := json.Marshal(newProduct)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		p, err := ph.service.Insert(pJSON)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		c.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}

func (ph *ProductHandler) UpdateProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		updateProduct := updateProduct{}
		err := c.ShouldBindJSON(&updateProduct)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		p, err := json.Marshal(updateProduct)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		product, err := ph.service.Update(id, p)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

func (ph *ProductHandler) DeleteProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		err = ph.service.Delete(id)
		if err != nil {
			status, msg := customerrors.ErrorHandleResponse(err)
			res := web.NewResponse(status, nil, msg)
			c.JSON(status, res)
			return
		}
		c.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "asdasdasdasdasd"))
	}
}

type saveProduct struct {
	ProductCode                    string  `json:"product_code" binding:"required"`
	Description                    string  `json:"description" binding:"required"`
	Width                          float64 `json:"width" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	NetWeight                      float64 `json:"net_weight" binding:"required"`
	Expiration_rate                int     `json:"expiration_rate" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	FreezingRate                   float64 `json:"freezing_rate" binding:"required"`
	ProductTypeId                  int     `json:"product_type_id" binding:"required"`
	SellerId                       int     `json:"seller_id" binding:"required"`
}

type updateProduct struct {
	ProductCode                    *string  `json:"product_code,omitempty"`
	Description                    *string  `json:"description,omitempty"`
	Width                          *float64 `json:"width,omitempty"`
	Height                         *float64 `json:"height,omitempty"`
	Length                         *float64 `json:"length,omitempty"`
	NetWeight                      *float64 `json:"net_weight,omitempty"`
	ExpirationRate                 *int     `json:"expiration_rate,omitempty"`
	RecommendedFreezingTemperature *float64 `json:"recommended_freezing_temperature,omitempty"`
	FreezingRate                   *float64 `json:"freezing_rate,omitempty"`
	ProducTypeId                   *int     `json:"product_type_id,omitempty"`
	SellerId                       *int     `json:"seller_id,omitempty"`
}
