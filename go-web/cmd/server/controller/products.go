package controller

import (
	"encoding/json"
	"errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/products"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/web"
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
	Product_code                     string  `json:"product_code" binding:"required"`
	Description                      string  `json:"description" binding:"required"`
	Width                            float64 `json:"width" binding:"required"`
	Height                           float64 `json:"height" binding:"required"`
	Length                           float64 `json:"length" binding:"required"`
	Net_weight                       float64 `json:"netweight" binding:"required"`
	Expiration_rate                  int     `json:"expiration_rate" binding:"required"`
	Recommended_freezing_temperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	Freezing_rate                    float64 `json:"freezing_rate" binding:"required"`
	Product_type_id                  int     `json:"product_type_id" binding:"required"`
	Seller_id                        int     `json:"seller_id" binding:"required"`
}

type updateProduct struct {
	Product_code                     *string  `json:"product_code,omitempty"`
	Description                      *string  `json:"description,omitempty"`
	Width                            *float64 `json:"width,omitempty"`
	Height                           *float64 `json:"height,omitempty"`
	Length                           *float64 `json:"length,omitempty"`
	Net_weight                       *float64 `json:"netweight,omitempty"`
	Expiration_rate                  *int     `json:"expiration_rate,omitempty"`
	Recommended_freezing_temperature *float64 `json:"recommended_freezing_temperature,omitempty"`
	Freezing_rate                    *float64 `json:"freezing_rate,omitempty"`
	Product_type_id                  *int     `json:"product_type_id,omitempty"`
	Seller_id                        *int     `json:"seller_id,omitempty"`
}
