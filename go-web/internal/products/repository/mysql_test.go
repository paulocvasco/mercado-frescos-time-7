package repository_test

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/products/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
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
	t.Run("Should return success", func(t *testing.T){
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
	t.Run("Should return product", func(t *testing.T) {
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
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Should return product", func(t *testing.T) {
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
}