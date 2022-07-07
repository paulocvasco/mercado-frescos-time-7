package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/reportsellers"
	"mercado-frescos-time-7/go-web/internal/reportsellers/repository"
	"mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceReportSellers(e *gin.Engine) {
	Mysqlrepo := repository.NewSQLrepository(db.StorageDB)
	service := reportsellers.NewService(Mysqlrepo)
	p := controller.NewReportSellers(service)

	r := e.Group("api/v1") 

	r.GET("/localities/reportSellers", p.ReportSellers())
}