package sections

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db"
)

type Section models.Section

var storage models.Sections

type repository struct {
	database db.DB
}

func (s *repository) ValidateID(id int) bool {
	err := s.database.Load(&storage)
	if err != nil {
		return false
	}

	if id < 0 || id > storage.LastID {
		return false
	}
	return true
}

type Repository interface {
	GetAll() ([]models.Section, error)
	GetById(int) (models.Section, error)
	Store(Section) (Section, error)
	Update(int, Section) error
	Delete(int) error
	ValidateID(int) bool
	VerifySectionNumber(int) error
}

func NewRepository(database db.DB) Repository {
	return &repository{
		database: database,
	}
}

func (s *repository) GetAll() ([]models.Section, error) {
	err := s.database.Load(&storage)
	if err != nil {
		return []models.Section{}, nil
	}
	return storage.SectionList, nil
}

func (s *repository) GetById(id int) (models.Section, error) {
	err := s.database.Load(&storage)
	if err != nil {
		return models.Section{}, err
	}
	for _, section := range storage.SectionList {
		if section.ID == id {
			return section, nil
		}
	}
	return models.Section{}, nil
}

func (s *repository) VerifySectionNumber(sectionRequestedNumber int) error {
	for _, section := range storage.SectionList {
		if section.SectionNumber == sectionRequestedNumber {
			return customErrors.ErrorConflict
		}
	}
	return nil
}

func (s *repository) Store(newSection Section) (Section, error) {
	err := s.database.Load(&storage)
	if err != nil {
		return Section{}, err
	}
	newSection.ID = storage.LastID
	storage.SectionList = append(storage.SectionList, models.Section(newSection))
	storage.LastID = storage.LastID + 1
	err = s.database.Save(&storage)

	if err != nil {
		return Section{}, err
	}

	return newSection, nil
}

func (s *repository) Update(id int, newSection Section) error {
	for i, section := range storage.SectionList {
		if section.ID == newSection.ID {

			section = models.Section(newSection)
			storage.SectionList[i] = section

			err := s.database.Save(&storage)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return customErrors.ErrorEmptySection
}

func (s *repository) Delete(id int) error {
	err := s.database.Load(&storage)
	if err != nil {
		return err
	}

	for i, section := range storage.SectionList {
		if section.ID == id {
			storage.SectionList = append(storage.SectionList[:i], storage.SectionList[i+1:]...)
			err = s.database.Save(&storage)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return customErrors.ErrorInvalidID
}
