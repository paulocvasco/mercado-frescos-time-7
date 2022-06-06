package buyer

import b "mercado-frescos-time-7/go-web/internal/models"

type Service interface {
	GetAll() ([]b.Buyer, error)
	GetId(id int) (b.Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	Update(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]b.Buyer, error) {
	response, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s service) GetId(id int) (b.Buyer, error) {
	response, err := s.repository.GetId(id)
	if err != nil {
		return b.Buyer{}, err
	}
	return response, nil
}

func (s service) Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error) {
	response, err := s.repository.Creat(id, card_number_id, first_name, last_name)
	if err != nil {
		return b.Buyer{}, err
	}
	return response, nil
}

func (s service) Update(id, card_number_id int, first_name, last_name string) (b.Buyer, error) {
	response, err := s.repository.Update(id, card_number_id, first_name, last_name)
	if err != nil {
		return b.Buyer{}, err
	}
	return response, nil
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
