package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"mercado-frescos-time-7/go-web/internal/employees"
	"mercado-frescos-time-7/go-web/internal/employees/repository"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"
)

func TestRepository_Create(t *testing.T) {
	e := employees.Employee{ID: 1, CardNumberId: "1234", FirstName: "Vitoria", LastName: "Souza", WareHouseId: 1}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("INSERT INTO employees (`id_card_number`, `first_name`, `last_name`, `warehouse_id`) VALUES (?, ?, ?, ?)").
		ExpectExec().WithArgs(e.CardNumberId, e.FirstName, e.LastName, e.WareHouseId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := repository.NewRepository(db)
	result, err := repo.Create(e.ID, e.CardNumberId, e.FirstName, e.LastName, e.WareHouseId)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestRepository_Create_Conflict_Exec(t *testing.T) {
	e := employees.Employee{ID: 1, CardNumberId: "1234", FirstName: "Vitoria", LastName: "Souza", WareHouseId: 1}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("INSERT INTO employees (`id_card_number`, `first_name`, `last_name`, `warehouse_id`) VALUES (?, ?, ?, ?)").
		ExpectExec().WithArgs(e.CardNumberId, e.FirstName, e.LastName, e.WareHouseId).
		WithArgs(e.CardNumberId, e.FirstName, e.LastName, e.WareHouseId).
		WillReturnError(customErrors.ErrorInvalidDB)
	repo := repository.NewRepository(db)
	result, err := repo.Create(e.ID, e.CardNumberId, e.FirstName, e.LastName, e.WareHouseId)
	assert.Equal(t, err, customErrors.ErrorInvalidDB)
	assert.NotNil(t, result)
}

func TestRepository(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	repo := repository.NewRepository(db)

	repo.LastID()

}

func TestRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	mock.ExpectQuery("SELECT * FROM employees").
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "warehouse_id"}).
			AddRow(1, "2323", "Vitoria", "Souza", 1))
	repo := repository.NewRepository(db)
	result, err := repo.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	mock.ExpectQuery("SELECT * FROM employees where id = ?").WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "warehouse_id"}).
			AddRow(1, "2323", "Vitoria", "Souza", 1))
	repo := repository.NewRepository(db)
	result, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestRepository_GetAll_Error(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT * FROM employees").
		WillReturnRows(sqlmock.NewRows([]string{"nothing"}).
			AddRow(1))
	repo := repository.NewRepository(db)
	_, err2 := repo.GetAll()
	assert.Equal(t, err2, err2)
}

func TestRepository_GetById_Error(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT * FROM employees where id = ?").WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"nothing"}).
			AddRow(1))
	repo := repository.NewRepository(db)
	_, err2 := repo.GetByID(1)
	assert.Equal(t, err2, err2)
}

func TestRepository_GetById_Error_Query(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	db.Query("SELECT * FROM employees where id = ?", 1)

	repo := repository.NewRepository(db)
	_, err2 := repo.GetByID(1)
	assert.Equal(t, err2, err2)
}

func TestRepository_GetAll_Error_Query(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	db.Query("SELECT * FROM employees")

	repo := repository.NewRepository(db)
	_, err2 := repo.GetAll()
	assert.Equal(t, err2, err2)
}

func TestRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("DELETE FROM	employees where id = ?").
		ExpectExec().WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := repository.NewRepository(db)
	result := repo.Delete(1)

	assert.NotNil(t, result)
}

func TestRepository_Delete_Error_Prepare(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("DELETE FROM	employees where id = ?").WillReturnError(err)

	repo := repository.NewRepository(db)
	result := repo.Delete(1)

	assert.Error(t, result)
}

func TestRepository_Update(t *testing.T) {
	e := employees.Employee{CardNumberId: "1234", FirstName: "Vitoria", LastName: "Souza", WareHouseId: 1}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("UPDATE employees SET id_card_number = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?").
		ExpectExec().WithArgs(e.CardNumberId, e.FirstName, e.LastName, e.WareHouseId, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := repository.NewRepository(db)
	result, err := repo.Update(e, 1)

	assert.Nil(t, err)
	assert.Equal(t, result, e)
}

func TestRepository_Update_Error(t *testing.T) {
	e := employees.Employee{CardNumberId: "1234", FirstName: "Vitoria", LastName: "Souza", WareHouseId: 1}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectPrepare("UPDATE employees SET id_card_number = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?").
		ExpectExec().WithArgs(e.CardNumberId, e.FirstName, e.LastName, e.WareHouseId, 1).
		WillReturnError(err)
	repo := repository.NewRepository(db)
	result, err := repo.Update(e, 1)

	assert.Error(t, err)
	assert.Equal(t, result, employees.Employee{})
}
