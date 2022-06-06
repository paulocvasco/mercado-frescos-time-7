package buyer

import (
	"fmt"
	b "mercado-frescos-time-7/go-web/internal/models"
)

type Repository interface {
	GetAll() ([]b.Buyer, error)
	GetId(id int) (b.Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	Update(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	// Delete(id int) error
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

func (r *repository) Update(id, card_number_id int, first_name, last_name string) (b.Buyer, error) {
	var returnDB b.Buyer
	newBuyer := b.Buyer{id, card_number_id, first_name, last_name}
	update := false
	for i := range db {
		if db[i].ID == id {
			db[i] = newBuyer
			update = true
			returnDB = db[i]
		}
	}
	if !update {
		return returnDB, fmt.Errorf("product %d não encontrado", id)
	}
	return returnDB, nil
}
