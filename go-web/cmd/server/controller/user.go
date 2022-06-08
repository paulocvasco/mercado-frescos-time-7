package controller

import (
	"mercado-frescos-time-7/go-web/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	NewUser(*gin.Context)
	GetToken(*gin.Context)
}

type userController struct {
	service user.Service
}

func NewController(s user.Service) UserController {
	return &userController{
		service: s,
	}
}

func (u *userController) NewUser(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (u *userController) GetToken(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
