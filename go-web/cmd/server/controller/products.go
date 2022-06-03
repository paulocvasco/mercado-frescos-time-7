package controller

import (
	"errors"
	"fmt"
	"mercado-frescos-time-7/go-web/internal/products/model"
	service "mercado-frescos-time-7/go-web/internal/products/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	service service.Service
}

func NewRepository(p service.Service) ProductHandler {
	return ProductHandler{
		service: p,
	}
}

func (ph *ProductHandler) GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		obj, err := ph.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno, tente mais tarde",
			})
			return
		}
		c.JSON(http.StatusOK, obj)
	}
}

func (ph *ProductHandler) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id inválido",
			})
			return
		}
		obj, err := ph.service.GetById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno, tente mais tarde",
			})
			return
		}
		c.JSON(http.StatusOK, obj)
	}
}

func (ph *ProductHandler) SaveProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		newProduct := saveProduct{}
		err := c.ShouldBindJSON(&newProduct)
		var ve validator.ValidationErrors
		fmt.Println(err)
		if errors.As(err, &ve) {
			for _, v := range ve {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("erro no campo: %v", v.Field()),
				})
				return
			}
		}
		p := model.Product{
			Product_code:                     newProduct.Product_code,
			Description:                      newProduct.Description,
			Width:                            newProduct.Width,
			Height:                           newProduct.Height,
			Length:                           newProduct.Length,
			Net_weight:                       newProduct.Net_weight,
			Expiration_rate:                  newProduct.Expiration_rate,
			Recommended_freezing_temperature: newProduct.Recommended_freezing_temperature,
			Freezing_rate:                    newProduct.Freezing_rate,
			Product_type_id:                  newProduct.Product_type_id,
			Seller_id:                        newProduct.Seller_id,
		}
		p, err = ph.service.Insert(p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno, tente mais tarde",
			})
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (ph *ProductHandler) UpdateProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		newProduct := saveProduct{}
		err := c.ShouldBindJSON(&newProduct)
		var ve validator.ValidationErrors
		fmt.Println(err)
		if errors.As(err, &ve) {
			for _, v := range ve {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("erro no campo: %v", v.Field()),
				})
				return
			}
		}
		p := model.Product{
			Product_code:                     newProduct.Product_code,
			Description:                      newProduct.Description,
			Width:                            newProduct.Width,
			Height:                           newProduct.Height,
			Length:                           newProduct.Length,
			Net_weight:                       newProduct.Net_weight,
			Expiration_rate:                  newProduct.Expiration_rate,
			Recommended_freezing_temperature: newProduct.Recommended_freezing_temperature,
			Freezing_rate:                    newProduct.Freezing_rate,
			Product_type_id:                  newProduct.Product_type_id,
			Seller_id:                        newProduct.Seller_id,
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id inválido",
			})
			return
		}
		err = ph.service.Update(id, p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno, tente mais tarde",
			})
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (ph *ProductHandler) DeleteProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id inválido",
			})
			return
		}
		err = ph.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro interno, tente mais tarde",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	}
}

type saveProduct struct {
	Product_code                     string  `json:"product_code"`
	Description                      string  `json:"description"`
	Width                            float64 `json:"width"`
	Height                           float64 `json:"height"`
	Length                           float64 `json:"lenght"`
	Net_weight                       float64 `json:"netweight"`
	Expiration_rate                  int     `json:"expiration_rate"`
	Recommended_freezing_temperature float64 `json:"recommended_freezing_temperature"`
	Freezing_rate                    float64 `json:"freezing_rate"`
	Product_type_id                  int     `json:"product_type_id" `
	Seller_id                        int     `json:"seller_id"`
}
