package repository_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"mercado-frescos-time-7/go-web/internal/sections/domain"
	"mercado-frescos-time-7/go-web/internal/sections/repository"
	"regexp"
	"testing"
)

func TestGetAllSections(t *testing.T) {
	mockExpectResponse := &domain.Sections{
		Sections: []domain.Section{
			{
				Id:                 1,
				SectionNumber:      1,
				CurrentTemperature: 1,
				MinimumTemperature: 1,
				CurrentCapacity:    1,
				MinimumCapacity:    1,
				MaximumCapacity:    1,
				WarehouseId:        1,
				ProductTypeId:      1,
			},
			{
				Id:                 2,
				SectionNumber:      2,
				CurrentTemperature: 10,
				MinimumTemperature: 2,
				CurrentCapacity:    100,
				MinimumCapacity:    10,
				MaximumCapacity:    150,
				WarehouseId:        2,
				ProductTypeId:      2,
			},
		},
	}
	t.Run("GetAll Sections", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rowsMockResponse := sqlmock.NewRows([]string{
			"id", "section_number", "current_capacity", "current_temperature",
			"maximum_capacity", "minimum_capacity", "minimum_temperature",
			"product_type_id", "warehouse_id"}).
			AddRow(1, 1, 1, 1, 1, 1, 1, 1, 1).
			AddRow(2, 2, 10, 2, 100, 10, 150, 2, 2)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM sections;")).
			WillReturnRows(rowsMockResponse)

		repo := repository.NewRepositorySection(db)

		res, err := repo.GetAll(context.TODO())
		assert.Equal(t, res, mockExpectResponse)
		assert.NoError(t, err)
	})

	t.Run("Error Query Context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM sections;")).
			WillReturnError(sqlmock.ErrCancelled)

		_, err = repo.GetAll(context.TODO())

		assert.Equal(t, sqlmock.ErrCancelled, err)
	})

	t.Run("Error Scan Query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rowsMockResponse := sqlmock.NewRows([]string{
			"id", "section_number", "current_capacity", "current_temperature",
			"maximum_capacity", "minimum_capacity", "minimum_temperature",
			"product_type_id"}).
			AddRow(-1, 1, 1, 1, 1, 1, 1, 1).
			AddRow(2, 2, 10, 2, 100, 10, 150, 2)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM sections;")).
			WillReturnRows(rowsMockResponse)

		repo := repository.NewRepositorySection(db)

		res, err := repo.GetAll(context.TODO())

		assert.Equal(t, &domain.Sections{}, res)
		assert.Error(t, err)
	})
}

func TestGetByIdSection(t *testing.T) {
	mockExpectResponse := &domain.Section{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	t.Run("GetById Section", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rowMockResponse := sqlmock.NewRows([]string{
			"id", "section_number", "current_capacity", "current_temperature",
			"maximum_capacity", "minimum_capacity", "minimum_temperature",
			"product_type_id", "warehouse_id"}).
			AddRow(1, 1, 1, 1, 1, 1, 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM sections WHERE id = ?;")).
			WithArgs(1).WillReturnRows(rowMockResponse)

		repo := repository.NewRepositorySection(db)

		res, err := repo.GetById(context.TODO(), 1)

		assert.Equal(t, res, mockExpectResponse)
		assert.NoError(t, err)
	})

	t.Run("Error No Rows", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM sections WHERE id = ?;")).
			WithArgs(1).WillReturnError(sql.ErrNoRows)

		res, err := repo.GetById(context.TODO(), 1)

		assert.Equal(t, sql.ErrNoRows, err)
		assert.Equal(t, &domain.Section{}, res)
	})

	t.Run("Error No Rows", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM sections WHERE id = ?;")).
			WithArgs(1).WillReturnError(sqlmock.ErrCancelled)

		res, err := repo.GetById(context.TODO(), 1)

		assert.Equal(t, sqlmock.ErrCancelled, err)
		assert.Equal(t, &domain.Section{}, res)
	})

}

func TestStoreSection(t *testing.T) {
	mockExpectResponse := &domain.Section{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	t.Run("Create Section", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO sections"+
			"(section_number, current_temperature, minimum_temperature, current_capacity,"+
			"minimum_capacity, maximum_capacity, warehouse_id, product_type_id)"+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?);")).ExpectExec().
			WithArgs(mockExpectResponse.SectionNumber, mockExpectResponse.CurrentTemperature,
				mockExpectResponse.MinimumTemperature, mockExpectResponse.CurrentCapacity,
				mockExpectResponse.MinimumCapacity, mockExpectResponse.MaximumCapacity,
				mockExpectResponse.WarehouseId, mockExpectResponse.ProductTypeId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res, err := repo.Store(context.TODO(), mockExpectResponse)

		assert.Equal(t, res, mockExpectResponse)
		assert.NoError(t, err)
	})

	t.Run("Error Exec Query Section", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO sections"+
			"(section_number, current_temperature, minimum_temperature, current_capacity,"+
			"minimum_capacity, maximum_capacity, warehouse_id, product_type_id)"+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?);")).ExpectExec().
			WithArgs(mockExpectResponse.SectionNumber, mockExpectResponse.CurrentTemperature,
				mockExpectResponse.MinimumTemperature, mockExpectResponse.CurrentCapacity,
				mockExpectResponse.MinimumCapacity, mockExpectResponse.MaximumCapacity,
				mockExpectResponse.WarehouseId, mockExpectResponse.ProductTypeId).
			WillReturnError(sqlmock.ErrCancelled)

		res, err := repo.Store(context.TODO(), mockExpectResponse)

		assert.Equal(t, res, &domain.Section{})
		assert.Error(t, err)
	})

	t.Run("Error Prepare Query Section", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO sections" +
			"(section_number, current_temperature, minimum_temperature, current_capacity," +
			"minimum_capacity, maximum_capacity, warehouse_id, product_type_id)" +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?);")).
			WillReturnError(sqlmock.ErrCancelled)

		res, err := repo.Store(context.TODO(), &domain.Section{})

		assert.Equal(t, res, &domain.Section{})
		assert.Error(t, err)
	})
	t.Run("Error Last Id", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO sections"+
			"(section_number, current_temperature, minimum_temperature, current_capacity,"+
			"minimum_capacity, maximum_capacity, warehouse_id, product_type_id)"+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?);")).ExpectExec().
			WithArgs(mockExpectResponse.SectionNumber, mockExpectResponse.CurrentTemperature,
				mockExpectResponse.MinimumTemperature, mockExpectResponse.CurrentCapacity,
				mockExpectResponse.MinimumCapacity, mockExpectResponse.MaximumCapacity,
				mockExpectResponse.WarehouseId, mockExpectResponse.ProductTypeId).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		res, err := repo.Store(context.TODO(), mockExpectResponse)

		assert.Equal(t, &domain.Section{}, res)
		assert.Error(t, err)
	})

}

func TestUpdateSection(t *testing.T) {
	sendMockUpdate := &domain.Section{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	t.Run("Update Section with no errors", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE sections SET "+
			"section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?,"+
			"minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?;")).
			ExpectExec().WithArgs(sendMockUpdate.SectionNumber, sendMockUpdate.CurrentTemperature,
			sendMockUpdate.MinimumTemperature, sendMockUpdate.CurrentCapacity,
			sendMockUpdate.MinimumCapacity, sendMockUpdate.MaximumCapacity,
			sendMockUpdate.WarehouseId, sendMockUpdate.ProductTypeId, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		_, err = repo.Update(context.TODO(), sendMockUpdate)
		sendMockUpdate.Id = 1

		assert.NoError(t, err)

	})
	t.Run("update prepare error", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE sections SET " +
			"section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?," +
			"minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?;")).
			WillReturnError(sqlmock.ErrCancelled)

		_, err = repo.Update(context.TODO(), sendMockUpdate)

		assert.Equal(t, err, sqlmock.ErrCancelled)
	})
	t.Run("update exec context error", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE sections SET "+
			"section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?,"+
			"minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?;")).
			ExpectExec().WithArgs(sendMockUpdate.SectionNumber, sendMockUpdate.CurrentTemperature,
			sendMockUpdate.MinimumTemperature, sendMockUpdate.CurrentCapacity,
			sendMockUpdate.MinimumCapacity, sendMockUpdate.MaximumCapacity,
			sendMockUpdate.WarehouseId, sendMockUpdate.ProductTypeId, 1).
			WillReturnError(sqlmock.ErrCancelled)

		_, err = repo.Update(context.TODO(), sendMockUpdate)

		assert.Equal(t, err, sqlmock.ErrCancelled)
	})

}

func TestDeleteSection(t *testing.T) {
	t.Run("Delete Section with no errors", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM sections WHERE id=?;")).
			ExpectExec().WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Delete(context.TODO(), 1)

		assert.NoError(t, err)

	})
	t.Run("Delete prepare error", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM sections WHERE id=?;")).
			WillReturnError(sqlmock.ErrCancelled)

		err = repo.Delete(context.TODO(), 1)

		assert.Error(t, err, sqlmock.ErrCancelled)
	})
	t.Run("Delete exec error", func(*testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewRepositorySection(db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM sections WHERE id=?;")).
			ExpectExec().WithArgs(1).
			WillReturnError(sqlmock.ErrCancelled)

		err = repo.Delete(context.TODO(), 1)

		assert.Error(t, err, sqlmock.ErrCancelled)
	})

}

func TestGetReportProductsById(t *testing.T) {
	mockExpectResponse := &domain.ProductReports{
		ProductReports: []domain.ProductReport{
			{
				SectionId:     1,
				SectionNumber: 1,
				ProductsCount: 100,
			},
			{
				SectionId:     2,
				SectionNumber: 2,
				ProductsCount: 100,
			},
		},
	}
	mockExpectResponseSingle := &domain.ProductReports{
		ProductReports: []domain.ProductReport{
			{
				SectionId:     1,
				SectionNumber: 1,
				ProductsCount: 100,
			},
		},
	}
	t.Run("Get All Report Products", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rowsMockResponse := sqlmock.NewRows([]string{
			"section_id", "section_number", "products_count"}).
			AddRow(1, 1, 100).
			AddRow(2, 2, 100)

		mock.ExpectPrepare(regexp.QuoteMeta("SELECT s.id, s.section_number, COUNT(*) as products_count from sections s Inner join products_batches p on s.id = p.section_id WHERE s.id > ? GROUP BY(s.id);")).
			ExpectQuery().
			WithArgs(0).
			WillReturnRows(rowsMockResponse)

		repo := repository.NewRepositorySection(db)

		res, err := repo.GetReportProducts(context.TODO(), 0)
		assert.Equal(t, res, mockExpectResponse)
		assert.NoError(t, err)
	})
	t.Run("Get Report Products prepare error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta("SELECT s.id, s.section_number, COUNT(*) as products_count from sections s Inner join products_batches p on s.id = p.section_id WHERE s.id > ? GROUP BY(s.id);")).
			WillReturnError(sqlmock.ErrCancelled)

		repo := repository.NewRepositorySection(db)

		res, err := repo.GetReportProducts(context.TODO(), 0)
		assert.Equal(t, &domain.ProductReports{}, res)
		assert.Error(t, err)
	})
	t.Run("Get Report Products query error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta("SELECT s.id, s.section_number, COUNT(*) as products_count from sections s Inner join products_batches p on s.id = p.section_id WHERE s.id > ? GROUP BY(s.id);")).
			ExpectQuery().
			WithArgs(0).
			WillReturnError(sqlmock.ErrCancelled)

		repo := repository.NewRepositorySection(db)

		res, err := repo.GetReportProducts(context.TODO(), 0)
		assert.Equal(t, &domain.ProductReports{}, res)
		assert.Error(t, err)
	})
	t.Run("Get Report Products scan", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rowsMockResponse := sqlmock.NewRows([]string{
			"section_id", "section_number", "products_count"}).
			AddRow(-1, 1, 100).
			AddRow("", 2, -100)

		mock.ExpectPrepare(regexp.QuoteMeta("SELECT s.id, s.section_number, COUNT(*) as products_count from sections s Inner join products_batches p on s.id = p.section_id WHERE s.id > ? GROUP BY(s.id);")).
			ExpectQuery().
			WithArgs(0).
			WillReturnRows(rowsMockResponse)

		repo := repository.NewRepositorySection(db)

		res, err := repo.GetReportProducts(context.TODO(), 0)
		assert.Equal(t, &domain.ProductReports{}, res)
		assert.Error(t, err)
	})
	t.Run("Get Report Products", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rowsMockResponse := sqlmock.NewRows([]string{
			"section_id", "section_number", "products_count"}).
			AddRow(1, 1, 100)
		mock.ExpectPrepare(regexp.QuoteMeta("SELECT s.id, s.section_number, COUNT(*) as products_count from sections s INNER JOIN products_batches pb ON s.id = pb.section_id WHERE s.id = ? GROUP BY(s.id);")).
			ExpectQuery().
			WithArgs(1).
			WillReturnRows(rowsMockResponse)

		repo := repository.NewRepositorySection(db)

		res, err := repo.GetReportProducts(context.TODO(), 1)
		assert.Equal(t, res, mockExpectResponseSingle)
		assert.NoError(t, err)
	})

}
