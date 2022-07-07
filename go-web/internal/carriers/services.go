package carriers

import (
	"mercado-frescos-time-7/go-web/internal/carriers/repository"
	"mercado-frescos-time-7/go-web/internal/models"
)

type Service interface {
	Create(models.CarrierRequest) (models.Carrier, error)
	Get(int) (models.CarriersReport, error)
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	s := &service{
		repository: r,
	}
	return s
}

func (s *service) Create(new models.CarrierRequest) (models.Carrier, error) {
	return models.Carrier{}, nil
}

func (s *service) Get(id int) (models.CarriersReport, error) {
	return models.CarriersReport{}, nil
}
