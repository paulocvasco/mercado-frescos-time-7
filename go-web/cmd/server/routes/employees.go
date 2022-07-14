package routes

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/employees"

	"github.com/gin-gonic/gin"
)

func InstanceEmployee(e *gin.Engine) {
	repository := employees.NewRepository()
	service := employees.NewService(repository)
	controller := controller.NewEmployee(service)

	r := e.Group("api/v1/employees")
	r.GET("", controller.GetAll())
	r.GET("/:id", controller.GetByID())
	r.POST("", controller.Create())
	r.DELETE("/:id", controller.Delete())
	r.PATCH("/:id", controller.Update())

}
