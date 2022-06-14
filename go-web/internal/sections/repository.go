package sections

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type Section models.Section
type repository struct{}

var db []Section
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
	GetAll() []Section
	GetById(int) (Section, error)
	Store(Section) (Section, error)
	Update(int, Section) error
	Delete(int) error
	VerifySectionNumber(int) error
}

func NewRepository() Repository {
	return &repository{}
}

func (s *repository) GetAll() []Section {
	return db
}

func (s *repository) GetById(id int) (Section, error) {
	if !ValidateID(id) {
		return Section{}, customErrors.ErrorInvalidID
	}

	for _, section := range db {
		if section.ID == id {
			return section, nil
		}
	}
	return Section{}, nil
}

func (s *repository) VerifySectionNumber(sectionRequestedNumber int) error {
	for _, section := range db {
		if section.SectionNumber == sectionRequestedNumber {
			return customErrors.ErrorConflict
		}
	}
	return nil
}

func (s *repository) Store(newSection Section) (Section, error) {
	newSection.ID = lastID
	db = append(db, newSection)
	lastID++

	return newSection, nil
}

func (s *repository) Update(id int, newSection Section) error {
	if !ValidateID(id) {
		return customErrors.ErrorInvalidID
	}

	for i, section := range db {
		if section.ID == newSection.ID {
			db[i] = newSection
			return nil
		}
	}
	return customErrors.ErrorEmptySection
}

func (s *repository) Delete(id int) error {
	for i, section := range db {
		if section.ID == id {
			db = append(db[:i], db[i+1:]...)
			return nil
		}
	}

	return customErrors.ErrorInvalidID
}
