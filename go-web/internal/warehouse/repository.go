package warehouse

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"

	"github.com/google/uuid"
)

type repository struct{}

var db []models.Warehouse

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
	new.ID = lastID + 1
	new.WarehouseCode = uuid.NewString()

	db = append(db, new)
	lastID++

	return new
}

func (r *repository) Update(id int, patchedWarehouse models.Warehouse) error {
	if id < 0 || id > lastID {
		return customerrors.ErrorInvalidID
	}

	for i, warehouse := range db {
		if warehouse.ID == id {
			warehouse = patchedWarehouse
			db[i] = warehouse
			return nil
		}
	}
	return customerrors.ErrorItemNotFound
}

func (r *repository) GetAll() []models.Warehouse {
	return db
}

func (r *repository) GetByID(id int) (models.Warehouse, error) {
	if id < 0 || id > lastID {
		return models.Warehouse{}, customerrors.ErrorInvalidID
	}

	for _, w := range db {
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

	for index, warehouse := range db {
		if warehouse.ID == id {
			db = append(db[:index], db[index+1:]...)
			return nil
		}
	}
	return customerrors.ErrorInvalidID
}
