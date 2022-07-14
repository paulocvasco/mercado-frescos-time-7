package products

import (
	"encoding/json"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Service interface {
	Insert(newProduct []byte) (models.Product, error)
	GetAll() (models.Products, error)
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
	product := models.Product{}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return models.Product{}, err
	}
	productJSON, err = jsonpatch.MergePatch(productJSON, newProduct)
	if err != nil {
		return models.Product{}, err
	}

	json.Unmarshal(productJSON, &product)

	if products, err := s.repository.GetAll(); err != nil {
		return models.Product{}, err
	} else {
		for _, v := range products.Products {
			if v.ProductCode == product.ProductCode {
				return models.Product{}, customerrors.ErrorConflict
			}
		}
	}

	p, err := s.repository.Insert(product)
	if err != nil {
		return models.Product{}, err
	}
	return p, nil
}

func (s *service) GetAll() (models.Products, error) {
	pp, err := s.repository.GetAll()
	if err != nil {
		return models.Products{}, err
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

	if oldProduct.ProductCode != updateProduct.ProductCode {
		if products, err := s.repository.GetAll(); err != nil {
			return models.Product{}, err
		} else {
			for _, v := range products.Products {
				if v.ProductCode == updateProduct.ProductCode {
					return models.Product{}, customerrors.ErrorConflict
				}
			}
		}
	}

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
