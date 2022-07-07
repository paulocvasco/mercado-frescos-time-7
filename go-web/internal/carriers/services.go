package carriers

import (
	"encoding/json"
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
	rawCarrier, _ := json.Marshal(new)
	var storeCarrier models.Carrier
	json.Unmarshal(rawCarrier, &storeCarrier)

	storedCarrier, err := s.repository.Create(storeCarrier)
	if err != nil {
		return models.Carrier{}, err
	}

	return storedCarrier, nil
}

func (s *service) Get(id int) (models.CarriersReport, error) {
	report, err := s.repository.Get(id)
	if err != nil {
		return models.CarriersReport{}, err
	}
	return report, nil
}
