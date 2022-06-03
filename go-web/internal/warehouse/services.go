package warehouse

import (
	"encoding/json"
	customerrors "mercado-frescos-time-7/go-web/internal/custom_errors"
	"strconv"
)

type Service interface {
	GetAll() (string, error)
	GetByID(string) (string, error)
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

func (s *service) GetAll() (string, error) {
	data := s.repository.GetAll()
	response, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (s *service) GetByID(id string) (string, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return "", customerrors.ErrorInvalidIDParameter
	}
	data, err := s.repository.GetByID(index)
	if err != nil {
		return "", err
	}
	response, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(response), nil
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

	var updatedWarehouse Warehouse
	err = json.Unmarshal(data, &updatedWarehouse)
	if err != nil {
		return err
	}

	err = s.repository.Update(index, updatedWarehouse)
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
