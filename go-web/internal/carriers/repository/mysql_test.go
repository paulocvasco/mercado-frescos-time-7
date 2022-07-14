package repository

import (
	"database/sql/driver"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
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
		storeModel    models.Carrier
		prepError     error
		mockResponse  driver.Result
		mockError     error
		expectedModel models.Carrier
		expectedError error
	}{
		{
			"Success", models.Carrier{}, nil,
			sqlmock.NewResult(1, 1), nil,
			models.Carrier{ID: 1}, nil,
		},
		{
			"PrepError", models.Carrier{}, customerrors.ErrorInvalidDB,
			nil, nil,
			models.Carrier{}, customerrors.ErrorInvalidDB,
		},
		{
			"ExecError", models.Carrier{}, nil,
			nil, customerrors.ErrorInvalidID,
			models.Carrier{}, customerrors.ErrorInvalidID,
		},
		{
			"RowsError", models.Carrier{}, nil,
			&mockedResult{errorRows: customerrors.ErrorInvalidDB}, nil,
			models.Carrier{}, customerrors.ErrorInvalidDB,
		},
		{
			"ErrorStore", models.Carrier{}, nil,
			sqlmock.NewResult(1, 0), nil,
			models.Carrier{}, customerrors.ErrorStoreFailed,
		},
		{
			"LastIdError", models.Carrier{}, nil,
			&mockedResult{rows: 1, errorID: customerrors.ErrorInvalidID}, nil,
			models.Carrier{}, customerrors.ErrorInvalidID,
		},
	}
	query := "INSERT INTO carriers(cid, company_name, address, locality_id) VALUES (?, ?, ?, ?)"
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

		repo := NewRepository(db)
		new, err := repo.Create(v.storeModel)
		if v.expectedError != err {
			t.Errorf("Create test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}

		if v.expectedModel != new {
			t.Errorf("Create test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedModel, new)
		}
	}
}

func TestGet(t *testing.T) {
	type mockResponse struct {
		localityId   int
		localityName string
		carrierCount int
	}
	testCases := []struct {
		testName      string
		id            int
		mockResponse  []mockResponse
		mockError     error
		expectedModel models.CarriersReport
		expectedError error
	}{
		{
			"IdValue0", 0,
			[]mockResponse{{localityId: 1, localityName: "foo", carrierCount: 10}}, nil,
			models.CarriersReport{Data: []models.CarrierInfo{{LocalityID: 1, LocalityName: "foo", CarriersCount: 10}}}, nil,
		},
		{
			"IdValue1", 1,
			[]mockResponse{{localityId: 1, localityName: "foo", carrierCount: 10}}, nil,
			models.CarriersReport{Data: []models.CarrierInfo{{LocalityID: 1, LocalityName: "foo", CarriersCount: 10}}}, nil,
		},
		{
			"QueryError", 0,
			[]mockResponse{}, customerrors.ErrorInvalidDB,
			models.CarriersReport{Data: []models.CarrierInfo{{LocalityID: 1, LocalityName: "foo", CarriersCount: 10}}}, customerrors.ErrorInvalidDB,
		},
		{
			"ItemNotFound", 4,
			[]mockResponse{}, nil,
			models.CarriersReport{}, customerrors.ErrorItemNotFound,
		},
	}
	query := "^SELECT (.*)"
	for _, v := range testCases {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		rows := mock.NewRows([]string{"locality_id", "locality_name", "carriers_count"})
		for _, r := range v.mockResponse {
			rows.AddRow(r.localityId, r.localityName, r.carrierCount)
		}

		exec := mock.ExpectQuery(query)
		exec.WillReturnRows(rows)
		exec.WillReturnError(v.mockError)

		repo := NewRepository(db)
		report, err := repo.Get(v.id)
		if v.expectedError != err {
			t.Errorf("Get test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
			continue
		}

		for i, r := range report.Data {
			if v.expectedModel.Data[i] != r {
				t.Errorf("Get test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedModel, report)
			}
		}
	}
}
