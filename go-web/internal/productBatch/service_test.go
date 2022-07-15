package productBatch_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mercado-frescos-time-7/go-web/internal/productBatch"
	"mercado-frescos-time-7/go-web/internal/productBatch/domain"
	"mercado-frescos-time-7/go-web/internal/productBatch/mock/mockRepository"
	"testing"
)

func TestCreateOK(t *testing.T) {
	repo := mockRepository.NewProductBatchRepository(t)
	serv := productBatch.NewService(repo)

	productBatch := &domain.ProductBatch{
		BatchNumber:        1013,
		CurrentQuantity:    100,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    1,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		ProductId:          1,
		SectionId:          1,
	}

	repo.On("CreateProductBatch", mock.Anything, mock.Anything).Return(productBatch, nil).Maybe()
	res, err := serv.Store(context.TODO(), productBatch)

	assert.Equal(t, productBatch, res)
	assert.NoError(t, err)
}

func TestCreateError(t *testing.T) {
	repo := mockRepository.NewProductBatchRepository(t)
	serv := productBatch.NewService(repo)

	productBatch := &domain.ProductBatch{
		BatchNumber:        1013,
		CurrentQuantity:    100,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    1,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		ProductId:          1,
		SectionId:          1,
	}

	repo.On("CreateProductBatch", mock.Anything, mock.Anything).Return(&domain.ProductBatch{}, sqlmock.ErrCancelled).Maybe()
	res, err := serv.Store(context.TODO(), productBatch)

	assert.Equal(t, &domain.ProductBatch{}, res)
	assert.Equal(t, err, sqlmock.ErrCancelled)
}
