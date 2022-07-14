package server

import (
	"database/sql"

	"github.com/paulocvasco/mercado-frescos-time-7/go-web/cmd/server/routes"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(r *gin.Engine, connDB *sql.DB) error {
	if r == nil {
		return customerrors.ErrorInstaceGin
	}

	err := db.InstanceDB(connDB)
	if err != nil {
		return customerrors.ErrorInvalidDB
	}

	r.Use(cors.Default())
	routes.InstanceEmployee(r)
	routes.InstanceSeller(r)
	routes.InstanceBuyer(r)
	routes.InstanceProducts(r)
	routes.InstanceWarehouse(r)
	routes.InstanceSection(r)
	routes.InstanceProductBatch(r)
	routes.InstancePurchaseOrders(r)
	routes.InstanceProductRecords(r)
	routes.InstanceInboudOrders(r)
	routes.InstanceLocality(r)
	routes.InstanceReportSellers(r)
	routes.InstanceCarriers(r)

	return r.Run()
}
