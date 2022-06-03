package buyer

type Service interface {
	GetAll() ([]Buyer, error)
	GetId(id int) (Buyer, error)
	Creat(id, card_number_id int, first_name, last_name string) (Buyer, error)
	Update(id, card_number_id int, first_name, last_name string) (Buyer, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Buyer, error) {
	response, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}
