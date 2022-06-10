package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/login"

	"github.com/gin-gonic/gin"
)

func InstanceLogin(r *gin.Engine) {
	repo := login.NewRepository()
	service := login.NewService(repo)
	c := controller.LoginNewController(service)

	routes := r.Group("/login")

	routes.POST("/", c.CreateUser())

}
