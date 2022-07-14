package locality

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/locality/repository"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
)

type Service interface {
	Store(loc models.Locality) (models.Locality, error)
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Store(loc models.Locality) (models.Locality, error) {
	product, err := s.repository.Store(loc)
	if err != nil {
		return models.Locality{}, err
	}
	return product, nil
}
