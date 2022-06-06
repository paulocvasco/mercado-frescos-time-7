package routers

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/warehouse"

	"github.com/gin-gonic/gin"
)

func InstanceWarehouse(r *gin.Engine) {
	repository := warehouse.NewRepository()
	service := warehouse.NewService(repository)
	controller := controller.NewController(service)

	wr := r.Group("/api/v1/warehouses")
	wr.GET("", controller.GetAll)
	wr.GET("/:id", controller.GetByID)
	wr.POST("", controller.Create)
	wr.PATCH("/:id", controller.Update)
	wr.DELETE("/:id", controller.Delete)

}
