package employees

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
)

type Employee models.Employee

var employees []Employee
var lastID int

type Repository interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(id int, card_number_id int, first_name string, last_name string, warehouse_id int) (Employee, error)
	Update(e Employee, id int) (Employee, error)
	Delete(id int) error
	LastID() (int, error)
}

type repository struct {
}

// Create implements Repository
func (r *repository) Create(id int, card_number_id int, first_name string, last_name string, warehouse_id int) (Employee, error) {
	e := Employee{
		ID:           id,
		CardNumberId: card_number_id,
		FirstName:    first_name,
		LastName:     last_name,
		WareHouseId:  warehouse_id,
	}
	employees = append(employees, e)
	lastID = e.ID
	return e, nil
}

// Delete implements Repository
func (r *repository) Delete(id int) error {
	deleted := false

	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			deleted = true
		}
	}

	if !deleted {
		return errors.New("O ID digitado não foi encontrado")
	}

	return nil
}

// GetAll implements Repository
func (r *repository) GetAll() ([]Employee, error) {
	return employees, nil
}

// GetByID implements Repository
func (*repository) GetByID(id int) (Employee, error) {
	for _, e := range employees {
		if e.ID == id {
			return e, nil
		}
	}

	return Employee{}, errors.New("O ID digitado não foi encontrado")
}

// LastID implements Repository
func (r *repository) LastID() (int, error) {
	return lastID, nil
}

// Update implements Repository
func (r *repository) Update(e Employee, id int) (Employee, error) {

	updated := false
	for i := range employees {
		if employees[i].ID == id {
			employees[i] = e
			updated = true
		}
	}

	if !updated {
		return Employee{}, errors.New("O ID digitado não foi encontrado")
	}

	return e, nil
}

func NewRepository() Repository {
	return &repository{}
}
