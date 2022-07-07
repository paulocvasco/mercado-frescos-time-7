package repository

import (
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/models"
)

type Repository interface {
	Insert(record models.ProductRecord) (models.ProductRecord, error)
	GetByProductId(id int) (models.ProductsRecordsResponse, error)
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

func (r *repository) GetByProductId(id int) (models.ProductsRecordsResponse, error) {
	var query string
	if id == 0 {
		query = sqlGetAllRecords
	} else {
		query = sqlGetRecordById
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return models.ProductsRecordsResponse{}, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return models.ProductsRecordsResponse{}, err
	}
	records := models.ProductsRecordsResponse{Records: []models.ProductRecordsResponse{}}
	for rows.Next() {
		record := models.ProductRecordsResponse{}
		err := rows.Scan(&record.ProductId, &record.Description, &record.RecordsCount)
		if err != nil {
			return models.ProductsRecordsResponse{}, err
		}
		records.Records = append(records.Records, record)
	}
	return records, nil
}
