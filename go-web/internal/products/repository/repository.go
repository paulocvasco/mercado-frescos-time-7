package repository

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/products/model"
)

type Repository interface {
	Insert(product model.Product) (model.Product, error)
	GetAll() ([]model.Product, error)
	GetById(id int) (model.Product, error)
	Update(product model.Product) error
	Delete(id int) error
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Insert(product model.Product) (model.Product, error) {
	model.Products = append(model.Products, product)
	return product, nil
}

func (r *repository) GetAll() ([]model.Product, error) {
	return model.Products, nil
}

func (r *repository) GetById(id int) (model.Product, error) {
	p := model.Products
	for _, v := range p {
		if v.Id == id {
			return v, nil
		}
	}
	return model.Product{}, errors.New("id não encontrado")
}

func (r *repository) Update(product model.Product) error {
	p := model.Products
	for i, v := range p {
		if v.Id == product.Id {
			p[i] = product
			return nil
		}
	}
	return errors.New("id não encontrado")

}

func (r *repository) Delete(id int) error {
	p := model.Products
	for i, v := range p {
		if v.Id == id {
			model.Products = append(p[:i], p[i+1:]...)
			return nil
		}
	}
	return errors.New("id não encontrado")
}

func (r *repository) LastId() (int, error) {
	ts := model.Products
	maxId := ts[0].Id
	for i := 1; i <= len(ts)-1; i++ {
		if ts[i].Id > maxId {
			maxId = ts[i].Id
		}
	}
	return maxId + 1, nil
}
