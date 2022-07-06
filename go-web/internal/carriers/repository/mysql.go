package repository

import (
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/pkg/db"
)

type Repository interface {
	Create(models.Carrier) (models.Carrier, error)
	Get(int) (models.CarriersReport, error)
}

type mysqlDB struct {
	db *sql.DB
}

func NewRepository() Repository {
	database := db.Get()
	return &mysqlDB{
		db: database,
	}
}

func (m *mysqlDB) Create(new models.Carrier) (models.Carrier, error) {
	return models.Carrier{}, nil
}

func (m *mysqlDB) Get(id int) (models.CarriersReport, error) {
	return models.CarriersReport{}, nil
}
