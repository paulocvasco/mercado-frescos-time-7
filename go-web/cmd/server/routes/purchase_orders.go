package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/purchaseOrders"
	"mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
	"mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstancePurchaseOrders(r *gin.Engine) {

	repo := repository.NewRepositoryMySql(db.StorageDB) //(database)
	service := purchaseOrders.NewService(repo)
	c := controller.PurchaseOrdersNewController(service)

	routes := r.Group("/api/v1/purchaseOrders")

	routes.POST("/", c.CreatePurchaseOrders())

}
