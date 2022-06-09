package warehouse

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch"
)

type Service interface {
	GetAll() []models.Warehouse
	GetByID(int) (models.Warehouse, error)
	Create(models.Warehouse) (models.Warehouse, error)
	Update(int, []byte) (models.Warehouse, error)
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

func (s *service) GetAll() ([]models.Warehouse, error) {
	data, err := s.repository.GetAll()
	if err != nil {
		return nil, err
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

	newWarehouse, err = s.repository.Create(newWarehouse)
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
		return models.Warehouse{}, nil
	}

	return warehouse, nil
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
