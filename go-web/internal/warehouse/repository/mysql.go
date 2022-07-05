package repository

import (
	"database/sql"
	"fmt"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/pkg/db/mysql"
)

type mysqlDB struct {
	db sql.DB
}

func NewSqlRepository() Repository {
	database := mysql.Get()
	return &mysqlDB{
		db: database}
}

func (m *mysqlDB) Create(new models.Warehouse) (models.Warehouse, error) {
	arg := fmt.Sprintf("INSERT INTO warehouse(address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES ('%s', '%s', '%s', %d, %d)", new.Address, new.Telephone, new.WarehouseCode, new.MinimunCapacity, new.MinimunTemperature)
	res, err := m.db.Exec(arg)
	if err != nil {
		return models.Warehouse{}, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return models.Warehouse{}, err
	}
	newWarehouse := new
	newWarehouse.ID = int(lastID)
	return newWarehouse, nil
}

func (m *mysqlDB) Update(id int, patchedWarehouse models.Warehouse) error {
	return nil
}

func (m *mysqlDB) GetAll() (models.Warehouses, error) {
	res, err := m.db.Query("SELECT * FROM warehouse")
	if err != nil {
		return models.Warehouses{}, err
	}

	var all models.Warehouses
	for res.Next() {
		var w models.Warehouse
		err := res.Scan(&w.ID, &w.Address, &w.MinimunCapacity, &w.MinimunCapacity, &w.MinimunTemperature, &w.WarehouseCode)
		if err != nil {
			return models.Warehouses{}, err
		}
		all.Warehouses = append(all.Warehouses, w)
	}
	return all, nil
}

func (m *mysqlDB) GetByID(id int) (models.Warehouse, error) {
	return models.Warehouse{}, nil
}

func (m *mysqlDB) Delete(id int) error {
	return nil
}
