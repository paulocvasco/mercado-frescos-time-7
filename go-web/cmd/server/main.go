package main

import (
	routes "mercado-frescos-time-7/go-web/cmd/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.InstanceBayer(r)
	routes.InstanceProducts(r)
	// routes.InstanceWarehouse(r) s

	r.Run()
}
