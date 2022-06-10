package warehouse

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db"

	"github.com/google/uuid"
)

const path = "./warehouses.db"

type repository struct{}

var database []models.Warehouse

var cache models.WarehouseMetaData

var lastID int

type Repository interface {
	Create(models.Warehouse) (models.Warehouse, error)
	Update(int, models.Warehouse) error
	GetAll() ([]models.Warehouse, error)
	GetByID(int) (models.Warehouse, error)
	Delete(int) error
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(new models.Warehouse) (models.Warehouse, error) {
	var warehouses models.WarehouseMetaData
	err := readDB(&warehouses)
	if err != nil {
		return models.Warehouse{}, err
	}

	warehouses.LastID++
	new.ID = warehouses.LastID
	new.WarehouseCode = uuid.NewString()

	warehouses.Warehouses = append(warehouses.Warehouses, new)
	err = writeDB(warehouses)
	if err != nil {
		return models.Warehouse{}, err
	}

	cache = warehouses
	return new, nil
}

func (r *repository) Update(id int, patchedWarehouse models.Warehouse) error {
	if id < 0 || id > cache.LastID {
		return customerrors.ErrorInvalidID
	}

	for i, warehouse := range cache.Warehouses {
		if warehouse.ID == id {
			warehouse = models.Warehouse(patchedWarehouse)

			var warehouses models.WarehouseMetaData
			err := readDB(&warehouses)
			if err != nil {
				return err
			}

			warehouses.Warehouses[i] = warehouse

			err = writeDB(warehouses)
			if err != nil {
				return err
			}

			cache = warehouses
			return nil
		}
	}
	return customerrors.ErrorItemNotFound
}

func (r *repository) GetAll() ([]models.Warehouse, error) {
	if cache.LastID == 0 {
		err := readDB(&cache)
		if err != nil {
			return nil, err
		}
	}
	return cache.Warehouses, nil
}

func (r *repository) GetByID(id int) (models.Warehouse, error) {
	if cache.LastID == 0 {
		err := readDB(&cache)
		if err != nil {
			return models.Warehouse{}, err
		}
	}

	if id < 0 || id > cache.LastID {
		return models.Warehouse{}, customerrors.ErrorInvalidID
	}

	for _, w := range cache.Warehouses {
		if w.ID == id {
			return w, nil
		}
	}
	return models.Warehouse{}, customerrors.ErrorInvalidID
}

func (r *repository) Delete(id int) error {
	if id < 0 || id > cache.LastID {
		return customerrors.ErrorInvalidID
	}

	for index, warehouse := range cache.Warehouses {
		if warehouse.ID == id {

			var warehouses models.WarehouseMetaData
			err := readDB(&warehouses)
			if err != nil {
				return err
			}

			warehouses.Warehouses = append(warehouses.Warehouses[:index], warehouses.Warehouses[index+1:]...)

			err = writeDB(warehouses)
			if err != nil {
				return err
			}
			cache = warehouses
			return nil
		}
	}
	return customerrors.ErrorInvalidID
}
