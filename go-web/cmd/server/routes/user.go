package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/user"

	"github.com/gin-gonic/gin"
)

func InstanceUser(r *gin.Engine) {
	repo := user.NewRpository()
	service := user.NewService(repo)
	control := controller.NewUserController(service)

	ur := r.Group("/api/v1/users")
	ur.POST("/", control.NewUser)
	ur.GET("/token", control.GetToken)
}
