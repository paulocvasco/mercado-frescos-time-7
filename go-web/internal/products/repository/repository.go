package repository

import "mercado-frescos-time-7/go-web/internal/products/model"

type Repository interface {
	Insert(product model.Product) (model.Product, error)
	GetAll() ([]model.Product, error)
	GetById(id int) (model.Product, error)
	Update(id int, product model.Product) (model.Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (s *repository) Insert(product model.Product) (model.Product, error) {

}

func (s *repository) GetAll() ([]model.Product, error) {

}

func (s *repository) GetById(id int) (model.Product, error) {

}

func (s *repository) Update(id int, product model.Product) (model.Product, error) {

}

func (s *repository) Delete(id int) error {

}
