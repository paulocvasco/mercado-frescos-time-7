package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/inbound_orders"
	repository2 "github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/inbound_orders/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"
)

func InstanceInboudOrders(i *gin.Engine) {
	repository := repository2.NewRepository(db.StorageDB)
	service := inbound_orders.NewService(repository)
	controller := controller.NewInboundOrders(service)

	r := i.Group("api/v1/inboundOrders")
	r.POST("/", controller.Create())

}
