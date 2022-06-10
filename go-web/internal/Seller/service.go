package seller

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	GetAll() ([]Seller, error)
	GetId(indice int) (Seller, error)
	Update(s []byte, id int) (Seller, error)
	Delete(id int) error
	Store(cid int, company_name string, address string, telephone string) (Seller, error)
}

type service struct {
	repository Repository 
}

func (s *service) GetAll() ([]Seller, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetId(indice int) (Seller, error) {
	ps, err := s.repository.GetId(indice)
	if err != nil {
		return Seller{}, err
	}
	return ps, nil
}

func (s *service) Update(sel []byte, id int) (Seller, error) {
	oldSeller, err := s.repository.GetId(id)
	if err != nil {
		return Seller{}, err
	}
	oldSellerJSON, _ := json.Marshal(oldSeller)
	patch, err := jsonpatch.MergePatch(oldSellerJSON, sel)
	if err != nil {
		return Seller{}, err
	}

	var updatedSeller Seller
	err = json.Unmarshal(patch, &updatedSeller)
	if err != nil {
		return Seller{}, err
	}
	ps, err := s.repository.Update(updatedSeller, id)
	if err != nil {
		return Seller{}, err
	}
	return ps, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Store(cid int, company_name string, address string, telephone string) (Seller, error) {
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

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
