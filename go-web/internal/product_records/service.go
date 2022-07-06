package productrecords

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/product_records/repository"

	jsonpatch "github.com/evanphx/json-patch"
)

type Service interface {
	Insert(record []byte) (models.ProductRecord, error)
	GetByProductId(id int) (models.ProductRecords, error)
}

type service struct {
	repository repository.Repository
}

func NewServiceProductRecord(repo repository.Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Insert(recordBytes []byte) (models.ProductRecord, error) {
	record := models.ProductRecord{}

	recordJSON, err := json.Marshal(record)
	if err != nil {
		return models.ProductRecord{}, err
	}
	recordJSON, err = jsonpatch.MergePatch(recordJSON, recordBytes)
	if err != nil {
		return models.ProductRecord{}, err
	}

	err = json.Unmarshal(recordJSON, &record)
	if err != nil {
		return models.ProductRecord{}, err
	}

	insertedRecord, err := s.repository.Insert(record)
	if err != nil {
		return models.ProductRecord{}, err
	}
	return insertedRecord, nil
}

func (s *service) GetByProductId(id int) (models.ProductRecords, error) {
	records, err := s.repository.GetByProductId(id)
	if err != nil {
		return models.ProductRecords{}, err
	}
	return records, nil
}
