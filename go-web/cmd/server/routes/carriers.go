package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/carriers"
	"mercado-frescos-time-7/go-web/internal/carriers/repository"

	"github.com/gin-gonic/gin"
)

func InstanceCarriers(r *gin.Engine) {
	repo := repository.NewRepository()
	s := carriers.NewService(repo)
	ctrl := controller.NewControllerCarriers(s)

	cr := r.Group("/api/v1/")
	cr.POST("carriers/", ctrl.CreateCarrier)
	cr.GET("localities/reportCarriers/", ctrl.GetCarriers)
}
