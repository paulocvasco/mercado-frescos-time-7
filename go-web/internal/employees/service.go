package employees

import (
	"encoding/json"
	jsonpatch "github.com/evanphx/json-patch/v5"
)

type RequestPatch struct {
	CardNumberId string `json:"card_number_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	WareHouseId  int    `json:"warehouse_id,omitempty"`
}
type Service interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(card_number_id string, first_name string, last_name string, warehouse_id int) (Employee, error)
	Update(e RequestPatch, id int) (Employee, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

// Create implements Service
func (s *service) Create(card_number_id string, first_name string, last_name string, warehouse_id int) (Employee, error) {
	err := s.repository.ValidationCardNumberID(card_number_id)
	if err != nil {
		return Employee{}, err
	}

	lastID, err := s.repository.LastID()

	if err != nil {
		return Employee{}, err
	}

	lastID++

	employee, err := s.repository.Create(lastID, card_number_id, first_name, last_name, warehouse_id)

	if err != nil {
		return Employee{}, err
	}

	return employee, nil

}

func (s *service) GetAll() ([]Employee, error) {
	employees, err := s.repository.GetAll()

	if err != nil {
		return []Employee{}, err
	}

	return employees, nil

}

// Delete implements Service
func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}

// GetByID implements Service
func (s *service) GetByID(id int) (Employee, error) {
	employees, err := s.repository.GetByID(id)

	if err != nil {
		return Employee{}, err
	}

	return employees, nil
}

func (s *service) Update(e RequestPatch, id int) (Employee, error) {
	err := s.repository.ValidationCardNumberID(e.CardNumberId)
	if err != nil {
		return Employee{}, err
	}

	employee, err := s.repository.GetByID(id)

	if err != nil {
		return Employee{}, err
	}

	oldEmployeeJson, err := json.Marshal(employee)

	if err != nil {
		return Employee{}, err
	}

	newEmployee, err := json.Marshal(e)

	if err != nil {
		return Employee{}, err
	}

	updatedEmployee, err := jsonpatch.MergePatch(oldEmployeeJson, newEmployee)

	if err != nil {
		return Employee{}, err
	}

	err = json.Unmarshal(updatedEmployee, &employee)

	if err != nil {
		return Employee{}, err
	}

	employees, err := s.repository.Update(employee, id)

	if err != nil {
		return Employee{}, err
	}

	return employees, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
