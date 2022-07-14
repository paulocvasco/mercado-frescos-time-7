package repository_test

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/seller"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/seller/repository"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {

	u := models.Seller{Cid: 195, Company_name: "Apple", Address: "Rua six", Telephone: "38998988978", LocalityID: "6700"}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO sellers (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").ExpectExec().
		WithArgs(u.Cid, u.Company_name, u.Address, u.Telephone, 6700).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewSQLrepository(db)

	result, err := repo.Store(u)
	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestStoreErrorConv(t *testing.T) {

	u := models.Seller{Cid: 195, Company_name: "Apple", Address: "Rua six", Telephone: "38998988978", LocalityID: "abcd"}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO sellers (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").ExpectExec().
		WithArgs(u.Cid, u.Company_name, u.Address, u.Telephone, 6700).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewSQLrepository(db)

	_, err = repo.Store(u)
	assert.NotNil(t, err)

}

func TestStoreErrorPrepare(t *testing.T) {

	u := models.Seller{Cid: 195, Company_name: "Apple", Address: "Rua six", Telephone: "38998988978", LocalityID: "6700"}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO sellers (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").WillReturnError(customerrors.ErrorInvalidDB)

	repo := repository.NewSQLrepository(db)

	result, err := repo.Store(u)
	assert.Equal(t, err, customerrors.ErrorInvalidDB)
	assert.NotNil(t, result)

}

func TestStoreErrorExec(t *testing.T) {

	u := models.Seller{Cid: 195, Company_name: "Apple", Address: "Rua six", Telephone: "38998988978", LocalityID: "6700"}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO sellers (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
		ExpectExec().
		WithArgs(u.Cid, u.Company_name, u.Address, u.Telephone, 6700).
		WillReturnError(customerrors.ErrorInvalidDB)

	repo := repository.NewSQLrepository(db)

	result, err := repo.Store(u)
	assert.Equal(t, err, customerrors.ErrorInvalidDB)
	assert.NotNil(t, result)

}

func TestEmptyFuncion(t *testing.T) {
	u := models.Seller{Cid: 195, Company_name: "Apple", Address: "Rua six", Telephone: "38998988978", LocalityID: "abcd"}

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	assert.NoError(t, err)

	repo := repository.NewSQLrepository(db)

	repo.CheckCid(1)
	repo.Delete(1)
	repo.GetAll()
	repo.GetId(1)
	repo.LastID()
	repo.Update(seller.Seller(u), 1)
}
