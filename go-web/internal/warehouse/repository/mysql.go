package repository

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db"
)

type mysqlDB struct {
	db *sql.DB
}

func NewSqlRepository() Repository {
	database := db.Get()
	return &mysqlDB{
		db: database}
}

func (m *mysqlDB) Create(new models.Warehouse) (models.Warehouse, error) {
	query := "INSERT INTO warehouse(address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := m.db.Exec(query, new.Address, new.Telephone, new.WarehouseCode, new.MinimunCapacity, new.MinimunTemperature, new.LocalityID)
	if err != nil {
		log.Println(err)
		return models.Warehouse{}, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return models.Warehouse{}, err
	}
	newWarehouse := new
	newWarehouse.ID = int(lastID)
	return newWarehouse, nil
}

func (m *mysqlDB) Update(id int, w models.Warehouse) error {
	query := "UPDATE warehouse SET address = ?, telephone = ?, warehouse_code = ?, minimum_capacity = ?, minimum_temperature = ?, locality_id = ?, WHERE id = ?"
	_, err := m.db.Exec(query, w.Address, w.Telephone, w.WarehouseCode, w.MinimunCapacity, w.MinimunTemperature, w.LocalityID, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (m *mysqlDB) GetAll() (models.Warehouses, error) {
	res, err := m.db.Query("SELECT * FROM warehouse")
	if err != nil {
		log.Println(err)
		return models.Warehouses{}, err
	}

	var all models.Warehouses
	for res.Next() {
		var w models.Warehouse
		err := res.Scan(&w.ID, &w.Address, &w.MinimunCapacity, &w.MinimunCapacity, &w.MinimunTemperature, &w.WarehouseCode)
		if err != nil {
			log.Println(err)
			return models.Warehouses{}, err
		}
		all.Warehouses = append(all.Warehouses, w)
	}
	return all, nil
}

func (m *mysqlDB) GetByID(id int) (models.Warehouse, error) {
	query := "SELECT * FROM warehouse WHERE id = ?"
	var w models.Warehouse
	err := m.db.QueryRow(query, id).Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimunCapacity, &w.MinimunTemperature)
	if err != nil {
		return models.Warehouse{}, customerrors.ErrorItemNotFound
	}

	return w, nil
}

func (m *mysqlDB) Delete(id int) error {
	query := "DELETE FROM warehouse WHERE id = ?"
	err := m.db.QueryRow(query, id).Scan()
	if err != nil {
		return customerrors.ErrorItemNotFound
	}
	return nil
}
