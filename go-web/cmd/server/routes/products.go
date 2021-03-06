package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/products"
	"mercado-frescos-time-7/go-web/internal/products/repository"
	"mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceProducts(r *gin.Engine) {
	repo := repository.NewRepositoryMysql(db.StorageDB)
	serv := products.NewService(repo)
	pr := controller.NewProductHandler(serv)

	group := r.Group("/api/v1/products")
	{
		group.GET("/", pr.GetAllProducts())
		group.GET("/:id", pr.GetProduct())
		group.DELETE("/:id", pr.DeleteProducts())
		group.POST("/", pr.SaveProducts())
		group.PATCH("/:id", pr.UpdateProducts())
	}
}
