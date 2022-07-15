package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/routes"
	"mercado-frescos-time-7/go-web/pkg/logger"
	"mercado-frescos-time-7/go-web/pkg/web"
	"net/http"

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
	routes.InstanceProductBatch(r)
	routes.InstancePurchaseOrders(r)
	routes.InstanceProductRecords(r)
	routes.InstanceInboudOrders(r)
	routes.InstanceLocality(r)
	routes.InstanceReportSellers(r)
	routes.InstanceCarriers(r)

	instanceLog(r)
	r.Run()
}

func instanceLog(r *gin.Engine) {
	r.GET("/api/v1/logs/", controllerLog)
}

func controllerLog(c *gin.Context) {
	logs := logger.GetLogs()
	response := web.NewResponse(http.StatusOK, logs, "")
	c.JSON(http.StatusOK, response)
}
