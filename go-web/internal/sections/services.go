package sections

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

//go:generate mockery --name=Service --output=./mock/mockService --outpkg=mockService
type Service interface {
	GetAll() (models.Sections, error)
	GetById(string) (models.Section, error)
	Store([]byte) (models.Section, error)
	Update(string, []byte) (models.Section, error)
	Delete(string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() (models.Sections, error) {
	data, err := s.repository.GetAll()
	if err != nil {
		return models.Sections{}, err
	}
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

func (s *service) Store(section []byte) (models.Section, error) {
	var newSection models.Section
	err := json.Unmarshal(section, &newSection)

	if err != nil {
		return models.Section{}, nil
	}

	err = s.repository.VerifySectionNumber(newSection.SectionNumber)

	if err != nil {
		return models.Section{}, customErrors.ErrorConflict
	}

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
	newSection, err = s.repository.Store(newSection)
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
	if !s.repository.ValidateID(index) {
		return models.Section{}, customErrors.ErrorInvalidID
	}

	sectionFound, err := s.repository.GetById(index)
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
	if !s.repository.ValidateID(index) {
		return customErrors.ErrorInvalidID
	}

	_, err = s.repository.GetById(index)

	if err != nil {
		return customErrors.ErrorInvalidID
	}

	err = s.repository.Delete(index)

	if err != nil {
		return err
	}

	return nil
}
