package routers

import (
	server "mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/buyer"

	"github.com/gin-gonic/gin"
)

func InstanceBayer(r *gin.Engine) {
	repo := buyer.NewRepository()
	service := buyer.NewService(repo)
	b := server.NewController(service)

	routes := r.Group("/api/v1/buyers")

	routes.GET("/", b.GetAll())
	routes.GET("/:id", b.GetId())
	routes.POST("/", b.Creat())
	routes.PATCH("/:id", b.Update())
	routes.DELETE("/:id", b.Delete())
}
