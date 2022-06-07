package products

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
)

type Repository interface {
	Insert(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetById(id int) (models.Product, error)
	Update(product models.Product) error
	Delete(id int) error
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
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
	return models.Product{}, errors.New("id não encontrado")
}

func (r *repository) Update(product models.Product) error {
	p := models.Products
	for i, v := range p {
		if v.Id == product.Id {
			p[i] = product
			return nil
		}
	}
	return errors.New("id não encontrado")

}

func (r *repository) Delete(id int) error {
	p := models.Products
	for i, v := range p {
		if v.Id == id {
			models.Products = append(p[:i], p[i+1:]...)
			return nil
		}
	}
	return errors.New("id não encontrado")
}

func (r *repository) LastId() (int, error) {
	ts := models.Products
	maxId := ts[0].Id
	for i := 1; i <= len(ts)-1; i++ {
		if ts[i].Id > maxId {
			maxId = ts[i].Id
		}
	}
	return maxId + 1, nil
}
