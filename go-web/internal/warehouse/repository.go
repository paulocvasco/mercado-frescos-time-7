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

type Warehouse models.Warehouse

var database []Warehouse

var cache models.WarehouseMetaData

var lastID int

type Repository interface {
	Create(models.Warehouse) models.Warehouse
	Update(int, models.Warehouse) error
	GetAll() []models.Warehouse
	GetByID(int) (models.Warehouse, error)
	Delete(int) error
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(new models.Warehouse) models.Warehouse {
	tmpCache := cache
	tmpCache.LastID++
	new.ID = tmpCache.LastID
	new.WarehouseCode = uuid.NewString()

	tmpCache.Warehouses = append(tmpCache.Warehouses, models.Warehouse(new))

	rawWarehouses, err := json.Marshal(tmpCache)
	if err != nil {
		return Warehouse{}, err
	}
	err = db.Save(path, rawWarehouses)
	if err != nil {
		return Warehouse{}, err
	}

	cache = tmpCache
	return new, nil
}

func (r *repository) Update(id int, patchedWarehouse models.Warehouse) error {
	if id < 0 || id > lastID {
		return customerrors.ErrorInvalidID
	}

	for i, warehouse := range database {
		if warehouse.ID == id {
			warehouse = patchedWarehouse
			database[i] = warehouse
			return nil
		}
	}
	return customerrors.ErrorItemNotFound
}

func (r *repository) GetAll() []models.Warehouse {
	return database
}

func (r *repository) GetByID(id int) (models.Warehouse, error) {
	if id < 0 || id > lastID {
		return models.Warehouse{}, customerrors.ErrorInvalidID
	}

	for _, w := range database {
		if w.ID == id {
			return w, nil
		}
	}
	return models.Warehouse{}, customerrors.ErrorInvalidID
}

func (r *repository) Delete(id int) error {
	if id < 0 || id > lastID {
		return customerrors.ErrorInvalidID
	}

	for index, warehouse := range database {
		if warehouse.ID == id {
			database = append(database[:index], database[index+1:]...)
			return nil
		}
	}
	return customerrors.ErrorInvalidID
}
