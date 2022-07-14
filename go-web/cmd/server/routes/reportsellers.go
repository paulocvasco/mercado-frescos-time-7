package routes

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/controller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/reportsellers"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/reportsellers/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceReportSellers(e *gin.Engine) {
	Mysqlrepo := repository.NewSQLrepository(db.StorageDB)
	service := reportsellers.NewService(Mysqlrepo)
	p := controller.NewReportSellers(service)

	r := e.Group("api/v1")

	r.GET("/localities/reportSellers", p.ReportSellers())
}
