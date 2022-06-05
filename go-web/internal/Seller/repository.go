package seller

import ("errors"
"golang.org/x/exp/slices"
"mercado-frescos-time-7/go-web/internal/Seller/models"
)

type Seller models.Seller
var ps []Seller
var lastID int

type Repository interface {
	GetAll() ([]Seller, error)
	GetId(indice int) (Seller, error)
	Update(s Seller, id int) (Seller, error)
	Delete(id int) (error)
	Store(id int, cid string, company_name string, address string, telephone string) (Seller, error)
	LastID() (int, error)
	}

type repository struct {}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Seller, error) {
	return ps, nil
}

func (r *repository) Delete(id int) (error) {
	for k, v := range ps{
		if v.ID == id {
			ps = slices.Delete(ps, k, k+1)
			return nil
		}
	}
	return errors.New("id não encontrado")
}

func (r *repository) GetId(indice int) (Seller, error) {
	for	_, v := range ps{
		if v.ID == indice {
			return v, nil
		}
	}
	return Seller{}, errors.New("id não encontrado")
}

func (r *repository) Update(s Seller, id int) (Seller, error) {
	for	k, v := range ps{
		if v.ID == id {
			v.Cid = s.Cid
			v.Company_name = s.Company_name
			v.Address = s.Address
			v.Telephone = s.Telephone
			ps[k] = v
			return v, nil
		}
	}
	return Seller{}, errors.New("id não encontrado")
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, cid string, company_name string, address string, telephone string) (Seller, error) {
	p := Seller{id, cid, company_name, address, telephone}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}