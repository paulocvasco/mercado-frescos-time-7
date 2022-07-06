package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/productBatch"
	"mercado-frescos-time-7/go-web/internal/productBatch/repository"
)

func InstanceProductBatch(eng *gin.Engine) {
	var database *sql.DB
	repository := repository.NewRepository(database)
	services := productBatch.NewService(repository)
	controller := controller.NewControllerProductBatch(services)
	pb := eng.Group("/api/v1/productBatches")

	pb.POST("/", controller.Store)
}
