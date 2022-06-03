package service

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/products/model"
	r "mercado-frescos-time-7/go-web/internal/products/repository"
)

type Service interface {
	Insert(product model.Product) (model.Product, error)
	GetAll() ([]model.Product, error)
	GetById(id int) (model.Product, error)
	Update(id int, product model.Product) error
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
	id, err := s.repository.LastId()
	if err != nil {
		return model.Product{}, errors.New("erro interno tente mais tarde")
	}
	product.Id = id
	p, err := s.repository.Insert(product)
	if err != nil {
		return model.Product{}, errors.New("erro interno tente mais tarde")
	}
	return p, nil
}

func (s *service) GetAll() ([]model.Product, error) {
	pp, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return pp, nil
}

func (s *service) GetById(id int) (model.Product, error) {
	p, err := s.repository.GetById(id)
	if err != nil {
		return model.Product{}, err
	}
	return p, err
}

func (s *service) Update(id int, product model.Product) error {
	oldProduct, err := s.repository.GetById(id)
	if err != nil {
		return err
	}
	product.Id = oldProduct.Id
	err = s.repository.Update(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
