package routes

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/purchaseOrders"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstancePurchaseOrders(r *gin.Engine) {

	repo := repository.NewRepositoryMySql(db.StorageDB) //(database)
	service := purchaseOrders.NewService(repo)
	c := controller.PurchaseOrdersNewController(service)

	routes := r.Group("/api/v1/purchaseOrders")
	r.GET("/api/v1/buyers/reportPurchaseOrders", c.GetPurchaseOrder())
	routes.POST("/", c.CreatePurchaseOrders())

}
