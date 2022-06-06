package main

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/products/repository"
	service "mercado-frescos-time-7/go-web/internal/products/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repo := repository.NewRepository()
	serv := service.NewService(repo)
	pr := controller.NewRepository(serv)

	group := r.Group("/products")
	{
		group.GET("/", pr.GetAllProducts())
		group.GET("/:id", pr.GetProduct())
		group.DELETE("/:id", pr.DeleteProducts())
		group.POST("/", pr.SaveProducts())
		group.PATCH("/:id", pr.UpdateProducts())
	}

	r.Run()
}
