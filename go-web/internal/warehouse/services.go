package warehouse

import (
	"encoding/json"
	customerrors "mercado-frescos-time-7/go-web/internal/custom_errors"
	"strconv"
)

type Service interface {
	GetAll() (Warehouses, error)
	GetByID(string) (*Warehouse, error)
	Create([]byte) error
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

func (s *service) GetAll() (Warehouses, error) {
	data := s.repository.GetAll()
	return data, nil
}

func (s *service) GetByID(id string) (*Warehouse, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return nil, customerrors.ErrorInvalidIDParameter
	}
	data, err := s.repository.GetByID(index)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *service) Create(data []byte) error {
	var newWarehouse Warehouse
	err := json.Unmarshal(data, &newWarehouse)
	if err != nil {
		return err
	}

	// validate request fields
	if newWarehouse.Address == "" {
		return customerrors.ErrorMissingAddres
	}
	if newWarehouse.Telephone == "" {
		return customerrors.ErrorMissingTelephone
	}
	if newWarehouse.MinimunCapacity == 0 {
		return customerrors.ErrorMissingCapacity
	}
	if newWarehouse.MinimunTemperature == 0 {
		return customerrors.ErrorMissingTemperature
	}

	s.repository.Create(newWarehouse)

	return nil
}

func (s *service) Update(id string, data []byte) error {
	index, err := strconv.Atoi(id)
	if err != nil {
		return customerrors.ErrorInvalidIDParameter
	}

	err = s.repository.Update(index, data)
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
