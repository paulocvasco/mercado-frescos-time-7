package reportsellers

import (
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/reportsellers/repository"
)

type Service interface {
	ReportSellers(id int) ([]models.ReportSeller, error)
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) ReportSellers(id int) ([]models.ReportSeller, error) {
	product, err := s.repository.ReportSellers(id)
	if err != nil {
		return []models.ReportSeller{}, err
	}
	return product, nil
}
