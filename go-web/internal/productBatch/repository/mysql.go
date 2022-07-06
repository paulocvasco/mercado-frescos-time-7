package repository

import (
	"context"
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/productBatch/domain"
	db2 "mercado-frescos-time-7/go-web/pkg/db"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.ProductBatchRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateProductBatch(ctx context.Context, productBatch *domain.ProductBatch) (*domain.ProductBatch, error) {

	db := db2.StorageDB

	stmt, err := db.Prepare(`INSERT INTO products_batches (batch_number, current_quantity, current_tempertature, 
                              due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) 
								 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var result sql.Result

	result, err = stmt.ExecContext(ctx, &productBatch.BatchNumber, &productBatch.CurrentQuantity, &productBatch.CurrentTemperature,
		&productBatch.DueDate, &productBatch.InitialQuantity, &productBatch.ManufacturingDate, &productBatch.ManufacturingHour,
		&productBatch.MinimumTemperature, &productBatch.ProductId, &productBatch.SectionId)

	if err != nil {
		return &domain.ProductBatch{}, err
	}
	insertedId, _ := result.LastInsertId()
	productBatch.Id = int(insertedId)

	return productBatch, nil
}
