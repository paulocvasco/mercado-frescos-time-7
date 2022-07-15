package productrecords_test

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	productrecords "mercado-frescos-time-7/go-web/internal/product_records"
	"mercado-frescos-time-7/go-web/internal/product_records/mock/mockRepository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsert(t *testing.T) {

	t.Run("Should return record", func(t *testing.T){
		mockRepo := mockRepository.NewRepository(t)
		service := productrecords.NewServiceProductRecord(mockRepo)
		time := time.Now()
		record := models.ProductRecord{Id: 0, LastUpdateDate: time, PurchasePrice: 10.0, SalePrice: 10.0, ProductId: 1}
		recordBytes, _ := json.Marshal(record)
	
		mockRepo.On("InsertProductRecords", mock.Anything).Return(record, nil)
	
		res, err := service.Insert(recordBytes)
	
		assert.Equal(t, record, res)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return time error", func(t *testing.T){
		mockRepo := mockRepository.NewRepository(t)
		service := productrecords.NewServiceProductRecord(mockRepo)
		time, _ := time.Parse("2006-01-02", "2008-08-08")
		record := models.ProductRecord{Id: 0, LastUpdateDate: time, PurchasePrice: 10.0, SalePrice: 10.0, ProductId: 1}
		recordBytes, _ := json.Marshal(record)
	
		mockRepo.On("InsertProductRecords", mock.Anything).Return(record, nil).Maybe()
	
		res, err := service.Insert(recordBytes)
	
		assert.Equal(t, models.ProductRecord{}, res)
		assert.Equal(t, customerrors.ErrorInvalidDate, err)
	})

	t.Run("Should return invalid db", func(t *testing.T){
		mockRepo := mockRepository.NewRepository(t)
		service := productrecords.NewServiceProductRecord(mockRepo)
		time := time.Now()
		record := models.ProductRecord{Id: 0, LastUpdateDate: time, PurchasePrice: 10.0, SalePrice: 10.0, ProductId: 1}
		recordBytes, _ := json.Marshal(record)
	
		mockRepo.On("InsertProductRecords", mock.Anything).Return(models.ProductRecord{}, customerrors.ErrorInvalidDB)
	
		res, err := service.Insert(recordBytes)
	
		assert.Equal(t, models.ProductRecord{}, res)
		assert.Equal(t, customerrors.ErrorInvalidDB, err)
	})
}

func TestGetProductRecords(t *testing.T) {

	t.Run("Should return records", func(t *testing.T){
		mockRepo := mockRepository.NewRepository(t)
		service := productrecords.NewServiceProductRecord(mockRepo)
		records := models.ProductsRecordsResponse{Records: []models.ProductRecordsResponse{
			{ProductId: 1, Description: "foo", RecordsCount: 5},
		}}
	
		mockRepo.On("GetProductRecords", mock.Anything).Return(records, nil).Maybe()
	
		res, err := service.GetProductRecords(1)
	
		assert.Equal(t, records, res)
		assert.Equal(t, nil, err)
	})
	t.Run("Should return invalid id", func(t *testing.T){
		mockRepo := mockRepository.NewRepository(t)
		service := productrecords.NewServiceProductRecord(mockRepo)
		records := models.ProductsRecordsResponse{}
	
		mockRepo.On("GetProductRecords", mock.Anything).Return(records, nil).Maybe()
	
		res, err := service.GetProductRecords(1)
	
		assert.Equal(t, records, res)
		assert.Equal(t, customerrors.ErrorInvalidID, err)
	})
	t.Run("Should return records", func(t *testing.T){
		mockRepo := mockRepository.NewRepository(t)
		service := productrecords.NewServiceProductRecord(mockRepo)
		records := models.ProductsRecordsResponse{Records: []models.ProductRecordsResponse{}}
	
		mockRepo.On("GetProductRecords", mock.Anything).Return(records, nil).Maybe()
	
		res, err := service.GetProductRecords(0)
	
		assert.Equal(t, records, res)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return invalid db", func(t *testing.T){
		mockRepo := mockRepository.NewRepository(t)
		service := productrecords.NewServiceProductRecord(mockRepo)
		records := models.ProductsRecordsResponse{}
	
		mockRepo.On("GetProductRecords", mock.Anything).Return(records, customerrors.ErrorInvalidDB).Maybe()
	
		res, err := service.GetProductRecords(0)
	
		assert.Equal(t, records, res)
		assert.Equal(t, customerrors.ErrorInvalidDB, err)
	})
}