package buyer

import (
	"encoding/json"
	model "mercado-frescos-time-7/go-web/internal/models"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	GetAll() []model.Buyer
	GetId(id int) (model.Buyer, error)
	Create(card_number_id int, first_name, last_name string) (model.Buyer, error)
	Update(id int, body RequestPatch) (model.Buyer, error)
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

func (s *service) GetAll() []model.Buyer {
	response := s.repository.GetAll()
	return response
}
func (s *service) GetId(id int) (model.Buyer, error) {
	response, err := s.repository.GetId(id)
	if err != nil {
		return model.Buyer{}, err
	}
	return response, nil
}

func (s *service) Create(cardNumberID int, firstName, lastName string) (model.Buyer, error) {

	response, err := s.repository.Create(cardNumberID, firstName, lastName)
	if err != nil {
		return model.Buyer{}, err
	}
	return response, nil
}

func (s *service) Update(id int, newData RequestPatch) (model.Buyer, error) {
	getById, err := s.repository.GetId(id)
	var emptyBuyer model.Buyer

	if err != nil {
		return emptyBuyer, err
	}
	buyerByte, err := json.Marshal(getById)
	if err != nil {
		return emptyBuyer, err
	}
	newDataByte, err := json.Marshal(newData)
	if err != nil {
		return emptyBuyer, err
	}

	buyerPatch, err := jsonpatch.MergeMergePatches(buyerByte, newDataByte)

	if err != nil {
		return emptyBuyer, err
	}

	err = json.Unmarshal(buyerPatch, &getById)
	if err != nil {
		return emptyBuyer, err
	}
	newUpdate, err := s.repository.Update(id, getById)
	if err != nil {
		return emptyBuyer, err
	}
	return newUpdate, nil
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
