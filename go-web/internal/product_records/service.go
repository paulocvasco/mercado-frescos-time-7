package productrecords

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/product_records/repository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
)

type Service interface {
	Insert(record []byte) (models.ProductRecord, error)
	GetProductRecords(id int) (models.ProductsRecordsResponse, error)
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
	if record.LastUpdateDate.Before(time.Now().AddDate(0, 0, -1)) {
		return models.ProductRecord{}, customerrors.ErrorInvalidDate
	}

	insertedRecord, err := s.repository.InsertProductRecords(record)
	if err != nil {
		return models.ProductRecord{}, err
	}
	return insertedRecord, nil
}

func (s *service) GetProductRecords(id int) (models.ProductsRecordsResponse, error) {
	records, err := s.repository.GetProductRecords(id)
	if err != nil {
		return models.ProductsRecordsResponse{}, err
	}
	if id != 0 && len(records.Records) == 0 {
		return models.ProductsRecordsResponse{}, customerrors.ErrorInvalidID
	}
	return records, nil
}
