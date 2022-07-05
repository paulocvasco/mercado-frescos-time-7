package repository

import (
	"database/sql"
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
	query := "INSERT INTO warehouse(address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (? ? ? ? ?)"
	res, err := m.db.Exec(query, new.Address, new.Telephone, new.WarehouseCode, new.MinimunCapacity, new.MinimunTemperature)
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
	query := "UPDATE warehouse SET address = ?, telephone = ?, warehouse_code = ?, minimum_capacity = ?, minimum_temperature = ?, where id = ?"
	_, err := m.db.Exec(query, patchedWarehouse.Address, patchedWarehouse.Telephone, patchedWarehouse.WarehouseCode, patchedWarehouse.MinimunCapacity, patchedWarehouse.MinimunTemperature, id)
	if err != nil {
		return err
	}
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
