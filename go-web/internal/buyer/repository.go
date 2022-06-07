package buyer

import (
	"fmt"
	model "mercado-frescos-time-7/go-web/internal/models"
)

type RequestPost struct {
	CardNumberID *int   `json:"card_number_id,omitempty" `
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
}
type Request struct {
	ID           int    `json:"id" binding:"required"`
	CardNumberID int    `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

type Repository interface {
	GetAll() []model.Buyer
	GetId(id int) (model.Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (model.Buyer, error)
	Update(id int, body model.Buyer) (model.Buyer, error)
	Delete(id int) error
}

var db []model.Buyer = []model.Buyer{}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() []model.Buyer {
	return db
}

func (r *repository) GetId(id int) (model.Buyer, error) {
	var getById model.Buyer
	get := false
	for i := range db {
		if db[i].ID == id {
			getById = db[i]
			get = true
		}
	}
	if !get {
		return getById, fmt.Errorf("product %d não encontrado", id)
	}
	return getById, nil
}

func (r *repository) Creat(id, card_number_id int, first_name, last_name string) (model.Buyer, error) {
	newBuyer := model.Buyer{ID: id, CardNumberID: card_number_id, FirstName: first_name, LastName: last_name}
	db = append(db, newBuyer)
	return newBuyer, nil
}

func (r *repository) Update(id int, body model.Buyer) (model.Buyer, error) {
	var returnDB model.Buyer
	update := false

	for i := range db {
		if db[i].ID == id {
			db[i] = model.Buyer(body)
			returnDB = model.Buyer(body)
			update = true
		}
	}

	if !update {
		return returnDB, fmt.Errorf("product %d não encontrado", id)
	}
	return returnDB, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	for i := range db {
		if db[i].ID == id {
			db = append(db[:i], db[i+1:]...)
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("product %d não encontrado", id)
	}
	return nil

}
