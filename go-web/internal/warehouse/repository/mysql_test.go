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
	type mockResponse struct {
		id            int
		address       string
		telephone     string
		warehouseCode string
		minCap        int
		minTmp        int
		localID       int
	}
	testCases := []struct {
		testName         string
		queryResponse    []mockResponse
		queryError       error
		expectedResponse models.Warehouses
		expectedError    error
	}{
		{
			"FailedToQueryDB",
			nil,
			customerrors.ErrorInvalidDB,
			models.Warehouses{},
			customerrors.ErrorInvalidDB,
		},
		{
			"EmptyList",
			nil,
			nil,
			models.Warehouses{},
			nil,
		},
		{
			"FilledList",
			[]mockResponse{{1, "foo", "1111", "bar", 10, 20, 1}, {2, "dummy", "2222", "foobar", 12, 14, 2}},
			nil,
			models.Warehouses{Warehouses: []models.Warehouse{{ID: 1, Address: "foo", Telephone: "1111", WarehouseCode: "bar", MinimunCapacity: 10, MinimunTemperature: 20, LocalityID: 1},
				{ID: 2, Address: "dummy", Telephone: "2222", WarehouseCode: "foobar", MinimunCapacity: 12, MinimunTemperature: 14, LocalityID: 2}}},
			nil,
		},
	}

	query := "SELECT * FROM warehouse"
	for _, v := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		rows := mock.NewRows([]string{"id", "address", "telephone", "warehouse_code", "minimun_capacity", "minimun_temperatue", "locality_id"})
		for _, r := range v.queryResponse {
			rows.AddRow(r.id, r.address, r.telephone, r.warehouseCode, r.minCap, r.minTmp, r.localID)
		}

		exec := mock.ExpectQuery(query)
		exec.WillReturnRows(rows)
		exec.WillReturnError(v.queryError)

		repo := mysqlDB{db: db}
		all, err := repo.GetAll()
		if v.expectedError != err {
			t.Errorf("GetAll test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
			continue
		}

		for i, w := range all.Warehouses {
			if v.expectedResponse.Warehouses[i] != w {
				t.Errorf("GetAll test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedResponse, all)
			}
		}
	}
}

func TestGetByID(t *testing.T) {
	type mockResponse struct {
		id            int
		address       string
		telephone     string
		warehouseCode string
		minCap        int
		minTmp        int
		localID       int
	}
	testCases := []struct {
		testName         string
		id               int
		queryResponse    []mockResponse
		queryError       error
		expectedResponse models.Warehouse
		expectedError    error
	}{
		{
			"GetID", 1,
			[]mockResponse{{1, "foo", "1111", "bar", 10, 20, 1}, {2, "dummy", "2222", "foobar", 12, 14, 2}},
			nil,
			models.Warehouse{ID: 1, Address: "foo", Telephone: "1111", WarehouseCode: "bar", MinimunCapacity: 10, MinimunTemperature: 20, LocalityID: 1},
			nil,
		},
		{
			"ItemNotFound", 1,
			nil,
			customerrors.ErrorItemNotFound,
			models.Warehouse{},
			customerrors.ErrorItemNotFound,
		},
	}

	query := "SELECT * FROM warehouse WHERE id = ?"
	for _, v := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		rows := mock.NewRows([]string{"id", "address", "telephone", "warehouse_code", "minimun_capacity", "minimun_temperatue", "locality_id"})
		for _, r := range v.queryResponse {
			rows.AddRow(r.id, r.address, r.telephone, r.warehouseCode, r.minCap, r.minTmp, r.localID)
		}

		exec := mock.ExpectQuery(query).WithArgs(v.id)
		exec.WillReturnRows(rows)
		exec.WillReturnError(v.queryError)

		repo := mysqlDB{db: db}
		w, err := repo.GetByID(v.id)
		if v.expectedError != err {
			t.Errorf("GetById test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}

		if v.expectedResponse != w {
			t.Errorf("GetById test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedResponse, w)
		}
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		testName      string
		id            int
		prepError     error
		mockResponse  driver.Result
		mockError     error
		expectedError error
	}{
		{
			"DeleteOk", 1,
			nil, sqlmock.NewResult(1, 1), nil,
			nil,
		},
		{
			"PrepError", 1,
			customerrors.ErrorInvalidDB, sqlmock.NewResult(1, 1), nil,
			customerrors.ErrorInvalidDB,
		},
		{
			"ItemNotFound", 1,
			nil, sqlmock.NewResult(1, 0), nil,
			customerrors.ErrorItemNotFound,
		},
		{
			"ExecError", 1,
			nil, sqlmock.NewResult(0, 0), customerrors.ErrorInvalidDB,
			customerrors.ErrorInvalidDB,
		},
	}

	query := "DELETE FROM warehouse WHERE id = ?"
	for _, v := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		stmt := mock.ExpectPrepare(query)
		stmt.WillReturnError(v.prepError)
		exec := stmt.ExpectExec()
		exec.WillReturnResult(v.mockResponse)
		exec.WillReturnError(v.mockError)

		repo := mysqlDB{db: db}
		err = repo.Delete(v.id)
		if v.expectedError != err {
			t.Errorf("Delete test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		testName      string
		id            int
		updatedModel  models.Warehouse
		prepError     error
		mockResponse  driver.Result
		mockError     error
		expectedError error
	}{
		{
			"UpdateOk", 1, models.Warehouse{},
			nil, sqlmock.NewResult(1, 1), nil,
			nil,
		},
		{
			"ItemNotFound", 1, models.Warehouse{},
			nil, sqlmock.NewResult(1, 0), nil,
			customerrors.ErrorItemNotFound,
		},
		{
			"ErrorPrep", 1, models.Warehouse{},
			customerrors.ErrorInvalidDB, sqlmock.NewResult(1, 0), nil,
			customerrors.ErrorInvalidDB,
		},
		{
			"ErrorExec", 1, models.Warehouse{},
			nil, sqlmock.NewResult(1, 0), customerrors.ErrorConflict,
			customerrors.ErrorConflict,
		},
	}

	query := "UPDATE warehouse SET address = ?, telephone = ?, warehouse_code = ?, minimum_capacity = ?, minimum_temperature = ?, locality_id = ? WHERE id = ?"
	for _, v := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		stmt := mock.ExpectPrepare(query)
		stmt.WillReturnError(v.prepError)
		exec := stmt.ExpectExec()
		exec.WillReturnResult(v.mockResponse)
		exec.WillReturnError(v.mockError)

		repo := mysqlDB{db: db}
		err = repo.Update(v.id, v.updatedModel)
		if v.expectedError != err {
			t.Errorf("Update test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}
	}
}
