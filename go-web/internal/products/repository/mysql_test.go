package repository_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/products/repository"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/go-sql-driver/mysql"
)

func TestInsertProduct(t *testing.T) {
	model := models.Product{
		Id:                             0,
		Description:                    "mod",
		ExpirationRate:                 10,
		FreezingRate:                   20,
		Height:                         6.40,
		Length:                         4.50,
		NetWeight:                      3.40,
		ProductCode:                    "ssd-Editado",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	t.Run("Should return product", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("INSERT INTO products (description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		stmtMock.ExpectExec().WithArgs(model.Description, model.ExpirationRate, model.FreezingRate, model.Height, model.Length, model.NetWeight, model.ProductCode, model.RecommendedFreezingTemperature, model.Width, model.ProductTypeId, model.SellerId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		res, err := repo.Insert(model)
		modelExpected := model
		modelExpected.Id = 1

		assert.Equal(t, modelExpected, res)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return error in prepare statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("INSERT INTO products (description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		stmtMock.WillReturnError(errors.New("prepare error"))
		res, err := repo.Insert(model)

		assert.Equal(t, models.Product{}, res)
		assert.Equal(t, "prepare error", err.Error())
	})

	t.Run("Should return error in exec statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("INSERT INTO products (description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		stmtMock.ExpectExec().WithArgs(model.Description, model.ExpirationRate, model.FreezingRate, model.Height, model.Length, model.NetWeight, model.ProductCode, model.RecommendedFreezingTemperature, model.Width, model.ProductTypeId, model.SellerId).
			WillReturnError(&mysql.MySQLError{Number: 1062, Message: "err"})
		res, err := repo.Insert(model)
		modelExpected := model
		modelExpected.Id = 1

		assert.Equal(t, models.Product{}, res)
		assert.Equal(t, "Error 1062: err", err.Error())
	})

	t.Run("Should return error in lastId statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("INSERT INTO products (description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		stmtMock.ExpectExec().WithArgs(model.Description, model.ExpirationRate, model.FreezingRate, model.Height, model.Length, model.NetWeight, model.ProductCode, model.RecommendedFreezingTemperature, model.Width, model.ProductTypeId, model.SellerId).
			WillReturnResult(sqlmock.NewErrorResult(driver.ErrBadConn))
		res, err := repo.Insert(model)

		assert.Equal(t, models.Product{}, res)
		assert.Equal(t, driver.ErrBadConn, err)
	})
}

func TestGetAllProducts(t *testing.T) {
	model1 := models.Product{
		Id:                             1,
		Description:                    "foo",
		ExpirationRate:                 1.1,
		FreezingRate:                   1.1,
		Height:                         1.1,
		Length:                         1.1,
		NetWeight:                      1.1,
		ProductCode:                    "bar",
		RecommendedFreezingTemperature: 1.1,
		Width:                          1.1,
		ProductTypeId:                  1,
		SellerId:                       1,
	}
	model2 := models.Product{
		Id:                             2,
		Description:                    "foo",
		ExpirationRate:                 1.1,
		FreezingRate:                   1.1,
		Height:                         1.1,
		Length:                         1.1,
		NetWeight:                      1.1,
		ProductCode:                    "foo",
		RecommendedFreezingTemperature: 1.1,
		Width:                          1.1,
		ProductTypeId:                  1,
		SellerId:                       1,
	}
	t.Run("Should return success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		rows := mock.NewRows([]string{
			"id", "description", "expiration_rate", "freezing_rate",
			"heigth", "length", "net_weight", "product_code",
			"recommended_freezing_temperature", "width",
			"product_type_id", "seller_id",
		}).AddRow(1, "foo", 1.1, 1.1, 1.1, 1.1, 1.1, "bar", 1.1, 1.1, 1, 1).
			AddRow(2, "foo", 1.1, 1.1, 1.1, 1.1, 1.1, "foo", 1.1, 1.1, 1, 1)
		mock.ExpectQuery("SELECT * FROM products").
			WillReturnRows(rows)
		repo := repository.NewRepositoryMysql(db)

		res, err := repo.GetAll()
		expectRes := models.Products{Products: []models.Product{
			model1, model2,
		}}
		assert.Equal(t, expectRes, res)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return error in query statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		mock.ExpectQuery("SELECT * FROM products").
			WillReturnError(driver.ErrBadConn)
		repo := repository.NewRepositoryMysql(db)

		res, err := repo.GetAll()

		assert.Equal(t, models.Products{}, res)
		assert.Equal(t, "expected a connection to be available, but it is not", err.Error())
	})

	t.Run("Should return error in scan statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		rows := mock.NewRows([]string{
			"id", "description", "expiration_rate", "freezing_rate",
			"heigth", "length", "net_weight", "product_code",
			"recommended_freezing_temperature", "width",
			"product_type_id", "seller_id", "foo",
		}).AddRow(1, "foo", 1.1, 1.1, 1.1, 1.1, 1.1, "bar", 1.1, 1.1, 1, 1, "foo")
		mock.ExpectQuery("SELECT * FROM products").
			WillReturnRows(rows)
		repo := repository.NewRepositoryMysql(db)

		res, err := repo.GetAll()

		assert.Equal(t, models.Products{}, res)
		assert.Equal(t, "sql: expected 13 destination arguments in Scan, not 12", err.Error())
	})
}

func TestGetProductById(t *testing.T) {
	modelExpected := models.Product{
		Id:                             1,
		Description:                    "foo",
		ExpirationRate:                 1.1,
		FreezingRate:                   1.1,
		Height:                         1.1,
		Length:                         1.1,
		NetWeight:                      1.1,
		ProductCode:                    "bar",
		RecommendedFreezingTemperature: 1.1,
		Width:                          1.1,
		ProductTypeId:                  1,
		SellerId:                       1,
	}
	t.Run("Should return product", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}
		rows := mock.NewRows([]string{
			"id", "description", "expiration_rate", "freezing_rate",
			"heigth", "length", "net_weight", "product_code",
			"recommended_freezing_temperature", "width",
			"product_type_id", "seller_id",
		}).AddRow(1, "foo", 1.1, 1.1, 1.1, 1.1, 1.1, "bar", 1.1, 1.1, 1, 1)

		repo := repository.NewRepositoryMysql(db)
		mock.ExpectPrepare("SELECT * FROM products WHERE id = ?").
			ExpectQuery().WithArgs(1).WillReturnRows(rows)
		res, err := repo.GetById(1)

		assert.Equal(t, modelExpected, res)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return error in prepare statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		mock.ExpectPrepare("SELECT * FROM products WHERE id = ?").
			WillReturnError(driver.ErrRemoveArgument)
		res, err := repo.GetById(1)

		assert.Equal(t, models.Product{}, res)
		assert.Equal(t, driver.ErrRemoveArgument, err)
	})

	t.Run("Should return error in query statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		mock.ExpectPrepare("SELECT * FROM products WHERE id = ?").
			ExpectQuery().WithArgs(1).WillReturnError(sql.ErrNoRows)
		res, err := repo.GetById(1)

		assert.Equal(t, models.Product{}, res)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	model := models.Product{
		Id:                             1,
		Description:                    "mod",
		ExpirationRate:                 10,
		FreezingRate:                   20,
		Height:                         6.40,
		Length:                         4.50,
		NetWeight:                      3.40,
		ProductCode:                    "ssd-Editado",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	t.Run("Should not return any errors", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=? WHERE id=?")
		stmtMock.ExpectExec().WithArgs(model.Description, model.ExpirationRate, model.FreezingRate, model.Height, model.Length, model.NetWeight, model.ProductCode, model.RecommendedFreezingTemperature, model.Width, model.ProductTypeId, model.SellerId, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err = repo.Update(model)

		assert.Equal(t, nil, err)
	})

	t.Run("Should return error in prepare statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=? WHERE id=?")
		stmtMock.WillReturnError(driver.ErrRemoveArgument)
		err = repo.Update(model)

		assert.Equal(t, driver.ErrRemoveArgument, err)
	})

	t.Run("Should return error in exec statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=? WHERE id=?")
		stmtMock.ExpectExec().WithArgs(model.Description, model.ExpirationRate, model.FreezingRate, model.Height, model.Length, model.NetWeight, model.ProductCode, model.RecommendedFreezingTemperature, model.Width, model.ProductTypeId, model.SellerId, 1).
			WillReturnError(driver.ErrBadConn)
		err = repo.Update(model)

		assert.Equal(t, "expected a connection to be available, but it is not", err.Error())
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Should not return any errors", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("DELETE FROM products WHERE id=?")
		stmtMock.ExpectExec().WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err = repo.Delete(1)

		assert.Equal(t, nil, err)
	})

	t.Run("Should return error in prepare statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		mock.ExpectPrepare("DELETE FROM products WHERE id=?").WillReturnError(driver.ErrRemoveArgument)

		err = repo.Delete(0)

		assert.Equal(t, driver.ErrRemoveArgument, err)
	})

	t.Run("Should return error in exec statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("DELETE FROM products WHERE id=?")
		stmtMock.ExpectExec().WithArgs(1).
			WillReturnError(driver.ErrBadConn)
		err = repo.Delete(1)

		assert.Equal(t, "expected a connection to be available, but it is not", err.Error())
	})

	t.Run("Should return invalid id after exec statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("DELETE FROM products WHERE id=?")
		stmtMock.ExpectExec().WithArgs().
			WillReturnResult(sqlmock.NewResult(0, 0))
		err = repo.Delete(1)

		assert.Equal(t, customerrors.ErrorInvalidID, err)
	})

	t.Run("Should return error after exec statment", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewRepositoryMysql(db)
		stmtMock := mock.ExpectPrepare("DELETE FROM products WHERE id=?")
		stmtMock.ExpectExec().WithArgs(1).
			WillReturnResult(driver.ResultNoRows)
		err = repo.Delete(1)

		assert.Equal(t, customerrors.ErrorInvalidDB, err)
	})
}
