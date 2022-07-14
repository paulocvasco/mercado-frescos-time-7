package repository

import (
	"database/sql"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type repositoryMysql struct {
	db *sql.DB
}

func NewRepositoryMysql(db *sql.DB) Repository {
	return &repositoryMysql{
		db: db,
	}
}

func (r *repositoryMysql) Insert(product models.Product) (models.Product, error) {
	stmt, err := r.db.Prepare(sqlStore)
	if err != nil {
		return models.Product{}, err
	}

	res, err := stmt.Exec(&product.Description, &product.ExpirationRate, &product.FreezingRate, &product.Height, &product.Length, &product.NetWeight, &product.ProductCode, &product.RecommendedFreezingTemperature, &product.Width, &product.ProductTypeId, &product.SellerId)
	if err != nil {
		return models.Product{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}
	product.Id = int(lastId)

	return product, nil
}

func (r *repositoryMysql) GetAll() (models.Products, error) {
	products := models.Products{}
	rows, err := r.db.Query(sqlGetAll)
	if err != nil {
		return models.Products{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Description, &product.ExpirationRate, &product.FreezingRate, &product.Height, &product.Length, &product.NetWeight, &product.ProductCode, &product.RecommendedFreezingTemperature, &product.Width, &product.ProductTypeId, &product.SellerId)
		if err != nil {
			return models.Products{}, err
		}
		products.Products = append(products.Products, product)
	}

	return products, nil
}

func (r *repositoryMysql) GetById(id int) (models.Product, error) {
	product := models.Product{}
	stmt, err := r.db.Prepare(sqlGetById)
	if err != nil {
		return models.Product{}, err
	}
	err = stmt.QueryRow(id).Scan(&product.Id, &product.Description, &product.ExpirationRate, &product.FreezingRate, &product.Height, &product.Length, &product.NetWeight, &product.ProductCode, &product.RecommendedFreezingTemperature, &product.Width, &product.ProductTypeId, &product.SellerId)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *repositoryMysql) Update(product models.Product) error {
	stmt, err := r.db.Prepare(sqlUpdate)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(&product.Description, &product.ExpirationRate, &product.FreezingRate, &product.Height, &product.Length, &product.NetWeight, &product.ProductCode, &product.RecommendedFreezingTemperature, &product.Width, &product.ProductTypeId, &product.SellerId, &product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryMysql) Delete(id int) error {
	stmt, err := r.db.Prepare(sqlDelete)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		return customerrors.ErrorInvalidID
	} else if err != nil {
		return customerrors.ErrorInvalidDB
	}
	return nil
}
