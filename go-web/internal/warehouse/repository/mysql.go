package repository

import (
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/pkg/db/mysql"
)

type mysqlDB struct {
	database sql.DB
}

func NewSqlRepository() Repository {
	db := mysql.Get()
	return &mysqlDB{
		database: db}
}

func (m *mysqlDB) Create(new models.Warehouse) (models.Warehouse, error) {
	return models.Warehouse{}, nil
}

func (m *mysqlDB) Update(id int, patchedWarehouse models.Warehouse) error {
	return nil
}

func (m *mysqlDB) GetAll() (models.Warehouses, error) {
	return models.Warehouses{}, nil
}

func (m *mysqlDB) GetByID(id int) (models.Warehouse, error) {
	return models.Warehouse{}, nil
}

func (m *mysqlDB) Delete(id int) error {
	return nil
}
