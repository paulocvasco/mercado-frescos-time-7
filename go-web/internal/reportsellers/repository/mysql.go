package repository

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/pkg/db"
	"strconv"
)

type Repository interface {
    ReportSellers(id int) (models.ReportSeller, error)     
}

type SQLrepository struct {
    db *sql.DB
}

func NewSQLrepository(db *sql.DB) Repository {
	return &SQLrepository{
		db: db,
	}
}

func (r *SQLrepository) ReportSellers(id int) (models.ReportSeller, error) {
    var sellers models.ReportSeller
    db := db.StorageDB
    rows := db.QueryRow("SELECT COUNT(id) FROM `sellers` WHERE `locality_id` = ?", id)
    if rows.Err() != nil {
        log.Println(rows.Err())
        return models.ReportSeller{}, rows.Err()
    }
    var contagem string
    err := rows.Scan(&contagem)
    if err != nil {
        log.Println(err)
        return models.ReportSeller{}, err
    }

    rows = db.QueryRow("SELECT locality_name FROM `localities` WHERE id = ?", id)
    if rows.Err() != nil {
        log.Println(rows.Err())
        return models.ReportSeller{}, rows.Err()
    }
    var nome string
    err = rows.Scan(&nome)
    if err != nil {
        log.Println(err)
        return models.ReportSeller{}, err
    }

    sellers.LocalityID = strconv.Itoa(id)
    sellers.SellerCount = contagem
    sellers.Locality_name = nome
    return sellers, nil
}