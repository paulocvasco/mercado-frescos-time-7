package productBatch

import (
	"context"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/productBatch/domain"
)

type service struct {
	repository domain.ProductBatchRepository
}

func NewService(pbs domain.ProductBatchRepository) domain.ProductBatchService {
	return &service{repository: pbs}
}

func (s *service) Store(ctx context.Context, batch *domain.ProductBatch) (*domain.ProductBatch, error) {
	batch, err := s.repository.CreateProductBatch(ctx, batch)

	if err != nil {
		return &domain.ProductBatch{}, err
	}

	return batch, nil
}
