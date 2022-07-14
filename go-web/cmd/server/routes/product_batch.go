package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/productBatch"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/productBatch/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"
)

func InstanceProductBatch(eng *gin.Engine) {
	repository := repository.NewRepositoryProductBatch(db.StorageDB)
	services := productBatch.NewService(repository)
	controller := controller.NewControllerProductBatch(services)
	pb := eng.Group("/api/v1/productBatches")

	pb.POST("/", controller.Store)
}
