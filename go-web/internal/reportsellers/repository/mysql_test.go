package repository_test

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/reportsellers/repository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestStoreRepportAll(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(*) FROM sellers s INNER JOIN localities l ON s.locality_id = l.id WHERE s.locality_id > ? GROUP BY s.locality_id;").
		WithArgs(0).WillReturnRows(sqlmock.NewRows([]string{"id", "locality_name", "COUNT(*)"}).AddRow(1, "Juan", 5))

	repo := repository.NewSQLrepository(db)

	result, err := repo.ReportSellers(0)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestStoreRepportID(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(*) FROM sellers s INNER JOIN localities l ON s.locality_id = l.id WHERE s.locality_id = ? GROUP BY s.locality_id;").
		WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "locality_name", "COUNT(*)"}).AddRow(1, "Juan", 5))

	repo := repository.NewSQLrepository(db)

	result, err := repo.ReportSellers(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestStoreRepportAllErrorQuery(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(*) FROM sellers s INNER JOIN localities l ON s.locality_id = l.id WHERE s.locality_id > ? GROUP BY s.locality_id;").
		WithArgs(0).WillReturnError(customerrors.ErrorInvalidDB)

	repo := repository.NewSQLrepository(db)

	result, err := repo.ReportSellers(0)
	assert.Equal(t, err, customerrors.ErrorInvalidDB)
	assert.NotNil(t, result)
}

func TestStoreRepportAllErrorScan(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(*) FROM sellers s INNER JOIN localities l ON s.locality_id = l.id WHERE s.locality_id > ? GROUP BY s.locality_id;").
	WithArgs(0).WillReturnRows(sqlmock.NewRows([]string{"test"}).AddRow(1))

	repo := repository.NewSQLrepository(db)

	_, err = repo.ReportSellers(0)
	assert.Equal(t, errors.New("sql: expected 1 destination arguments in Scan, not 3"), err)

}
