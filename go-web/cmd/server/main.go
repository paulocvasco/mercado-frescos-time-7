package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/auth"
	"mercado-frescos-time-7/go-web/cmd/server/routes"
	"mercado-frescos-time-7/go-web/internal/login"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := login.NewRepository()
	service := login.NewService(repo)
	m := auth.Middeleware(service)

	r := gin.Default()
	routes.InstanceLogin(r)
	r.Use(m.ValidateToken())
	routes.InstanceEmployee(r)
	routes.InstanceSeller(r)
	routes.InstanceBuyer(r)
	routes.InstanceProducts(r)
	routes.InstanceWarehouse(r)

	r.Run()
}
