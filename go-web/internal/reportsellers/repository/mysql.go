package repository

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type Repository interface {
	ReportSellers(id int) ([]models.ReportSeller, error)
}

type SQLrepository struct {
	db *sql.DB
}

func NewSQLrepository(db *sql.DB) Repository {
	return &SQLrepository{
		db: db,
	}
}

func (r *SQLrepository) ReportSellers(id int) ([]models.ReportSeller, error) {
	var query string
	if id == 0 {
		query = "SELECT l.id, l.locality_name, COUNT(*) FROM sellers s INNER JOIN localities l ON s.locality_id = l.id WHERE s.locality_id > ? GROUP BY s.locality_id;"
	} else {
		query = "SELECT l.id, l.locality_name, COUNT(*) FROM sellers s INNER JOIN localities l ON s.locality_id = l.id WHERE s.locality_id = ? GROUP BY s.locality_id;"
	}
	var reporters []models.ReportSeller
	db := r.db
	rows, err := db.Query(query, id)
	if err != nil {
		log.Println(err)
		return []models.ReportSeller{}, err
	}

	for rows.Next() {
		var rapp models.ReportSeller
		if err := rows.Scan(&rapp.LocalityID, &rapp.Locality_name, &rapp.SellerCount); err != nil {
			log.Println(err)
			return []models.ReportSeller{}, err
		}
		reporters = append(reporters, rapp)
	}

	if len(reporters) == 0 {
		return []models.ReportSeller{}, customerrors.ErrorItemNotFound
	}

	return reporters, nil

}
