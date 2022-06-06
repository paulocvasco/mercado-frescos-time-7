package main

import (
	routers "mercado-frescos-time-7/go-web/cmd/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.InstanceBayer(r)

	r.Run()
}
