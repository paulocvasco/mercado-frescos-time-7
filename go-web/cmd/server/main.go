package server

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(r *gin.Engine) error {

	if r == nil {

	}

	r.Use(cors.Default())
	routes.InstanceEmployee(r)
	routes.InstanceSeller(r)
	routes.InstanceBuyer(r)
	routes.InstanceProducts(r)
	routes.InstanceWarehouse(r)
	routes.InstanceSection(r)
	routes.InstanceProductBatch(r)
	routes.InstancePurchaseOrders(r)
	routes.InstanceProductRecords(r)
	routes.InstanceInboudOrders(r)
	routes.InstanceLocality(r)
	routes.InstanceReportSellers(r)
	routes.InstanceCarriers(r)

	return r.Run()
}
