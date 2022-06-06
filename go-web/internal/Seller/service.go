package seller

type Service interface {
	GetAll() ([]Seller, error)
	Store(cid string, company_name string, address string, telephone string) (Seller, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Seller, error){
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(cid string, company_name string, address string, telephone string) (Seller, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Seller{}, err
	}
	lastID++
	product, err := s.repository.Store(lastID, cid, company_name, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return product, nil
}



func NewService(r Repository) Service  {
	return &service{
		repository: r,
	}
}