package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/buyer"

	"github.com/gin-gonic/gin"
)

func InstanceBuyer(r *gin.Engine) {
	repo := buyer.NewRepository()
	service := buyer.NewService(repo)
	c := controller.BuyerNewController(service)

	routes := r.Group("/api/v1/buyers")

	routes.GET("/", c.BuyerGetAll())
	routes.GET("/:id", c.BuyerGetId())
	routes.POST("/", c.BuyerCreate())
	routes.PATCH("/:id", c.BuyerUpdate())
	routes.DELETE("/:id", c.BuyerDelete())
}
