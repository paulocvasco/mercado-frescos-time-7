package warehouse

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"

	jsonpatch "github.com/evanphx/json-patch"
)

//go:generate mockery --name=Service --output=./mock/mockService --outpkg=mockService
type Service interface {
	GetAll() (models.Warehouses, error)
	GetByID(int) (models.Warehouse, error)
	Create(models.Warehouse) (models.Warehouse, error)
	Update(int, []byte) (models.Warehouse, error)
	Delete(int) error
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

func (s *service) GetAll() (models.Warehouses, error) {
	data, err := s.repository.GetAll()
	if err != nil {
		return models.Warehouses{}, err
	}
	return data, nil
}

func (s *service) GetByID(id int) (models.Warehouse, error) {
	data, err := s.repository.GetByID(id)
	if err != nil {
		return models.Warehouse{}, err
	}

	return data, nil
}

func (s *service) Create(newWarehouse models.Warehouse) (models.Warehouse, error) {
	// validate request fields
	if newWarehouse.Address == "" {
		return models.Warehouse{}, customerrors.ErrorMissingAddres
	}
	if newWarehouse.Telephone == "" {
		return models.Warehouse{}, customerrors.ErrorMissingTelephone
	}
	if newWarehouse.MinimunCapacity < 0 {
		return models.Warehouse{}, customerrors.ErrorMissingCapacity
	}
	if newWarehouse.MinimunTemperature == 0 {
		return models.Warehouse{}, customerrors.ErrorMissingTemperature
	}

	newWarehouse, err := s.repository.Create(newWarehouse)
	if err != nil {
		return models.Warehouse{}, err
	}

	return newWarehouse, nil
}

func (s *service) Update(id int, data []byte) (models.Warehouse, error) {
	warehouse, err := s.repository.GetByID(id)
	if err != nil {
		return models.Warehouse{}, err
	}

	warehouseBytes, err := json.Marshal(warehouse)
	if err != nil {
		return models.Warehouse{}, err
	}
	patchedWarehouse, err := jsonpatch.MergePatch(warehouseBytes, data)
	if err != nil {
		return models.Warehouse{}, err
	}
	err = json.Unmarshal(patchedWarehouse, &warehouse)
	if err != nil {
		return models.Warehouse{}, err
	}

	err = s.repository.Update(id, warehouse)
	if err != nil {
		return models.Warehouse{}, err
	}

	return warehouse, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
