package products

import (
	"encoding/json"
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	Insert(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetById(id int) (models.Product, error)
	Update(id int, product []byte) (models.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Insert(product models.Product) (models.Product, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return models.Product{}, errors.New("erro interno tente mais tarde")
	}
	product.Id = id
	p, err := s.repository.Insert(product)
	if err != nil {
		return models.Product{}, errors.New("erro interno tente mais tarde")
	}
	return p, nil
}

func (s *service) GetAll() ([]models.Product, error) {
	pp, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return pp, nil
}

func (s *service) GetById(id int) (models.Product, error) {
	p, err := s.repository.GetById(id)
	if err != nil {
		return models.Product{}, err
	}
	return p, err
}

func (s *service) Update(id int, product []byte) (models.Product, error) {
	oldProduct, err := s.repository.GetById(id)
	if err != nil {
		return models.Product{}, err
	}
	oldProductJSON, _ := json.Marshal(oldProduct)
	patch, err := jsonpatch.MergePatch(oldProductJSON, product)
	if err != nil {
		return models.Product{}, err
	}
	var updateProduct models.Product
	json.Unmarshal(patch, &updateProduct)

	err = s.repository.Update(updateProduct)
	if err != nil {
		return models.Product{}, err
	}
	return updateProduct, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
