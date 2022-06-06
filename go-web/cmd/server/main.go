package main

import (
	"mercado-frescos-time-7/go-web/cmd/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.InstanceBayer(r)
	r.Run()
}
