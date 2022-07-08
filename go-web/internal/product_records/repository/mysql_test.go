package repository_test

import (
	"database/sql/driver"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/product_records/repository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestProductRecordsInsertErrorPrepare(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	record := models.ProductRecord{}
	query := "INSERT INTO product_records (last_update_date, purchase_prince, sale_price, product_id) VALUES (?, ?, ?, ?)"
	prep := mock.ExpectPrepare(query)
	prep.WillReturnError(sqlmock.ErrCancelled)

	_, err = repository.Insert(record)

	assert.Equal(t, sqlmock.ErrCancelled, err)
}

func TestProductRecordsInsertSuccess(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	record := models.ProductRecord{}
	query := "INSERT INTO product_records (last_update_date, purchase_prince, sale_price, product_id) VALUES (?, ?, ?, ?)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(record.LastUpdateDate, record.PurchasePrince, record.SalePrice, record.ProductId).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repository.Insert(record)

	assert.Equal(t, nil, err)
}

func TestProductRecordsInsertErrorExec(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	record := models.ProductRecord{}
	query := "INSERT INTO product_records (last_update_date, purchase_prince, sale_price, product_id) VALUES (?, ?, ?, ?)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(record.LastUpdateDate, record.PurchasePrince, record.SalePrice, record.ProductId).WillReturnError(sqlmock.ErrCancelled)

	_, err = repository.Insert(record)

	assert.Equal(t, sqlmock.ErrCancelled, err)
}

func TestProductRecordsInsertErrorLastId(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	record := models.ProductRecord{}
	query := "INSERT INTO product_records (last_update_date, purchase_prince, sale_price, product_id) VALUES (?, ?, ?, ?)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(record.LastUpdateDate, record.PurchasePrince, record.SalePrice, record.ProductId).WillReturnResult(driver.ResultNoRows)

	_, err = repository.Insert(record)

	assert.NotNil(t, err)
}

//
//
//
func TestProductRecordsGetErrorPrepare(t *testing.T) {
	query := "SELECT pr.product_id, p.description, COUNT(*) AS records_count FROM product_records pr INNER JOIN products p ON pr.product_id = p.id WHERE pr.product_id > ? GROUP BY pr.product_id"
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	mock.ExpectPrepare(query).WillReturnError(sqlmock.ErrCancelled)

	_, err = repository.GetByProductId(0)

	assert.Equal(t, sqlmock.ErrCancelled, err)
}

func TestProductRecordsGetAllSuccess(t *testing.T) {
	query := "SELECT pr.product_id, p.description, COUNT(*) AS records_count FROM product_records pr INNER JOIN products p ON pr.product_id = p.id WHERE pr.product_id > ? GROUP BY pr.product_id"
	expectRes := models.ProductsRecordsResponse{Records: []models.ProductRecordsResponse{
		{ProductId: 1, Description: "teste", RecordsCount: 1},
		{ProductId: 2, Description: "teste", RecordsCount: 2},
	}}
	mockRes := sqlmock.NewRows([]string{"product_id", "description", "records_count"}).
		AddRow(1, "teste", 1).
		AddRow(2, "teste", 2)
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(0).WillReturnRows(mockRes)

	res, err := repository.GetByProductId(0)

	assert.Equal(t, expectRes, res)
	assert.Equal(t, nil, err)
}

func TestProductRecordsGetAllScanError(t *testing.T) {
	query := "SELECT pr.product_id, p.description, COUNT(*) AS records_count FROM product_records pr INNER JOIN products p ON pr.product_id = p.id WHERE pr.product_id > ? GROUP BY pr.product_id"
	expectRes := models.ProductsRecordsResponse{}
	mockRes := sqlmock.NewRows([]string{"product_id", "description", "records_count", "foo"}).
		AddRow(1, "teste", 1, 0).
		AddRow(2, "teste", 2, 0)
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(0).WillReturnRows(mockRes)

	res, err := repository.GetByProductId(0)

	assert.Equal(t, expectRes, res)
	assert.Equal(t, "sql: expected 4 destination arguments in Scan, not 3", err.Error())
}

func TestProductRecordsGetAllQueryError(t *testing.T) {
	query := "SELECT pr.product_id, p.description, COUNT(*) AS records_count FROM product_records pr INNER JOIN products p ON pr.product_id = p.id WHERE pr.product_id > ? GROUP BY pr.product_id"
	expectRes := models.ProductsRecordsResponse{}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(0).WillReturnError(customerrors.ErrorInvalidDB)

	res, err := repository.GetByProductId(0)

	assert.Equal(t, expectRes, res)
	assert.Equal(t, customerrors.ErrorInvalidDB, err)
}

func TestProductRecordsGetIdSuccess(t *testing.T) {
	query := "SELECT pr.product_id, p.description, COUNT(*) AS records_count FROM product_records pr INNER JOIN products p ON pr.product_id = p.id WHERE pr.product_id = ? GROUP BY pr.product_id"
	expectRes := models.ProductsRecordsResponse{Records: []models.ProductRecordsResponse{
		{ProductId: 1, Description: "teste", RecordsCount: 1},
	}}
	mockRes := sqlmock.NewRows([]string{"product_id", "description", "records_count"}).
		AddRow(1, "teste", 1)
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepositoryProductRecord(db)
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(1).WillReturnRows(mockRes)

	res, err := repository.GetByProductId(1)

	assert.Equal(t, expectRes, res)
	assert.Equal(t, nil, err)
}
