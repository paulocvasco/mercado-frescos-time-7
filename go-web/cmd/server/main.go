package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routers.InstanceSeller(r)

	r.Run()
}
