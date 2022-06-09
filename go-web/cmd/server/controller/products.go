package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/internal/products"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno, tente mais tarde",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": pp})
	}
}

func (ph *ProductHandler) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id inválido",
			})
			return
		}
		obj, err := ph.service.GetById(id)
		if err != nil {
			if errors.Is(err, customerrors.ErrorInvalidID) {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "id nao encontrado",
				})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "erro interno, tente mais tarde",
				})
				return
			}
		}
		c.JSON(http.StatusOK, obj)
	}
}

func (ph *ProductHandler) SaveProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		newProduct := saveProduct{}
		err := c.ShouldBindJSON(&newProduct)
		var ve validator.ValidationErrors
		var js *json.SyntaxError
		var jt *json.UnmarshalTypeError
		if err != nil {
			if errors.As(err, &ve) {
				for _, v := range ve {
					c.JSON(http.StatusUnprocessableEntity, gin.H{
						"message": fmt.Sprintf("erro no campo: %v", v.Field()),
					})
					return
				}
			} else if errors.As(err, &js) {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "confira a estrutura do JSON",
				})
				return
			} else if errors.As(err, &jt) {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("confira o tipo do campo %v", jt.Field),
				})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "erro interno, tente mais tarde",
				})
				return
			}
		}
		pJSON, err := json.Marshal(newProduct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno, tente mais tarde",
			})
			return
		}
		p, err := ph.service.Insert(pJSON)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno, tente mais tarde",
			})
			return
		}
		c.JSON(http.StatusCreated, p)
	}
}

func (ph *ProductHandler) UpdateProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		updateProduct := updateProduct{}
		err := c.ShouldBindJSON(&updateProduct)
		var ve validator.ValidationErrors
		var js *json.SyntaxError
		var jt *json.UnmarshalTypeError
		if err != nil {
			if errors.As(err, &ve) {
				for _, v := range ve {
					c.JSON(http.StatusUnprocessableEntity, gin.H{
						"message": fmt.Sprintf("erro no campo: %v", v.Field()),
					})
					return
				}
			} else if errors.As(err, &js) {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "confira a estrutura do JSON",
				})
				return
			} else if errors.As(err, &jt) {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("confira o tipo do campo %v", jt.Field),
				})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "erro interno",
				})
				return
			}
		}
		p, err := json.Marshal(updateProduct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno tente novamente",
			})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id inválido",
			})
			return
		}
		product, err := ph.service.Update(id, p)
		if err != nil {
			if errors.Is(err, customerrors.ErrorInvalidID) {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "id nao encontrado",
				})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "erro interno, tente mais tarde",
				})
				return
			}
		}
		c.JSON(http.StatusOK, product)
	}
}

func (ph *ProductHandler) DeleteProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id inválido",
			})
			return
		}
		err = ph.service.Delete(id)
		if err != nil {
			if errors.Is(err, customerrors.ErrorInvalidID) {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "id nao encontrado",
				})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "erro interno, tente mais tarde",
				})
				return
			}
		}
		c.JSON(http.StatusNoContent, gin.H{
			"status": 204,
		})
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
