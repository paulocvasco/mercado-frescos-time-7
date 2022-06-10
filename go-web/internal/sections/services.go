package sections

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	GetAll() (Sections, error)
	GetById(string) (models.Section, error)
	Store(models.Section) (models.Section, error)
	Update(string, []byte) (models.Section, error)
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

func (s *service) GetById(id string) (models.Section, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return models.Section{}, customErrors.ErrorInvalidID
	}
	data, err := s.repository.GetById(index)

	if err != nil || (data == models.Section{}) {
		return models.Section{}, customErrors.ErrorInvalidID
	}

	return data, nil
}

func (s *service) Store(section models.Section) (models.Section, error) {

	if section.SectionNumber < 0 {
		return models.Section{}, customErrors.ErrorSectionNumber
	}
	if section.CurrentCapacity < 0 {
		return models.Section{}, customErrors.ErrorCurrentCapacity
	}
	if section.MinimumCapacity < 0 {
		return models.Section{}, customErrors.ErrorMinimumCapacity
	}
	if section.MaximumCapacity < 0 {
		return models.Section{}, customErrors.ErrorMaximumCapacity
	}
	if section.WarehouseId < 0 {
		return models.Section{}, customErrors.ErrorWarehouseID
	}
	if section.ProductTypeId < 0 {
		return models.Section{}, customErrors.ErrorProductTypeID
	}

	newSection, err := s.repository.Store(section)
	if err != nil {
		return models.Section{}, customErrors.ErrorStoreFailed
	}
	return newSection, nil
}

func (s *service) Update(id string, data []byte) (models.Section, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return models.Section{}, customErrors.ErrorInvalidID
	}

	sectionFound, err := repository.GetById(index)
	if err != nil {
		return models.Section{}, customErrors.ErrorInvalidID
	}
	sectionFoundJSON, _ := json.Marshal(sectionFound)
	patch, err := jsonpatch.MergePatch(sectionFoundJSON, data)

	if err != nil {
		return models.Section{}, err
	}
	var newSection models.Section
	json.Unmarshal(patch, &newSection)

	if newSection.SectionNumber < 0 {
		return models.Section{}, customErrors.ErrorSectionNumber
	}
	if newSection.CurrentCapacity < 0 {
		return models.Section{}, customErrors.ErrorCurrentCapacity
	}
	if newSection.MinimumCapacity < 0 {
		return models.Section{}, customErrors.ErrorMinimumCapacity
	}
	if newSection.MaximumCapacity < 0 {
		return models.Section{}, customErrors.ErrorMaximumCapacity
	}
	if newSection.WarehouseId < 0 {
		return models.Section{}, customErrors.ErrorWarehouseID
	}
	if newSection.ProductTypeId < 0 {
		return models.Section{}, customErrors.ErrorProductTypeID
	}

	err = s.repository.Update(index, newSection)
	if err != nil {
		return models.Section{}, err
	}

	return newSection, nil
}

func (s *service) Delete(id string) error {
	index, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = repository.GetById(index)

	if err != nil {
		return customErrors.ErrorInvalidID
	}

	err = s.repository.Delete(index)

	if err != nil {
		return err
	}

	return nil
}
