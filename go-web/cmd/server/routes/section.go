package routes

import (
	"github.com/gin-gonic/gin"
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/sections"
	"mercado-frescos-time-7/go-web/internal/sections/repository"
	"mercado-frescos-time-7/go-web/pkg/db"
)

func InstanceSection(eng *gin.Engine) {
	repository := repository.NewRepositorySection(db.StorageDB)
	services := sections.NewServiceSection(repository)
	controller := controller.NewController(services)
	sec := eng.Group("/api/v1/sections")
	sec.GET("/", controller.GetAll)
	sec.GET("/:id", controller.GetById)
	sec.POST("/", controller.Store)
	sec.PATCH("/:id", controller.Update)
	sec.DELETE("/:id", controller.Delete)
	sec.GET("/reportProducts", controller.GetAllReportProducts)
	sec.GET("/reportProducts/:id", controller.GetReportProductsById)

}
