package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.InstanceWarehouse(r)

	r.Run()
}
