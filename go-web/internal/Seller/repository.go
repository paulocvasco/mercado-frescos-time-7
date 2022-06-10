package seller

import (
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"

	"golang.org/x/exp/slices"
)

type Seller models.Seller

var ps []Seller 
var lastID int

type Repository interface {
	GetAll() ([]Seller, error)
	GetId(indice int) (Seller, error)
	CheckCid(cid int) (Seller, error)
	Update(s Seller, id int) (Seller, error)
	Delete(id int) error
	Store(id int, cid int, company_name string, address string, telephone string) (Seller, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Seller, error) {
	return ps, nil
}

func (r *repository) Delete(id int) error {
	for k, v := range ps {
		if v.ID == id {
			ps = slices.Delete(ps, k, k+1)
			return nil
		}
	}
	return errors.New("id não encontrado")
}

func (r *repository) GetId(indice int) (Seller, error) {
	for _, v := range ps {
		if v.ID == indice {
			return v, nil
		}
	}
	return Seller{}, errors.New("id não encontrado")
}

func (r *repository) CheckCid(cid int) (Seller, error) {
	for _, v := range ps {
		if v.Cid == cid {
			return Seller{}, errors.New("cid já existente")
		}
	}
	return Seller{}, nil
}

func (r *repository) Update(newValues Seller, id int) (Seller, error) {
	_, err := r.CheckCid(newValues.Cid)
	if err != nil {
		return Seller{}, err
	}
	for k, v := range ps {
		if v.ID == id {
			if newValues.Address != "" {
				v.Address = newValues.Address
			}
			if newValues.Cid != 0 {
				v.Cid = newValues.Cid
			}
			if newValues.Company_name != "" {
				v.Company_name = newValues.Company_name
			}
			if newValues.Telephone != "" {
				v.Telephone = newValues.Telephone
			}
			ps[k] = v
			return v, nil
		}
	}
	return Seller{}, errors.New("id não encontrado")
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, cid int, company_name string, address string, telephone string) (Seller, error) {
	_, err := r.CheckCid(cid)
	if err != nil {
		return Seller{}, err
	}
	p := Seller{id, cid, company_name, address, telephone}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}
