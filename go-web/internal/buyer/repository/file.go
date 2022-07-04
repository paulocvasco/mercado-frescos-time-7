package repository

import (
	model "mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db"
)

var cache model.BuyersMetaData

type RequestPatch struct {
	CardNumberID string `json:"card_number_id,omitempty" `
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
}
type Repository interface {
	GetAll() (model.Buyers, error)
	GetId(id int) (model.Buyer, error)
	GetCardNumberId(id string) error
	Create(CardNumberID string, FirstName, LastName string) (model.Buyer, error)
	Update(id int, body model.Buyer) (model.Buyer, error)
	Delete(id int) error
}
type repository struct {
	data db.DB
}

func NewRepository(data db.DB) Repository {
	return &repository{
		data: data,
	}
}

func (r *repository) GetAll() (model.Buyers, error) {
	if cache.LastID == 0 {
		err := r.data.Load(&cache)
		if err != nil {
			return model.Buyers{}, err
		}
	}
	return cache.Content, nil
}

func (r *repository) GetId(id int) (model.Buyer, error) {

	if cache.LastID == 0 {
		err := r.data.Load(&cache)
		if err != nil {
			return model.Buyer{}, err
		}
	}

	if id < 0 || id > cache.LastID {
		return model.Buyer{}, customerrors.ErrorInvalidID
	}
	var getById model.Buyer
	for _, value := range cache.Content.Buyer {
		if value.ID == id {
			getById = value
			return getById, nil
		}
	}
	return getById, customerrors.ErrorInvalidID
}

func (r *repository) Create(cardNumberID string, firstName, lastName string) (model.Buyer, error) {
	var buyers model.BuyersMetaData
	err := r.data.Load(&buyers)
	if err != nil {
		return model.Buyer{}, err
	}

	newId := buyers.LastID + 1
	buyers.LastID = newId

	newBuyer := model.Buyer{ID: newId, CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName}
	buyers.Content.Buyer = append(buyers.Content.Buyer, newBuyer)
	err = r.data.Save(buyers)
	if err != nil {
		return model.Buyer{}, err
	}
	cache = buyers
	return newBuyer, nil
}

func (r *repository) Update(id int, newData model.Buyer) (model.Buyer, error) {

	if id < 0 || id > cache.LastID {
		return model.Buyer{}, customerrors.ErrorInvalidID
	}
	var returnDB model.Buyer

	for i, value := range cache.Content.Buyer {
		if value.ID == id {
			cache.Content.Buyer[i] = newData
			returnDB = model.Buyer(newData)

		}
	}
	err := r.data.Save(cache)
	if err != nil {
		return model.Buyer{}, err
	}

	return returnDB, nil
}

func (r *repository) Delete(id int) error {
	var buyers model.BuyersMetaData
	err := r.data.Load(&buyers)
	if err != nil {
		return err
	}
	if id < 0 || id > buyers.LastID {
		return customerrors.ErrorInvalidID
	}
	for i, value := range buyers.Content.Buyer {
		if value.ID == id {
			buyers.Content.Buyer = append(buyers.Content.Buyer[:i], buyers.Content.Buyer[i+1:]...)

			err = r.data.Save(buyers)
			if err != nil {
				return err
			}
			cache = buyers
			return nil
		}
	}
	return customerrors.ErrorInvalidID

}

func (r *repository) GetCardNumberId(cardId string) error {
	err := r.data.Load(&cache)
	if err != nil {
		return err
	}
	for _, value := range cache.Content.Buyer {
		if value.CardNumberID == cardId {
			return customerrors.ErrorCardIdAlreadyExists
		}
	}
	return nil
}
