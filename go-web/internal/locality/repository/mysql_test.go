package repository_test

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/locality/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestStoreLocality(t *testing.T) {
	u := models.Locality{Id: "1", Locality_name: "Buritizeiro", Province_name: "Minas Gerais", Country_name: "Brazil"}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT a.id FROM provinces a INNER JOIN countries b ON a.id_country_fk = b.Id AND b.country_name = ? AND a.province_name = ?;").
		WithArgs(u.Country_name, u.Province_name).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectPrepare("INSERT INTO `localities` (`id`, `locality_name`, `province_id`) VALUES (?, ?, ?)").ExpectExec().WithArgs(u.Id, u.Locality_name, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewSQLrepository(db)

	result, err := repo.Store(u)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestStoreLocalityErrorPrepare(t *testing.T) {
	u := models.Locality{Id: "1", Locality_name: "Buritizeiro", Province_name: "Minas Gerais", Country_name: "Brazil"}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT a.id FROM provinces a INNER JOIN countries b ON a.id_country_fk = b.Id AND b.country_name = ? AND a.province_name = ?;").
		WithArgs(u.Country_name, u.Province_name).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectPrepare("INSERT INTO `localities` (`id`, `locality_name`, `province_id`) VALUES (?, ?, ?)").
		WillReturnError(customerrors.ErrorInvalidDB)

	repo := repository.NewSQLrepository(db)

	result, err := repo.Store(u)
	assert.Equal(t, err, customerrors.ErrorInvalidDB)
	assert.NotNil(t, result)
}

func TestStoreLocalityErrorExec(t *testing.T) {
	u := models.Locality{Id: "1", Locality_name: "Buritizeiro", Province_name: "Minas Gerais", Country_name: "Brazil"}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT a.id FROM provinces a INNER JOIN countries b ON a.id_country_fk = b.Id AND b.country_name = ? AND a.province_name = ?;").
		WithArgs(u.Country_name, u.Province_name).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectPrepare("INSERT INTO `localities` (`id`, `locality_name`, `province_id`) VALUES (?, ?, ?)").ExpectExec().WillReturnError(customerrors.ErrorInvalidDB)

	repo := repository.NewSQLrepository(db)

	result, err := repo.Store(u)
	assert.Equal(t, err, customerrors.ErrorInvalidDB)
	assert.NotNil(t, result)
}

func TestStoreLocalityErrorQuery(t *testing.T) {
	u := models.Locality{Id: "1", Locality_name: "Buritizeiro", Province_name: "Minas Gerais", Country_name: "Brazil"}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	mock.ExpectQuery("SELECT a.id FROM provinces a INNER JOIN countries b ON a.id_country_fk = b.Id AND b.country_name = ? AND a.province_name = ?;").WillReturnError(customerrors.ErrorConflict)

	repo := repository.NewSQLrepository(db)

	result, err := repo.Store(u)
	assert.Equal(t, err, customerrors.ErrorConflict)
	assert.NotNil(t, result)
}
