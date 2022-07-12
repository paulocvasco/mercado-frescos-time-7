package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/employees"
	repository2 "mercado-frescos-time-7/go-web/internal/employees/repository"
	"mercado-frescos-time-7/go-web/pkg/db"
)

func InstanceEmployee(e *gin.Engine) {
	repository := repository2.NewRepository(db.StorageDB)
	service := employees.NewService(repository)
	controller := ctrl.NewEmployee(service)

	repositoryReport := repository2.NewRepositoryReport(db.StorageDB)
	serviceReport := employees.NewServiceReport(repositoryReport)
	controllerReport := ctrl.NewReport(serviceReport)

	r := e.Group("api/v1/employees")
	r.GET("", controller.GetAll())
	r.GET("/:id", controller.GetByID())
	r.POST("", controller.Create())
	r.DELETE("/:id", controller.Delete())
	r.PATCH("/:id", controller.Update())
	r.GET("/reportInboundOrders", controllerReport.GetReportInboundOrders())

}
