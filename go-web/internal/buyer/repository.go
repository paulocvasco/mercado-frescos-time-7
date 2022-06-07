package buyer

import (
	"fmt"
	"log"
	model "mercado-frescos-time-7/go-web/internal/models"
)

type RequestPatch struct {
	CardNumberID *int   `json:"card_number_id,omitempty" `
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
}
type Repository interface {
	GetAll() []model.Buyer
	GetId(id int) (model.Buyer, error)
	Create(CardNumberID int, FirstName, LastName string) (model.Buyer, error)
	Update(id int, body model.Buyer) (model.Buyer, error)
	Delete(id int) error
	GenerateID() int
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
	for _, value := range db {
		if value.ID == id {
			getById = value
			return getById, nil
		}
	}
	return getById, fmt.Errorf("Error")
}

func (r *repository) Create(cardNumberID int, firstName, lastName string) (model.Buyer, error) {
	newId := r.GenerateID()
	newBuyer := model.Buyer{ID: newId, CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName}
	db = append(db, newBuyer)
	return newBuyer, nil
}

func (r *repository) Update(id int, newData model.Buyer) (model.Buyer, error) {
	var returnDB model.Buyer

	for i, value := range db {
		if value.ID == id {
			db[i] = model.Buyer(newData)
			log.Println(value)
			returnDB = model.Buyer(newData)
			return returnDB, nil
		}
	}
	return returnDB, fmt.Errorf("Error")
}

func (r *repository) Delete(id int) error {
	for i, value := range db {
		if value.ID == id {
			db = append(db[:i], db[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Error")

}

var lastId []int

func (r *repository) GenerateID() int {

	if len(db) == 0 {
		lastId = append(lastId, 1)
		return 1
	}

	result := lastId[len(lastId)-1] + 1
	lastId = append(lastId, result)

	return result
}
