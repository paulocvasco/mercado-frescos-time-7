package routes

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/warehouse"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/warehouse/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceWarehouse(r *gin.Engine) {
	db := db.Get()
	repository := repository.NewSqlRepository(db)
	service := warehouse.NewService(repository)
	controller := controller.NewControllerWarehouse(service)

	wr := r.Group("/api/v1/warehouses")
	wr.GET("/", controller.GetAllWarehouse)
	wr.GET("/:id", controller.GetByIDWarehouse)
	wr.POST("/", controller.CreateWarehouse)
	wr.PATCH("/:id", controller.UpdateWarehouse)
	wr.DELETE("/:id", controller.DeleteWarehouse)

}
