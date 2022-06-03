package services

import (
	"mercado-frescos-time-7/go-web/internal/products/model"
	r "mercado-frescos-time-7/go-web/internal/products/repository"
)

type Service interface {
	Insert(product model.Product) (model.Product, error)
	GetAll() ([]model.Product, error)
	GetById(id int) (model.Product, error)
	Update(id int, product model.Product) (model.Product, error)
	Delete(id int) error
}

type service struct {
	repository r.Repository
}

func NewService(r r.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Insert(product model.Product) (model.Product, error) {

}

func (s *service) GetAll() ([]model.Product, error) {

}

func (s *service) GetById(id int) (model.Product, error) {

}

func (s *service) Update(id int, product model.Product) (model.Product, error) {

}

func (s *service) Delete(id int) error {

}
