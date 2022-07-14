package server

import (
	"github.com/gin-gonic/gin"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/routes"
)

func main() {

	r := gin.Default()

	routes.InstanceEmployee(r)
	routes.InstanceSeller(r)
	routes.InstanceBuyer(r)
	routes.InstanceProducts(r)
	routes.InstanceWarehouse(r)
	routes.InstanceSection(r)
	r.Run()
}
