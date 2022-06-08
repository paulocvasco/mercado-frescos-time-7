package warehouse

import (
	"encoding/json"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch"
)

type Service interface {
	GetAll() []Warehouse
	GetByID(string) (Warehouse, error)
	Create([]byte) (Warehouse, error)
	Update(string, []byte) error
	Delete(string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	newService := &service{
		repository: r,
	}
	return newService
}

func (s *service) GetAll() []Warehouse {
	data := s.repository.GetAll()
	return data
}

func (s *service) GetByID(id string) (Warehouse, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return Warehouse{}, customerrors.ErrorInvalidIDParameter
	}
	data, err := s.repository.GetByID(index)
	if err != nil {
		return Warehouse{}, err
	}

	return data, nil
}

func (s *service) Create(data []byte) (Warehouse, error) {
	var newWarehouse Warehouse
	err := json.Unmarshal(data, &newWarehouse)
	if err != nil {
		return Warehouse{}, nil
	}

	// validate request fields
	if newWarehouse.Address == "" {
		return Warehouse{}, customerrors.ErrorMissingAddres
	}
	if newWarehouse.Telephone == "" {
		return Warehouse{}, customerrors.ErrorMissingTelephone
	}
	if newWarehouse.MinimunCapacity == 0 {
		return Warehouse{}, customerrors.ErrorMissingCapacity
	}
	if newWarehouse.MinimunTemperature == 0 {
		return Warehouse{}, customerrors.ErrorMissingTemperature
	}

	newWarehouse = s.repository.Create(newWarehouse)

	return newWarehouse, nil
}

func (s *service) Update(id string, data []byte) error {
	index, err := strconv.Atoi(id)
	if err != nil {
		return customerrors.ErrorInvalidIDParameter
	}

	warehouse, err := s.repository.GetByID(index)
	if err != nil {
		return err
	}

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

	err = s.repository.Update(index, warehouse)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(id string) error {
	index, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.repository.Delete(index)
	if err != nil {
		return err
	}

	return nil
}
