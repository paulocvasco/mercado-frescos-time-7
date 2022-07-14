package routes

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/buyer"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/buyer/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceBuyer(r *gin.Engine) {
	// database := db.NewDatabase()
	// dataBase, err := db.ConectionDb()

	// if err != nil {
	//	log.Fatal("failed to connect to mariadb")
	// }
	repo := repository.NewRepositoryMySql(db.StorageDB) //(database)
	service := buyer.NewService(repo)
	c := controller.BuyerNewController(service)

	routes := r.Group("/api/v1/buyers")

	routes.GET("/", c.BuyerGetAll())
	routes.GET("/:id", c.BuyerGetId())
	routes.POST("/", c.BuyerCreate())
	routes.PATCH("/:id", c.BuyerUpdate())
	routes.DELETE("/:id", c.BuyerDelete())

}
