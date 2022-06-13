package products

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db"
)

var lastProductID int = 1

type Repository interface {
	Insert(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetById(id int) (models.Product, error)
	Update(product models.Product) error
	Delete(id int) error
	LastId() (int, error)
}

type repository struct{
	db db.DB
}

func NewRepository(db db.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(product models.Product) (models.Product, error) {
	models.Products = append(models.Products, product)
	return product, nil
}

func (r *repository) GetAll() ([]models.Product, error) {
	return models.Products, nil
}

func (r *repository) GetById(id int) (models.Product, error) {
	p := models.Products
	for _, v := range p {
		if v.Id == id {
			return v, nil
		}
	}
	return models.Product{}, customerrors.ErrorInvalidID
}

func (r *repository) Update(product models.Product) error {
	p := models.Products
	for i, v := range p {
		if v.Id == product.Id {
			p[i] = product
			return nil
		}
	}
	return customerrors.ErrorInvalidID

}

func (r *repository) Delete(id int) error {
	p := models.Products
	for i, v := range p {
		if v.Id == id {
			models.Products = append(p[:i], p[i+1:]...)
			return nil
		}
	}
	return customerrors.ErrorInvalidID
}

func (r *repository) LastId() (int, error) {
	lastProductID++
	return lastProductID, nil
}
