package repository_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/productBatch/domain"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/productBatch/repository"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSqlRepositoryCreateProductBatchOk(t *testing.T) {
	mockSendProductBatch := &domain.ProductBatch{
		BatchNumber:        1013,
		CurrentQuantity:    100,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    1,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		ProductId:          1,
		SectionId:          1,
	}

	mockResultProductBatch := &domain.ProductBatch{
		Id:                 1,
		BatchNumber:        1013,
		CurrentQuantity:    100,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    1,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		ProductId:          1,
		SectionId:          1,
	}

	queryInsert := `INSERT INTO products_batches (batch_number, current_quantity, current_tempertature,
					due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	t.Run("Create OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prepareQuery := mock.ExpectPrepare(regexp.QuoteMeta(queryInsert))

		prepareQuery.ExpectExec().WithArgs(
			mockSendProductBatch.BatchNumber,
			mockSendProductBatch.CurrentQuantity,
			mockSendProductBatch.CurrentTemperature,
			mockSendProductBatch.DueDate,
			mockSendProductBatch.InitialQuantity,
			mockSendProductBatch.ManufacturingDate,
			mockSendProductBatch.ManufacturingHour,
			mockSendProductBatch.MinimumTemperature,
			mockSendProductBatch.ProductId,
			mockSendProductBatch.SectionId,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := repository.NewRepositoryProductBatch(db)
		result, err := repo.CreateProductBatch(context.TODO(), mockSendProductBatch)
		assert.Equal(t, result, mockResultProductBatch)
		assert.NoError(t, err)
	})

	t.Run("Create Conflict", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		prepareQuery := mock.ExpectPrepare(regexp.QuoteMeta(queryInsert))

		prepareQuery.ExpectExec().WithArgs(
			mockSendProductBatch.BatchNumber,
			mockSendProductBatch.CurrentQuantity,
			mockSendProductBatch.CurrentTemperature,
			mockSendProductBatch.DueDate,
			mockSendProductBatch.InitialQuantity,
			mockSendProductBatch.ManufacturingDate,
			mockSendProductBatch.ManufacturingHour,
			mockSendProductBatch.MinimumTemperature,
			mockSendProductBatch.ProductId,
			mockSendProductBatch.SectionId,
		).WillReturnResult(sqlmock.NewErrorResult(sqlmock.ErrCancelled))

		repo := repository.NewRepositoryProductBatch(db)
		result, err := repo.CreateProductBatch(context.TODO(), &domain.ProductBatch{})
		assert.Equal(t, result, &domain.ProductBatch{})
		assert.Error(t, err)
	})

	t.Run("Create Error Prepare", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositoryProductBatch(db)

		prepareQuery := mock.ExpectPrepare(regexp.QuoteMeta(queryInsert))
		prepareQuery.WillReturnError(sqlmock.ErrCancelled)

		_, err = repo.CreateProductBatch(context.TODO(), &domain.ProductBatch{})

		assert.Equal(t, sqlmock.ErrCancelled, err)
	})

	t.Run("Create Error Last Id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositoryProductBatch(db)

		prepareQuery := mock.ExpectPrepare(regexp.QuoteMeta(queryInsert))
		prepareQuery.ExpectExec().WithArgs(
			mockSendProductBatch.BatchNumber,
			mockSendProductBatch.CurrentQuantity,
			mockSendProductBatch.CurrentTemperature,
			mockSendProductBatch.DueDate,
			mockSendProductBatch.InitialQuantity,
			mockSendProductBatch.ManufacturingDate,
			mockSendProductBatch.ManufacturingHour,
			mockSendProductBatch.MinimumTemperature,
			mockSendProductBatch.ProductId,
			mockSendProductBatch.SectionId,
		).WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		_, err = repo.CreateProductBatch(context.TODO(), mockSendProductBatch)

		assert.Error(t, err)
	})

}
