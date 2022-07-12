package repository

import (
	"context"
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/productBatch/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepositoryProductBatch(db *sql.DB) domain.ProductBatchRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateProductBatch(ctx context.Context, productBatch *domain.ProductBatch) (*domain.ProductBatch, error) {

	stmt, err := r.db.Prepare(`INSERT INTO products_batches (batch_number, current_quantity, current_tempertature, 
                              due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) 
								 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return &domain.ProductBatch{}, err
	}

	defer stmt.Close()

	var result sql.Result

	result, err = stmt.ExecContext(ctx, &productBatch.BatchNumber, &productBatch.CurrentQuantity, &productBatch.CurrentTemperature,
		&productBatch.DueDate, &productBatch.InitialQuantity, &productBatch.ManufacturingDate, &productBatch.ManufacturingHour,
		&productBatch.MinimumTemperature, &productBatch.ProductId, &productBatch.SectionId)

	if err != nil {
		return &domain.ProductBatch{}, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return &domain.ProductBatch{}, err
	}
	productBatch.Id = int(insertedId)

	return productBatch, nil
}
