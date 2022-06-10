package controller

import (
	"mercado-frescos-time-7/go-web/internal/login"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	service login.Service
}

func LoginNewController(l login.Service) *LoginController {
	return &LoginController{
		service: l,
	}
}

func (l *LoginController) CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input createUser
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		user := l.service.CreateUser(input.User, input.Password)
		context.JSON(http.StatusCreated, user)

	}
}

func (l *LoginController) GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input createUser
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		user, err := l.service.GetUser(input.User, input.Password)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "User e Password invalidos"})
			return
		}

		context.JSON(http.StatusCreated, user)

	}
}

type createUser struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
