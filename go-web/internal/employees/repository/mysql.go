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
	data := r.data
	query := "DELETE FROM employees WHERE id = ?"
	queryEmployee, err := data.Prepare(query)

	if err != nil {
		return err
	}
	result, err := queryEmployee.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return customerrors.ErrorInvalidID
	}

	return nil
}

// GetAll implements Repository
func (r *repository) GetAll() ([]employees.Employee, error) {
	data := r.data
	query := "SELECT * FROM employees"
	queryEmployee, err := data.Query(query)
	if err != nil {
		log.Println(customerrors.ErrorInvalidDB)
		return []employees.Employee{}, err
	}
	var result []employees.Employee
	for queryEmployee.Next() {
		var employee employees.Employee
		if err := queryEmployee.Scan(&employee.ID, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WareHouseId); err != nil {
			log.Println(err)
		}
		result = append(result, employee)
	}

	return result, nil
}

// GetByID implements Repository
func (r *repository) GetByID(id int) (employees.Employee, error) {
	data := r.data
	query := "SELECT * FROM employees where id = ?"

	queryId, err := data.Query(query, id)
	if err != nil {
		log.Println(customerrors.ErrorInvalidDB)
		return employees.Employee{}, err
	}

	var employee employees.Employee
	for queryId.Next() {
		if err := queryId.Scan(&employee.ID, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WareHouseId); err != nil {
			log.Println(err)
		}
	}

	if employee.ID != id {
		return employee, customerrors.ErrorInvalidID
	}

	return employee, nil
}

// LastID implements Repository
func (r *repository) LastID() (int, error) {
	return lastID, nil
}

// Update implements Repository
func (r *repository) Update(employee employees.Employee, id int) (employees.Employee, error) {
	data := r.data
	query := "UPDATE employees SET id_card_number = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?"
	queryEmployee, err := data.Prepare(query)
	if err != nil {
		return employees.Employee{}, err
	}

	_, err = queryEmployee.Exec(&employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WareHouseId, id)
	if err != nil {
		return employees.Employee{}, err
	}

	return employee, nil
}

func NewRepository(data *sql.DB) employees.Repository {
	return &repository{
		data: data,
	}
}
