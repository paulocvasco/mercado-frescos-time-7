package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.InstanceEmployee(r)
	routes.InstanceSeller(r)
	routes.InstanceBuyer(r)
	routes.InstanceProducts(r)
	routes.InstanceWarehouse(r)
	routes.InstanceLogin(r)

	r.Run()
}
