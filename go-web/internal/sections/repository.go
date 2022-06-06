package sections

import (
	"fmt"
	"mercado-frescos-time-7/go-web/internal/models"
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
	GetById(int) (Section, error)
	Store(models.Section) error
	Update(int, Section) error
	Delete(int) error
}

func NewRepository() Repository {
	return &repository
}

func (s *Sections) GetAll() Sections {
	return *s
}

func (s *Sections) GetById(id int) (Section, error) {
	if id < 0 || id > lastID {
		return Section{}, fmt.Errorf("invalid id")
	}
	for _, s := range repository.Section {
		if s.ID == id {
			return Section(s), nil
		}
	}
	return Section{}, nil
}

func (s *Sections) Store(newSection models.Section) error {
	newSection.ID = lastID
	s.Section = append(s.Section, models.Section(newSection))
	lastID++
	return nil
}

func (s *Sections) Update(id int, newSection Section) error {
	st, err := s.GetById(id)

	if err != nil {
		return err
	}

	if (st == Section{}) {
		return fmt.Errorf("empty section")
	}

	st.Current_temperature = newSection.Current_temperature
	st.Minimum_temperature = newSection.Minimum_temperature
	st.Section_number = newSection.Section_number
	st.Current_capacity = newSection.Current_capacity
	st.Minimum_capacity = newSection.Minimum_capacity
	st.Maximum_capacity = newSection.Maximum_capacity
	st.Warehouse_id = newSection.Warehouse_id
	st.Product_type_id = newSection.Product_type_id

	s.Section[id] = models.Section(st)

	return nil
}

func (s *Sections) Delete(id int) error {
	for i, section := range s.Section {
		if section.ID == id {
			s.Section = append(s.Section[:i], s.Section[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("invalid id")
}
