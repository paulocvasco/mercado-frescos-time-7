package buyer

type Service interface {
	GetAll() ([]Buyer, error)
	GetId() (Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (Buyer, error)
	Updat(id, card_number_id int, first_name, last_name string) (Buyer, error)
	Delete(id int) error
}
