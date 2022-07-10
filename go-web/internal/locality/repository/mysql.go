package repository

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type Repository interface {
	/*CheckId(cid string) (models.Locality, error)*/
	Store(loc models.Locality) (models.Locality, error)
}

type SQLrepository struct {
	db *sql.DB
}
//go:generate mockery --name=Repository --output=../mocks --outpkg=mockRepository
func NewSQLrepository(db *sql.DB) Repository {
	return &SQLrepository{
		db: db,
	}
}

func (r *SQLrepository) Store(loc models.Locality) (models.Locality, error) {
	/*_, err := r.CheckId(loc.Id)
	if err != nil {
		return models.Locality{}, err
	}*/
	var exists int
	db1 := r.db
	sqlquery := "SELECT a.id FROM provincies a INNER JOIN countries b ON a.id_country_fk = b.Id AND b.country_name = ? AND a.provincie_name = ?;"
	result := db1.QueryRow(sqlquery, loc.Country_name, loc.Province_name)
	if result.Err() != nil {
		log.Println(result.Err())
		return models.Locality{}, result.Err()
	}

	err := result.Scan(&exists)
	if err != nil {
		log.Println(err)
		return models.Locality{}, customerrors.ErrorConflict
	}
	db := r.db
	stmt, err := db.Prepare("INSERT INTO `localities` (`id`, `locality_name`, `province_id`) VALUES (?, ?, ?)")
	if err != nil {
		return models.Locality{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(loc.Id, loc.Locality_name, exists)
	if err != nil {
		log.Println(err)
		return models.Locality{}, err
	}

	localityResult := models.Locality{Id: loc.Id, Locality_name: loc.Locality_name, Province_name: loc.Province_name, Country_name: loc.Country_name}
	return localityResult, nil
}

/*func (r *SQLrepository) CheckId(cid string) (models.Locality, error) {
	seller := models.Locality{}
	var exists int
	db := db.StorageDB
	result := db.QueryRow("SELECT exists (SELECT * FROM localities WHERE id = ?)", cid)
	if result.Err() != nil {
		return seller, result.Err()
	}
	err := result.Scan(&exists)
	if err != nil {
		return seller, result.Err()
	}
	if exists == 1 {
		return seller, customerrors.ErrorConflict
	}
	return seller, nil

}

func (r *SQLrepository) LastID() (int, error) {
	return 0, nil
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
}*/
