package warehouse

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	customerrors "mercado-frescos-time-7/go-web/internal/custom_errors"
	"mercado-frescos-time-7/go-web/internal/models"
	"strconv"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
)

type Warehouses models.Warehouses

type Warehouse models.Warehouse

var repository Warehouses

var lastID int

type Repository interface {
	Create(Warehouse)
	Update(int, []byte) error
	GetAll() Warehouses
	GetByID(int) (Warehouse, error)
	Delete(int) error
}

func NewRepository() Repository {
	return &repository
}

func (w *Warehouses) Create(new Warehouse) {
	new.ID = lastID + 1
	new.WarehouseCode = calculateCode(new.Address, new.Telephone, strconv.Itoa(new.MinimunCapacity), strconv.Itoa(new.MinimunTemperature))

	w.Warehouse = append(w.Warehouse, models.Warehouse(new))
	lastID++
}

func (w *Warehouses) Update(id int, data []byte) error {
	if id < 0 || id > lastID {
		return customerrors.ErrorInvalidID
	}

	for i, warehouse := range w.Warehouse {
		if warehouse.ID == id {
			warehouseBytes, err := json.Marshal(warehouse)
			if err != nil {
				return err
			}
			patchedWarehouse, err := jsonpatch.MergePatch(warehouseBytes, data)
			if err != nil {
				return err
			}
			err = json.Unmarshal(patchedWarehouse, &warehouse)
			if err != nil {
				return err
			}
			warehouse.WarehouseCode = calculateCode(warehouse.Address, warehouse.Telephone, strconv.Itoa(warehouse.MinimunCapacity), strconv.Itoa(warehouse.MinimunTemperature))
			w.Warehouse[i] = warehouse
			return nil
		}
	}
	return customerrors.ErrorItemNotFound
}

func (w *Warehouses) GetAll() Warehouses {
	return *w
}

func (w *Warehouses) GetByID(id int) (Warehouse, error) {
	if id < 0 || id > lastID {
		return Warehouse{}, customerrors.ErrorInvalidID
	}

	for _, w := range repository.Warehouse {
		if w.ID == id {
			return Warehouse(w), nil
		}
	}
	return Warehouse{}, customerrors.ErrorInvalidID
}

func (w *Warehouses) Delete(id int) error {
	if id < 0 || id > lastID {
		return customerrors.ErrorInvalidID
	}

	for index, warehouse := range w.Warehouse {
		if warehouse.ID == id {
			w.Warehouse = append(w.Warehouse[:index], w.Warehouse[index+1:]...)
			return nil
		}
	}
	return customerrors.ErrorInvalidID
}

func calculateCode(info ...string) string {
	var builder strings.Builder
	for _, s := range info {
		builder.WriteString(s)
	}
	code := sha256.Sum256([]byte(builder.String()))

	return base64.StdEncoding.EncodeToString(code[:])
}
