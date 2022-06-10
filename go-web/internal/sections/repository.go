package sections

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type Section models.Section
type Sections models.Sections

var repository Sections

var lastID int

func LastID() int {
	return lastID
}

func ValidateID(id int) bool {
	if id < 0 || id > lastID {
		return false
	}
	return true
}

type Repository interface {
	GetAll() Sections
	GetById(int) (models.Section, error)
	Store(models.Section) (models.Section, error)
	Update(int, models.Section) error
	Delete(int) error
}

func NewRepository() Repository {
	return &repository
}

func (s *Sections) GetAll() Sections {
	return *s
}

func (s *Sections) GetById(id int) (models.Section, error) {
	if !ValidateID(id) {
		return models.Section{}, customErrors.ErrorInvalidID
	}

	for _, sec := range s.Section {
		if sec.ID == id {
			return sec, nil
		}
	}
	return models.Section{}, nil
}

func (s *Sections) Store(newSection models.Section) (models.Section, error) {
	newSection.ID = lastID
	s.Section = append(s.Section, models.Section(newSection))
	lastID++

	return newSection, nil
}

func (s *Sections) Update(id int, newSection models.Section) error {
	sectionVec := s.Section
	for i, section := range sectionVec {
		if section.ID == newSection.ID {
			s.Section[i] = newSection
			return nil
		}
	}
	return customErrors.ErrorEmptySection
}

func (s *Sections) Delete(id int) error {
	for i, section := range s.Section {
		if section.ID == id {
			s.Section = append(s.Section[:i], s.Section[i+1:]...)
			return nil
		}
	}

	return customErrors.ErrorInvalidID
}
