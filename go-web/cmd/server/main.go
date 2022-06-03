package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	seller "mercado-frescos-time-7/go-web/internal/Seller"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := seller.NewRepository()
	service := seller.NewService(repo)
	p := controller.NewSellers(service)

	r := gin.Default()
	r.GET("/sellers", p.GetAll())
	r.GET("/sellers/:id", p.GetId())
	r.POST("/sellers", p.Store())
	r.PATCH("/sellers/:id", p.Update())
	r.DELETE("/sellers/:id", p.Delete())
	r.Run()
}
