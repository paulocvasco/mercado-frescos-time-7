package sections

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Service interface {
	GetAll() (Sections, error)
	GetById(string) (*Section, error)
	Store([]byte) error
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

func (s *service) GetAll() (Sections, error) {
	data := s.repository.GetAll()
	return data, nil
}

func (s *service) GetById(id string) (*Section, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}
	data, err := s.repository.GetById(index)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *service) Store(data []byte) error {
	var newSection Section
	err := json.Unmarshal(data, &newSection)
	if err != nil {
		return err
	}

	// validate request fields
	if newSection.Current_capacity < 0 {
		return fmt.Errorf("missing current capacity")
	}
	if newSection.Minimum_capacity < 0 {
		return fmt.Errorf("missing current capacity")
	}
	if newSection.Maximim_capacity < 0 {
		return fmt.Errorf("missing current capacity")
	}
	if newSection.Warehouse_id < 0 {
		return fmt.Errorf("missing current capacity")
	}
	if newSection.Product_type_id < 0 {
		return fmt.Errorf("missing current capacity")
	}

	s.repository.Store(newSection)

	return nil
}

func (s *service) Update(id string, data []byte) error {
	index, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid id")
	}

	var newSection Section
	err = json.Unmarshal(data, &newSection)
	if err != nil {
		return err
	}

	err = s.repository.Update(index, newSection)
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
