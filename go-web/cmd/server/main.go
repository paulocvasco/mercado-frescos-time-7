package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(cors.Default())
	routes.InstanceEmployee(r)
	routes.InstanceSeller(r)
	routes.InstanceBuyer(r)
	routes.InstanceProducts(r)
	routes.InstanceWarehouse(r)
	routes.InstanceSection(r)
	routes.InstanceProductRecords(r)
	routes.InstanceInboudOrders(r)
	routes.InstanceLocality(r)
	routes.InstanceReportSellers(r)
	routes.InstanceCarriers(r)
	r.Run()
}
