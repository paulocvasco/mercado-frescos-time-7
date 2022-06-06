package buyer

import (
	"encoding/json"
	b "mercado-frescos-time-7/go-web/internal/models"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	GetAll() ([]b.Buyer, error)
	GetId(id int) (b.Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	Update(id int, body RequestPost) (b.Buyer, error)
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

func (s *service) GetAll() ([]b.Buyer, error) {
	response, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *service) GetId(id int) (b.Buyer, error) {
	response, err := s.repository.GetId(id)
	if err != nil {
		return b.Buyer{}, err
	}
	return response, nil
}

func (s *service) Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error) {
	response, err := s.repository.Creat(id, card_number_id, first_name, last_name)
	if err != nil {
		return b.Buyer{}, err
	}
	return response, nil
}

func (s *service) Update(id int, body RequestPost) (b.Buyer, error) {
	getById, err := s.repository.GetId(id)

	if err != nil {
		return getById, err
	}
	buyerMarch, err := json.Marshal(getById)
	bodyMarch, err := json.Marshal(body)

	buyerPatch, err := jsonpatch.MergeMergePatches(buyerMarch, bodyMarch)

	if err != nil {
		return getById, err
	}

	err = json.Unmarshal(buyerPatch, &getById)
	if err != nil {
		return getById, err
	}
	newUpdate, err := s.repository.Update(id, getById)
	if err != nil {
		return getById, err
	}
	return newUpdate, nil
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
