package controller

import (
	"mercado-frescos-time-7/go-web/internal/warehouse/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetToken(*gin.Context)
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
}

type controller struct {
	service services.Service
}

func (control *controller) GetAll(c *gin.Context) {
	response, err := control.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// func (control *controller) GetByID(c *gin.Context) {
//	param := c.Param("id")
//	response, err := control.service.GetByID(param)
//	if err != nil {
//		switch err {
//		case customErrors.ErrorInvalidIDParameter:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorInvalidID:
//			c.JSON(http.StatusNotFound, err)
//		default:
//			c.JSON(http.StatusInternalServerError, err)
//		}
//		return
//	}
//	c.JSON(http.StatusOK, response)
// }

// func (control *controller) Store(c *gin.Context) {
//	body := c.Request.Body
//	defer body.Close()

//	data, err := ioutil.ReadAll(body)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}

//	err = control.service.Store(data)
//	if err != nil {
//		switch err {
//		case customErrors.ErrorMissingCode:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorMissingCoin:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorMissingIssuer:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorMissingReciever:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorMissingValue:
//			c.JSON(http.StatusBadRequest, err)
//		default:
//			c.JSON(http.StatusInternalServerError, err)
//		}
//		return
//	}
//	c.JSON(http.StatusOK, nil)
// }

// func (control *controller) Update(c *gin.Context) {
//	id := c.Param("id")
//	body := c.Request.Body
//	defer body.Close()

//	data, err := ioutil.ReadAll(body)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}

//	err = control.service.Update(id, data)
//	if err != nil {
//		switch err {
//		case customErrors.ErrorMissingCode:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorMissingCoin:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorMissingIssuer:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorMissingReciever:
//			c.JSON(http.StatusBadRequest, err)
//		case customErrors.ErrorMissingValue:
//			c.JSON(http.StatusBadRequest, err)
//		default:
//			c.JSON(http.StatusInternalServerError, err)
//		}
//		return
//	}
//	c.JSON(http.StatusOK, nil)
// }
