package repository

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/employees"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type repository struct {
	data *sql.DB
}

var lastID int

func (r *repository) ValidationCardNumberID(card_number_id string) error {
	return nil
}

func (r *repository) Create(id int, card_number_id string, first_name string, last_name string, warehouse_id int) (employees.Employee, error) {
	data := r.data
	query := "INSERT INTO employees (`id_card_number`, `first_name`, `last_name`, `warehouse_id`) " +
		"VALUES (?, ?, ?, ?)"

	employeeQuery, err := data.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer employeeQuery.Close()

	var result sql.Result

	result, err = employeeQuery.Exec(card_number_id, first_name, last_name, warehouse_id)
	if err != nil {
		return employees.Employee{}, err
	}

	usedID, _ := result.LastInsertId()
	employeeInsert := employees.Employee{
		ID:           int(usedID),
		CardNumberId: card_number_id,
		FirstName:    first_name,
		LastName:     last_name,
		WareHouseId:  warehouse_id,
	}

	return employeeInsert, nil
}

// Delete implements Repository
func (r *repository) Delete(id int) error {
	return nil
}

// GetAll implements Repository
func (r *repository) GetAll() ([]employees.Employee, error) {
	return []employees.Employee{}, nil
}

// GetByID implements Repository
func (*repository) GetByID(id int) (employees.Employee, error) {

	return employees.Employee{}, customerrors.ErrorInvalidID
}

// LastID implements Repository
func (r *repository) LastID() (int, error) {
	return lastID, nil
}

// Update implements Repository
func (r *repository) Update(e employees.Employee, id int) (employees.Employee, error) {
	return e, nil
}

func NewRepository(data *sql.DB) employees.Repository {
	return &repository{
		data: data,
	}
}
