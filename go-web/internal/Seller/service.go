package seller

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	GetAll() ([]models.Seller, error)
	GetId(indice int) (models.Seller, error)
	Update(s []byte, id int) (models.Seller, error)
	Delete(id int) error
	Store(cid int, company_name string, address string, telephone string) (models.Seller, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]models.Seller, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetId(indice int) (models.Seller, error) {
	ps, err := s.repository.GetId(indice)
	if err != nil {
		return models.Seller{}, err
	}
	return ps, nil
}

func (s *service) Update(sel []byte, id int) (models.Seller, error) {
	oldSeller, err := s.repository.GetId(id)
	if err != nil {
		return models.Seller{}, err
	}

	var sc Seller
	err = json.Unmarshal(sel, &sc)
	if err != nil {
		return models.Seller{}, err
	}
	_, err = s.repository.CheckCid(sc.Cid)
	if err != nil {
		return models.Seller{}, customerrors.ErrorConflict
	}

	oldSellerJSON, _ := json.Marshal(oldSeller)
	patch, err := jsonpatch.MergePatch(oldSellerJSON, sel)
	var updatedSeller Seller
	if err != nil {
		return models.Seller{}, err
	}
	err = json.Unmarshal(patch, &updatedSeller)
	if err != nil {
		return models.Seller{}, err
	}
	ps, err := s.repository.Update(updatedSeller, id)
	if err != nil {
		return models.Seller{}, err
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

func (s *service) Store(cid int, company_name string, address string, telephone string) (models.Seller, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return models.Seller{}, err
	}
	lastID++
	product, err := s.repository.Store(lastID, cid, company_name, address, telephone)
	if err != nil {
		return models.Seller{}, err
	}
	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
