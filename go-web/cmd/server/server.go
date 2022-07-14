package server

import (
	"mercado-frescos-time-7/go-web/cmd/server/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	r := gin.Default()

	routes.InstanceEmployee(r)
	routes.InstanceSeller(r)
	routes.InstanceBuyer(r)
	routes.InstanceProducts(r)
	routes.InstanceWarehouse(r)
	routes.InstanceSection(r)
	r.Run()
}
