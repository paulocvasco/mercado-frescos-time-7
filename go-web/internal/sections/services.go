package sections

import (
	"encoding/json"
	"fmt"
	customerrors "mercado-frescos-time-7/go-web/internal/custom_errors"
	"mercado-frescos-time-7/go-web/internal/models"
	"strconv"
)

type Service interface {
	GetAll() (Sections, error)
	GetById(string) (*Section, error)
	Store(models.Section) error
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

func (s *service) Store(section models.Section) error {

	if section.SectionNumber < 0 {
		return customerrors.ErrorSectionNumber
	}
	if section.CurrentCapacity < 0 {
		return customerrors.ErrorCurrentCapacity
	}
	if section.MinimumCapacity < 0 {
		return customerrors.ErrorMinimumCapacity
	}
	if section.MaximumCapacity < 0 {
		return customerrors.ErrorMaximumCapacity
	}
	if section.WarehouseId < 0 {
		return customerrors.ErrorWarehouseID
	}
	if section.ProductTypeId < 0 {
		return customerrors.ErrorProductTypeID
	}

	err := s.repository.Store(section)
	if err != nil {
		return fmt.Errorf("")
	}
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

	if newSection.SectionNumber < 0 {
		return customerrors.ErrorSectionNumber
	}
	if newSection.CurrentCapacity < 0 {
		return customerrors.ErrorCurrentCapacity
	}
	if newSection.MinimumCapacity < 0 {
		return customerrors.ErrorMinimumCapacity
	}
	if newSection.MaximumCapacity < 0 {
		return customerrors.ErrorMaximumCapacity
	}
	if newSection.WarehouseId < 0 {
		return customerrors.ErrorWarehouseID
	}
	if newSection.ProductTypeId < 0 {
		return customerrors.ErrorProductTypeID
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
	if !ValidateID(index) {
		return err
	}

	err = s.repository.Delete(index)

	if err != nil {
		return err
	}

	return nil
}
