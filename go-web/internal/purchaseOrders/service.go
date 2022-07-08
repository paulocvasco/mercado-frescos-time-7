package purchaseOrders

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
)

type Service interface {
	Create(data models.PurchaseOrders) (repository.ResultPost, error)
}

type service struct {
	repository repository.RepositoryMysql
}

func NewService(r repository.RepositoryMysql) Service {
	return &service{
		repository: r,
	}
}

func (s service) Create(data models.PurchaseOrders) (repository.ResultPost, error) {
	response, err := s.repository.Create(data)
	if err != nil {
		return response, err
	}
	return response, nil
}
