package buyer

import b "mercado-frescos-time-7/go-web/internal/models"

type Repository interface {
	GetAll() ([]b.Buyers, error)
	GetId(id int) (b.Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	Update(id, card_number_id int, first_name, last_name string) (b.Buyer, error)
	Delete(id int) error
}

type repository struct{}

var db []b.Buyers = []b.Buyers{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]b.Buyers, error) {
	return db, nil
}
