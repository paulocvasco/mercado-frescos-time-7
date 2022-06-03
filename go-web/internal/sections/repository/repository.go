package sections

import (
	"fmt"
)

type Section struct {
	ID                  int `json:"id"`
	Section_number      int `json:"section_number"`
	Current_temperature int `json:"current_temperature"`
	Minimum_temperature int `json:"minimum_temperature"`
	Current_capacity    int `json:"current_capacity"`
	Minimum_capacity    int `json:"minimum_capacity"`
	Maximim_capacity    int `json:"maximim_capacity"`
	Warehouse_id        int `json:"warehouse_id"`
	Product_type_id     int `json:"product_type_id"`
}

type Sections struct {
	Section []Section `json:"data"`
}

type Repository interface {
	GetAll() Sections
	GetById(int) (Section, error)
	Store(Section)
	Update(int, Section) error
	Delete(int) error
}

var repository Sections

var lastID int

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
			return s, nil
		}
	}
	return Section{}, nil
}

func (s *Sections) Store(newSection Section) {
	newSection.ID = lastID
	s.Section = append(s.Section, newSection)
	lastID++
}

func (s *Sections) Update(id int, newSection Section) error {
	st, err := s.GetById(id)

	if err != nil {
		return err
	}

	if (st != Section{}) {
		return fmt.Errorf("empty section")
	}

	st.Current_temperature = newSection.Current_temperature
	st.Minimum_temperature = newSection.Minimum_temperature

	if newSection.Section_number < 0 {
		st.Section_number = newSection.Section_number
	}
	if newSection.Current_capacity < 0 {
		st.Current_capacity = newSection.Current_capacity
	}
	if newSection.Minimum_capacity < 0 {
		st.Minimum_capacity = newSection.Minimum_capacity
	}
	if newSection.Maximim_capacity < 0 {
		st.Maximim_capacity = newSection.Maximim_capacity
	}
	if newSection.Warehouse_id < 0 {
		st.Warehouse_id = newSection.Warehouse_id
	}
	if newSection.Product_type_id < 0 {
		st.Product_type_id = newSection.Product_type_id
	}

	s.Section[id] = st

	return nil
}

func (s *Sections) Delete(id int) error {
	if id < 0 || id > lastID {
		return fmt.Errorf("invalid id")
	}
	for i, section := range s.Section {
		if section.ID == id {
			s.Section = append(s.Section[:i], s.Section[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("invalid id")
}
