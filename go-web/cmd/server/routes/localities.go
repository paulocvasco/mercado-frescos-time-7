package routes

import (
	"mercado-frescos-time-7/go-web/cmd/server/controller"
	"mercado-frescos-time-7/go-web/internal/locality"
	"mercado-frescos-time-7/go-web/internal/locality/repository"
	"mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-gonic/gin"
)

func InstanceLocality(e *gin.Engine) {
	Mysqlrepo := repository.NewSQLrepository(db.StorageDB)
	service := locality.NewService(Mysqlrepo)
	p := controller.NewLocality(service)

	r := e.Group("api/v1") 

	r.POST("/localities/", p.LocalityStore())
}