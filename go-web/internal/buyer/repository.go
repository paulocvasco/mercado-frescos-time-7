package buyer

import (
	"fmt"
	"log"
	b "mercado-frescos-time-7/go-web/internal/models"
)

type RequestPost struct {
	CardNumberID int    `json:"card_number_id,omitempty" `
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
	GetAll() ([]b.Buyer, error)
	GetId(id int) (b.Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	Update(id int, body b.Buyer) (b.Buyer, error)
	Delete(id int) error
}

var db []b.Buyer = []b.Buyer{}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]b.Buyer, error) {
	return db, nil
}

func (r *repository) GetId(id int) (b.Buyer, error) {
	var getById b.Buyer
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

func (r *repository) Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error) {
	newBuyer := b.Buyer{id, card_number_id, first_name, last_name}
	db = append(db, newBuyer)
	return newBuyer, nil
}

func (r *repository) Update(id int, body b.Buyer) (b.Buyer, error) {
	var returnDB b.Buyer
	update := false

	log.Println(body)

	for i := range db {
		if db[i].ID == id {
			db[i] = b.Buyer(body)
			returnDB = b.Buyer(body)
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
