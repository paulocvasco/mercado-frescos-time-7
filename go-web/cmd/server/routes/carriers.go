package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/carriers"
	"mercado-frescos-time-7/go-web/internal/carriers/repository"
	"mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceCarriers(r *gin.Engine) {
	db := db.Get()
	repo := repository.NewRepository(db)
	s := carriers.NewService(repo)
	ctrl := controller.NewControllerCarriers(s)

	cr := r.Group("/api/v1/")
	cr.POST("carriers/", ctrl.CreateCarrier)
	cr.GET("localities/reportCarriers/", ctrl.GetCarriers)
}
