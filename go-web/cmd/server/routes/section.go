package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/sections"
	"mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceSection(eng *gin.Engine) {
	database := db.NewDatabase()
	repository := sections.NewRepository(database)
	services := sections.NewService(repository)
	controller := controller.NewController(services)
	sec := eng.Group("/api/v1/sections")
	sec.GET("/", controller.GetAll)
	sec.GET("/:id", controller.GetById)
	sec.POST("/", controller.Store)
	sec.PATCH("/:id", controller.Update)
	sec.DELETE("/:id", controller.Delete)
}
