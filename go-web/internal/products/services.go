package products

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	Insert(newProduct []byte) (models.Product, error)
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

func (s *service) Insert(newProduct []byte) (models.Product, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return models.Product{}, err
	}

	product := models.Product{}
	product.Id = id

	productJSON, err := json.Marshal(product)
	if err != nil {
		return models.Product{}, err
	}

	productJSON, err = jsonpatch.MergePatch(productJSON, newProduct)
	if err != nil {
		return models.Product{}, err
	}

	json.Unmarshal(productJSON, &product)

	p, err := s.repository.Insert(product)
	if err != nil {
		return models.Product{}, err
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
	return p, nil
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
