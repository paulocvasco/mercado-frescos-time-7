package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	productrecords "mercado-frescos-time-7/go-web/internal/product_records"
	"mercado-frescos-time-7/go-web/internal/product_records/repository"
	"mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceProductRecords(r *gin.Engine) {
	repo := repository.NewRepositoryProductRecord(db.StorageDB)
	serv := productrecords.NewServiceProductRecord(repo)
	recordController := controller.NewProductRecordsController(serv)

	group := r.Group("/api/v1")
	{
		group.GET("/products/reportRecords", recordController.GetProductRecordsById())
		group.POST("/productRecords", recordController.InsertProductRecord())
	}
}
