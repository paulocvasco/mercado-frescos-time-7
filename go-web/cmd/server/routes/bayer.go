package routes

import (
	server "mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/buyer"

	"github.com/gin-gonic/gin"
)

func InstanceBayer(r *gin.Engine) {
	repo := buyer.NewRepository()
	service := buyer.NewService(repo)
	b := server.BuyerNewController(service)

	routes := r.Group("/api/v1/buyers")

	routes.GET("/", b.BuyerGetAll())
	routes.GET("/:id", b.BuyerGetId())
	routes.POST("/", b.BuyerCreat())
	routes.PATCH("/:id", b.BuyerUpdate())
	routes.DELETE("/:id", b.BuyerDelete())
}
