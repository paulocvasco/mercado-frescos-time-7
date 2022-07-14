package inbound_orders

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/inbound_orders/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
)

type Service interface {
	Create(order_date string, order_number string, employee_id int, product_batch_id int, warehouse_id int) (models.InboundOrders, error)
}

type service struct {
	repository repository.Repository
}

func (s *service) Create(order_date string, order_number string, employee_id int, product_batch_id int, warehouse_id int) (models.InboundOrders, error) {

	inboudOrders, err := s.repository.Create(order_date, order_number, employee_id, product_batch_id, warehouse_id)

	if err != nil {
		return models.InboundOrders{}, err
	}

	return models.InboundOrders(inboudOrders), nil
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}
