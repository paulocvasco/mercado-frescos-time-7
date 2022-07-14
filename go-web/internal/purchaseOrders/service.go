package purchaseOrders

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/purchaseOrders/repository"
)

type Service interface {
	Create(data models.PurchaseOrders) (repository.ResultPost, error)
	GetPurchaseOrder(id int) ([]models.ResponsePurchaseByBuyer, error)
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

func (s *service) GetPurchaseOrder(id int) ([]models.ResponsePurchaseByBuyer, error) {
	if id == 0 {
		response, err := s.repository.GetAllPurchaseOrder()
		if err != nil {
			return []models.ResponsePurchaseByBuyer{}, err
		}
		return response, nil
	}
	response, err := s.repository.GetIdPurchaseOrder(id)
	if err != nil {
		return []models.ResponsePurchaseByBuyer{}, err
	}
	return response, nil
}
