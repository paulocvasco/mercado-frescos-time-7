package products

import (
	"errors"
)

type Repository interface {
	Insert(product Product) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Update(product Product) error
	Delete(id int) error
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Insert(product Product) (Product, error) {
	Products = append(Products, product)
	return product, nil
}

func (r *repository) GetAll() ([]Product, error) {
	return Products, nil
}

func (r *repository) GetById(id int) (Product, error) {
	p := Products
	for _, v := range p {
		if v.Id == id {
			return v, nil
		}
	}
	return Product{}, errors.New("id não encontrado")
}

func (r *repository) Update(product Product) error {
	p := Products
	for i, v := range p {
		if v.Id == product.Id {
			p[i] = product
			return nil
		}
	}
	return errors.New("id não encontrado")

}

func (r *repository) Delete(id int) error {
	p := Products
	for i, v := range p {
		if v.Id == id {
			Products = append(p[:i], p[i+1:]...)
			return nil
		}
	}
	return errors.New("id não encontrado")
}

func (r *repository) LastId() (int, error) {
	ts := Products
	maxId := ts[0].Id
	for i := 1; i <= len(ts)-1; i++ {
		if ts[i].Id > maxId {
			maxId = ts[i].Id
		}
	}
	return maxId + 1, nil
}
