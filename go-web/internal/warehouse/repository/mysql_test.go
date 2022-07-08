package repository

import (
	"database/sql/driver"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type mockedResult struct {
	lastID    int64
	rows      int64
	errorID   error
	errorRows error
}

func (m *mockedResult) LastInsertId() (int64, error) {
	return m.lastID, m.errorID
}

func (m *mockedResult) RowsAffected() (int64, error) {
	return m.rows, m.errorRows
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		testName      string
		storeModel    models.Warehouse
		mockResponse  driver.Result
		mockError     error
		expectedModel models.Warehouse
		expectedError error
	}{
		{"Success",
			models.Warehouse{Address: "foo", Telephone: "0000", WarehouseCode: "bar", MinimunCapacity: 10, MinimunTemperature: 20, LocalityID: 1},
			sqlmock.NewResult(1, 1),
			nil,
			models.Warehouse{ID: 1, Address: "foo", Telephone: "0000", WarehouseCode: "bar", MinimunCapacity: 10, MinimunTemperature: 20, LocalityID: 1},
			nil,
		},
		{"ExecError",
			models.Warehouse{Address: "foo", Telephone: "0000", WarehouseCode: "bar", MinimunCapacity: 10, MinimunTemperature: 20, LocalityID: 1},
			sqlmock.NewResult(1, 1),
			customerrors.ErrorInvalidDB,
			models.Warehouse{},
			customerrors.ErrorInvalidDB,
		},
		{"IdError",
			models.Warehouse{Address: "foo", Telephone: "0000", WarehouseCode: "bar", MinimunCapacity: 10, MinimunTemperature: 20, LocalityID: 1},
			&mockedResult{errorID: customerrors.ErrorInvalidID},
			nil,
			models.Warehouse{},
			customerrors.ErrorInvalidID,
		},
	}

	query := "INSERT INTO warehouse(address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) VALUES (?, ?, ?, ?, ?, ?)"
	for _, v := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		exec := mock.ExpectExec(query)
		exec.WillReturnResult(v.mockResponse)
		exec.WillReturnError(v.mockError)

		repo := mysqlDB{db: db}
		new, err := repo.Create(v.storeModel)
		if v.expectedError != err {
			t.Errorf("Create test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}

		if v.expectedModel != new {
			t.Errorf("Create test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedModel, new)
		}
	}
}

func TestGetAll(t *testing.T) {
	testCases := []struct {

	}
}
