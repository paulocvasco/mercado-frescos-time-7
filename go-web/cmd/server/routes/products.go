package routes

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/products"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceProducts(r *gin.Engine) {
	db := db.NewDatabase()
	repo := products.NewRepository(db)
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
