package repository

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db"
)

//go:generate mockery --name=Repository --output=./mock/mockRepository --outpkg=mockRepository
type Repository interface {
	Insert(product models.Product) (models.Product, error)
	GetAll() (models.Products, error)
	GetById(id int) (models.Product, error)
	Update(product models.Product) error
	Delete(id int) error
}

type repository struct {
	db db.DB
}

func NewRepository(db db.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(product models.Product) (models.Product, error) {
	var products models.ProductMetaData
	err := r.db.Load(&products)
	products.LastID++
	product.Id = products.LastID
	if err != nil {
		return models.Product{}, err
	}
	products.Content.Products = append(products.Content.Products, product)
	err = r.db.Save(products)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *repository) GetAll() (models.Products, error) {
	var products models.ProductMetaData
	err := r.db.Load(&products)
	if err != nil {
		return models.Products{}, err
	}
	return products.Content, nil
}

func (r *repository) GetById(id int) (models.Product, error) {
	var products models.ProductMetaData
	err := r.db.Load(&products)
	if err != nil {
		return models.Product{}, err
	}
	for _, v := range products.Content.Products {
		if v.Id == id {
			return v, nil
		}
	}
	return models.Product{}, customerrors.ErrorInvalidID
}

func (r *repository) Update(product models.Product) error {
	var products models.ProductMetaData
	err := r.db.Load(&products)
	if err != nil {
		return err
	}
	for i, v := range products.Content.Products {
		if v.Id == product.Id {
			products.Content.Products[i] = product
			err := r.db.Save(products)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return customerrors.ErrorInvalidID
}

func (r *repository) Delete(id int) error {
	var products models.ProductMetaData
	err := r.db.Load(&products)
	if err != nil {
		return err
	}
	for i, v := range products.Content.Products {
		if v.Id == id {
			products.Content.Products = append(products.Content.Products[:i], products.Content.Products[i+1:]...)
			err := r.db.Save(products)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return customerrors.ErrorInvalidID
}
