package repository

import (
	"database/sql"
	"fmt"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/logger"
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
		logger.Save(err.Error())
		return models.Product{}, err
	}

	res, err := stmt.Exec(&product.Description, &product.ExpirationRate, &product.FreezingRate, &product.Height, &product.Length, &product.NetWeight, &product.ProductCode, &product.RecommendedFreezingTemperature, &product.Width, &product.ProductTypeId, &product.SellerId)
	if err != nil {
		logger.Save(err.Error())
		return models.Product{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		logger.Save(err.Error())
		return models.Product{}, err
	}
	product.Id = int(lastId)

	logger.Save(fmt.Sprintf(logger.ProductsCreated, lastId))
	return product, nil
}

func (r *repositoryMysql) GetAll() (models.Products, error) {
	products := models.Products{}
	rows, err := r.db.Query(sqlGetAll)
	if err != nil {
		logger.Save(err.Error())
		return models.Products{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Description, &product.ExpirationRate, &product.FreezingRate, &product.Height, &product.Length, &product.NetWeight, &product.ProductCode, &product.RecommendedFreezingTemperature, &product.Width, &product.ProductTypeId, &product.SellerId)
		if err != nil {
			logger.Save(err.Error())
			return models.Products{}, err
		}
		products.Products = append(products.Products, product)
	}

	logger.Save(logger.ProductssResquested)
	return products, nil
}

func (r *repositoryMysql) GetById(id int) (models.Product, error) {
	product := models.Product{}
	stmt, err := r.db.Prepare(sqlGetById)
	if err != nil {
		logger.Save(err.Error())
		return models.Product{}, err
	}
	err = stmt.QueryRow(id).Scan(&product.Id, &product.Description, &product.ExpirationRate, &product.FreezingRate, &product.Height, &product.Length, &product.NetWeight, &product.ProductCode, &product.RecommendedFreezingTemperature, &product.Width, &product.ProductTypeId, &product.SellerId)
	if err != nil {
		logger.Save(err.Error())
		return models.Product{}, err
	}
	logger.Save(fmt.Sprintf(logger.ProductsRequested, id))
	return product, nil
}

func (r *repositoryMysql) Update(product models.Product) error {
	stmt, err := r.db.Prepare(sqlUpdate)
	if err != nil {
		logger.Save(err.Error())
		return err
	}
	_, err = stmt.Exec(&product.Description, &product.ExpirationRate, &product.FreezingRate, &product.Height, &product.Length, &product.NetWeight, &product.ProductCode, &product.RecommendedFreezingTemperature, &product.Width, &product.ProductTypeId, &product.SellerId, &product.Id)
	if err != nil {
		logger.Save(err.Error())
		return err
	}
	logger.Save(fmt.Sprintf(logger.ProductsChanged, product.Id))
	return nil
}

func (r *repositoryMysql) Delete(id int) error {
	stmt, err := r.db.Prepare(sqlDelete)
	if err != nil {
		logger.Save(err.Error())
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		logger.Save(err.Error())
		return err
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		logger.Save(err.Error())
		return customerrors.ErrorInvalidID
	} else if err != nil {
		logger.Save(err.Error())
		return customerrors.ErrorInvalidDB
	}
	logger.Save(fmt.Sprintf(logger.ProductsDeleted, id))
	return nil
}
