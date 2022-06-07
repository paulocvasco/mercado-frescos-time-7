package buyer

import (
	"encoding/json"
	"fmt"
	model "mercado-frescos-time-7/go-web/internal/models"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	GetAll() ([]model.Buyer, error)
	GetId(id int) (model.Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (model.Buyer, error)
	Update(id int, body RequestPost) (model.Buyer, error)
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

func (s *service) GetAll() ([]model.Buyer, error) {
	response, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *service) GetId(id int) (model.Buyer, error) {
	response, err := s.repository.GetId(id)
	if err != nil {
		return model.Buyer{}, err
	}
	return response, nil
}

func (s *service) Creat(id, card_number_id int, first_name, last_name string) (model.Buyer, error) {

	_, err := s.repository.GetId(id)

	if err == nil {
		return model.Buyer{}, fmt.Errorf("ID:%d já existente", id)

	}

	response, err := s.repository.Creat(id, card_number_id, first_name, last_name)
	if err != nil {
		return model.Buyer{}, err
	}
	return response, nil
}

func (s *service) Update(id int, body RequestPost) (model.Buyer, error) {
	getById, err := s.repository.GetId(id)

	if err != nil {
		return getById, err
	}
	buyerMarch, err := json.Marshal(getById)
	if err != nil {
		return getById, err
	}
	bodyMarch, _ := json.Marshal(body)
	if err != nil {
		return getById, err
	}

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
