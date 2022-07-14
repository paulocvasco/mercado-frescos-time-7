package routes

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	seller "github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/Seller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceSeller(e *gin.Engine) {
	mydb := db.NewDatabase()
	repo := seller.NewRepository(mydb)
	service := seller.NewService(repo)
	p := controller.NewSellers(service)

	r := e.Group("api/v1")
	r.GET("/sellers", p.SellersGetAll())
	r.GET("/sellers/:id", p.SellersGetId())
	r.POST("/sellers", p.SellersStore())
	r.PATCH("/sellers/:id", p.SellersUpdate())
	r.DELETE("/sellers/:id", p.SellersDelete())

}
