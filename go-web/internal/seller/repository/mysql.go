package repository

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/seller"
	"strconv"
)

type SQLrepository struct {
	db *sql.DB
}

func NewSQLrepository(db *sql.DB) seller.Repository {
	return &SQLrepository{
		db: db,
	}
}

func (r *SQLrepository) LocalityStore(id string, locality_name string, province_name string, country_name string) (models.Locality, error) {
	return models.Locality{}, nil
}

func (r *SQLrepository) Store(sel models.Seller) (models.Seller, error) {

	stmt, err := r.db.Prepare("INSERT INTO mercado_db.sellers (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return models.Seller{}, err
	}
	defer stmt.Close()
	var result sql.Result
	localityINT, err := strconv.Atoi(sel.LocalityID)
	if err != nil {
		log.Println(err)
		return models.Seller{}, err
	}
	result, err = stmt.Exec(sel.Cid, sel.Company_name, sel.Address, sel.Telephone, localityINT)
	if err != nil {
		log.Println(err)
		return models.Seller{}, err
	}

	insertedID, _ := result.LastInsertId()
	sellerResult := models.Seller{ID: int(insertedID), Cid: sel.Cid, Company_name: sel.Company_name, Address: sel.Address, Telephone: sel.Telephone, LocalityID: sel.LocalityID}
	return sellerResult, nil
}

func (r *SQLrepository) GetId(indice int) (models.Seller, error) {

	return models.Seller{}, nil
}

func (r *SQLrepository) Update(newValues seller.Seller, id int) (models.Seller, error) {

	return models.Seller{}, nil

}

func (r *SQLrepository) GetAll() ([]models.Seller, error) {

	return []models.Seller{}, nil

}

func (r *SQLrepository) Delete(id int) error {

	return nil
}

func (r *SQLrepository) CheckCid(cid int) (models.Seller, error) {
	return models.Seller{}, nil
}

func (r *SQLrepository) CheckCLocality(cid string) (models.Seller, error) {
	return models.Seller{}, nil
}

func (r *SQLrepository) LastID() (int, error) {
	return 0, nil
}

func (r *SQLrepository) ReportSellers(id int) (models.ReportSeller, error) {
	return models.ReportSeller{}, nil

}
