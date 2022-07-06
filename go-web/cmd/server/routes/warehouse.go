package routes

import (
	"github.com/gin-gonic/gin"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/warehouse"
	"mercado-frescos-time-7/go-web/internal/warehouse/repository"
)

func InstanceWarehouse(r *gin.Engine) {
	repository := repository.NewSqlRepository()
	service := warehouse.NewService(repository)
	controller := controller.NewControllerWarehouse(service)

	wr := r.Group("/api/v1/warehouses")
	wr.GET("/", controller.GetAllWarehouse)
	wr.GET("/:id", controller.GetByIDWarehouse)
	wr.POST("/", controller.CreateWarehouse)
	wr.PATCH("/:id", controller.UpdateWarehouse)
	wr.DELETE("/:id", controller.DeleteWarehouse)

}
