package buyer

import (
	"fmt"
	"log"
	b "mercado-frescos-time-7/go-web/internal/models"
)

type Repository interface {
	GetAll() ([]b.Buyer, error)
	GetId(id int) (b.Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	// Update(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	// Delete(id int) error
}

var db []b.Buyer = []b.Buyer{}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]b.Buyer, error) {
	log.Print(db)
	return db, nil
}

func (r *repository) GetId(id int) (b.Buyer, error) {
	log.Println("cheguei GetId", id)
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
	log.Println(db)
	newBuyer := b.Buyer{id, card_number_id, first_name, last_name}
	db = append(db, newBuyer)

	log.Println(db)

	return newBuyer, nil

}
