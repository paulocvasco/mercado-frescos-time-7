package sections

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customErrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/db"
)

type repository struct {
	database db.DB
}

var cache models.SectionMetaData

//go:generate mockery --name=Repository --output=./mock/mockRepository --outpkg=mockRepository
type Repository interface {
	GetAll() (models.Sections, error)
	GetById(int) (models.Section, error)
	Store(models.Section) (models.Section, error)
	Update(int, models.Section) error
	Delete(int) error
	ValidateID(int) bool
	VerifySectionNumber(int) error
}

func NewRepository(database db.DB) Repository {
	return &repository{
		database: database,
	}
}

func (s *repository) GetAll() (models.Sections, error) {

	err := s.database.Load(&cache)
	if err != nil {
		return models.Sections{}, nil
	}
	return cache.Content, nil
}

func (s *repository) GetById(id int) (models.Section, error) {
	err := s.database.Load(&cache)
	if err != nil {
		return models.Section{}, err
	}
	for _, section := range cache.Content.SectionList {
		if section.ID == id {
			return section, nil
		}
	}
	return models.Section{}, nil
}

func (s *repository) VerifySectionNumber(sectionRequestedNumber int) error {
	for _, section := range cache.Content.SectionList {
		if section.SectionNumber == sectionRequestedNumber {
			return customErrors.ErrorConflict
		}
	}
	return nil
}

func (s *repository) Store(newSection models.Section) (models.Section, error) {
	var sections models.SectionMetaData
	err := s.database.Load(&sections)
	if err != nil {
		return models.Section{}, err
	}
	newSection.ID = sections.LastID
	sections.Content.SectionList = append(sections.Content.SectionList, models.Section(newSection))
	sections.LastID = sections.LastID + 1

	err = s.database.Save(&sections)
	if err != nil {
		return models.Section{}, err
	}

	cache = sections
	return newSection, nil
}

func (s *repository) Update(id int, newSection models.Section) error {
	for i, section := range cache.Content.SectionList {
		if section.ID == newSection.ID {

			section = models.Section(newSection)
			cache.Content.SectionList[i] = section

			err := s.database.Save(&cache)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return customErrors.ErrorEmptySection
}

func (s *repository) Delete(id int) error {
	err := s.database.Load(&cache)
	if err != nil {
		return err
	}

	for i, section := range cache.Content.SectionList {
		if section.ID == id {
			cache.Content.SectionList = append(cache.Content.SectionList[:i], cache.Content.SectionList[i+1:]...)
			err = s.database.Save(&cache)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return customErrors.ErrorInvalidID
}

func (s *repository) ValidateID(id int) bool {
	err := s.database.Load(&cache)
	if err != nil {
		return false
	}

	if id < 0 || id > cache.LastID {
		return false
	}
	return true
}
