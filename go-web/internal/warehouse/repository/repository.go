package repository

import (
	"crypto/sha256"
	"encoding/base64"
	customerrors "mercado-frescos-time-7/go-web/internal/custom_errors"
	"strconv"
	"strings"
)

type Warehouses struct {
	Warehouse []Warehouse `json:"data"`
}

type Warehouse struct {
	ID                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

func (w *Warehouses) Create(new Warehouse) {
	new.ID = len(w.Warehouse) + 1
	new.WarehouseCode = calculateCode(new.Address, new.Telephone, strconv.Itoa(new.MinimunCapacity), strconv.Itoa(new.MinimunTemperature))

	w.Warehouse = append(w.Warehouse, new)
}

func (w *Warehouses) Update(id int, newValues Warehouse) error {
	if id < 0 || id > len(w.Warehouse) {
		return customerrors.ErrorInvalidID
	}

	warehouse := w.Warehouse[id]
	if newValues.Address != "" {
		warehouse.Address = newValues.Address
	}
	if newValues.Telephone != "" {
		warehouse.Telephone = newValues.Telephone
	}
	if newValues.MinimunCapacity != 0 {
		warehouse.MinimunCapacity = newValues.MinimunCapacity
	}
	if newValues.MinimunTemperature != 0 {
		warehouse.MinimunCapacity = newValues.MinimunTemperature
	}
	warehouse.WarehouseCode = calculateCode(warehouse.Address, warehouse.Telephone, strconv.Itoa(warehouse.MinimunCapacity), strconv.Itoa(warehouse.MinimunTemperature))

	w.Warehouse[id] = warehouse
	return nil
}

func (w *Warehouses) GetAll() Warehouses {
	return *w
}

func (w *Warehouses) GetByID(id int) (Warehouse, error) {
	if id < 0 || id >= len(w.Warehouse) {
		return Warehouse{}, customerrors.ErrorInvalidID
	}

	return (w.Warehouse)[id], nil
}

func calculateCode(info ...string) string {
	var builder strings.Builder
	for _, s := range info {
		builder.WriteString(s)
	}
	code := sha256.Sum256([]byte(builder.String()))

	return base64.StdEncoding.EncodeToString(code[:])
}
