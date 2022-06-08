package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	seller "mercado-frescos-time-7/go-web/internal/Seller"

	"github.com/gin-gonic/gin"
)


func InstanceSeller(r *gin.Engine) {
	repo := seller.NewRepository()
	service := seller.NewService(repo)
	p := controller.NewSellers(service)

	r.GET("/sellers", p.SellersGetAll())
	r.GET("/sellers/:id", p.SellersGetId())
	r.POST("/sellers", p.SellersStore())
	r.PATCH("/sellers/:id", p.SellersUpdate())
	r.DELETE("/sellers/:id", p.SellersDelete())

}

