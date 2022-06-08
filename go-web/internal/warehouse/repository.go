package warehouse

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"

	"github.com/google/uuid"
)

type repository struct{}

type Warehouse models.Warehouse

var db []Warehouse

var lastID int

type Repository interface {
	Create(Warehouse) Warehouse
	Update(int, Warehouse) error
	GetAll() []Warehouse
	GetByID(int) (Warehouse, error)
	Delete(int) error
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(new Warehouse) Warehouse {
	new.ID = lastID + 1
	new.WarehouseCode = uuid.NewString()

	db = append(db, new)
	lastID++

	return new
}

func (r *repository) Update(id int, patchedWarehouse Warehouse) error {
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

func (r *repository) GetAll() []Warehouse {
	return db
}

func (r *repository) GetByID(id int) (Warehouse, error) {
	if id < 0 || id > lastID {
		return Warehouse{}, customerrors.ErrorInvalidID
	}

	for _, w := range db {
		if w.ID == id {
			return w, nil
		}
	}
	return Warehouse{}, customerrors.ErrorInvalidID
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
