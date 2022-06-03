package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/warehouse"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := warehouse.NewRepository()
	service := warehouse.NewService(repository)
	controller := controller.NewController(service)

	routers := gin.Default()
	mapRouters(routers, controller)
	routers.Run()
}

func mapRouters(r *gin.Engine, c controller.Controller) {
	wr := r.Group("/api/v1/warehouses")
	wr.GET("", c.GetAll)
	wr.GET("/:id", c.GetByID)
	wr.POST("", c.Create)
	wr.PATCH("/:id", c.Update)
	wr.DELETE("/:id", c.Delete)
}
