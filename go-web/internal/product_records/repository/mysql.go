package repository

import (
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/models"
)

type Repository interface {
	Insert(record models.ProductRecord) (models.ProductRecord, error)
	GetByProductId(id int) (models.ProductRecords, error)
}

type repository struct {
	db *sql.DB
}

func NewRepositoryProductRecord(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(record models.ProductRecord) (models.ProductRecord, error) {
	stmt, err := r.db.Prepare(sqlStoreRecord)
	if err != nil {
		return models.ProductRecord{}, err
	}
	res, err := stmt.Exec(&record.LastUpdateDate, &record.PurchasePrince, &record.SalePrice, &record.ProductId)
	if err != nil {
		return models.ProductRecord{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return models.ProductRecord{}, err
	}
	record.Id = int(lastId)

	return record, nil
}

func (r *repository) GetByProductId(id int) (models.ProductRecords, error) {
	stmt, err := r.db.Prepare(sqlGetRecordById)
	if err != nil {
		return models.ProductRecords{}, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return models.ProductRecords{}, err
	}
	records := models.ProductRecords{Records: []models.ProductRecord{}}
	for rows.Next() {
		record := models.ProductRecord{}
		err := rows.Scan(&record.Id, &record.LastUpdateDate, &record.PurchasePrince, &record.SalePrice, &record.ProductId)
		if err != nil {
			return models.ProductRecords{}, err
		}
		records.Records = append(records.Records, record)
	}
	return records, nil
}
