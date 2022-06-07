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

	r.GET("/sellers", p.SeellersGetAll())
	r.GET("/sellers/:id", p.SeellersGetId())
	r.POST("/sellers", p.SeellersStore())
	r.PATCH("/sellers/:id", p.SeellersUpdate())
	r.DELETE("/sellers/:id", p.SeellersDelete())

}

