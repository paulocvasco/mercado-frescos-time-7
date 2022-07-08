package repository

import (
	"database/sql"
	"log"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/seller"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db"
	"strconv"
)

type SQLrepository struct {
    db *sql.DB
}

func NewSQLrepository (db *sql.DB) seller.Repository {
	return &SQLrepository {
		db: db,
	}
}

func (r *SQLrepository) LocalityStore(id string, locality_name string, province_name string, country_name string) (models.Locality, error) {
	_, err := r.CheckCLocality(id)
	if err != nil {
		return models.Locality{}, err
	}
	var exists int
	db1 := db.StorageDB
	result := db1.QueryRow("SELECT id FROM `provincies` WHERE `provincie_name` LIKE ?;", province_name)
	if result.Err() != nil {
		log.Println(result.Err())
		return models.Locality{}, result.Err()
	}

	err = result.Scan(&exists)
	if err != nil {
		log.Println(err)
		return models.Locality{}, customerrors.ErrorConflict
	}

	db := db.StorageDB
	stmt, err := db.Prepare("INSERT INTO `localities` (`id`, `locality_name`, `province_id`) VALUES (?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, locality_name, exists)
	if err != nil {
		log.Println(err)
		return models.Locality{}, err
	}

	localityResult := models.Locality{Id: id, Locality_name: locality_name, Province_name: province_name, Country_name: country_name}
	return localityResult, nil
}


func (r *SQLrepository) Store(sel models.Seller) (models.Seller, error) {

	log.Println("preparando consulta")
	stmt, err := r.db.Prepare("INSERT INTO mercado_db.sellers (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return models.Seller{}, err
	}
	defer stmt.Close()
	log.Println("iniciando consulta")
	var result sql.Result
	log.Println(sel)
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
	var seller models.Seller
	db := db.StorageDB
	rows, err := db.Query("SELECT * FROM `sellers` WHERE `id` = ?", indice)
	if err != nil {
		log.Fatal(err)
		return seller, err
	}
	for rows.Next() {
		if err := rows.Scan(&seller.ID, &seller.Cid, &seller.Company_name, &seller.Address, &seller.Telephone, &seller.LocalityID); err != nil {
			log.Println(err.Error())
			return seller, err
		}
	}
	if seller.ID != indice {
		return seller, customerrors.ErrorInvalidID
	}
	return seller, nil

}

func (r *SQLrepository) Update(newValues seller.Seller, id int) (models.Seller, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("UPDATE `sellers` SET `cid` = ?, `company_name` = ?, `address` = ?, `telephone` = ? WHERE `seller`.`id` = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(newValues.Cid, newValues.Company_name, newValues.Address, newValues.Telephone, id)
	if err != nil {
		return models.Seller{}, err
	}
	return models.Seller{ID: id, Cid: newValues.Cid, Company_name: newValues.Company_name, Address: newValues.Address, Telephone: newValues.Telephone}, nil
}

func (r *SQLrepository) GetAll() ([]models.Seller, error) {
	var sellers []models.Seller
	db := db.StorageDB
	rows, err := db.Query("SELECT * FROM `sellers`")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var seller models.Seller
		if err := rows.Scan(&seller.ID, &seller.Cid, &seller.Company_name, &seller.Address, &seller.Telephone, &seller.LocalityID); err != nil {
			log.Fatal(err)
			return nil, err
		}
		sellers = append(sellers, seller)
	}
	return sellers, nil

}

func (r *SQLrepository) Delete(id int) error {
	db := db.StorageDB
	stmt, err := db.Prepare("DELETE FROM `sellers` WHERE `seller`.`id` = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return customerrors.ErrorInvalidID
	}
	return nil
}

func (r *SQLrepository) CheckCid(cid int) (models.Seller, error) {
	seller := models.Seller{}
	var exists int
	db := db.StorageDB
	result := db.QueryRow("SELECT exists (SELECT * FROM `sellers` WHERE `cid` = ?)", cid)
	log.Println(result)
	if result.Err() != nil {
		return seller, result.Err()
	}
	err := result.Scan(&exists)
	if err != nil {
		return seller, result.Err()
	}
	log.Println(exists)
	if exists == 1 {
		return seller, customerrors.ErrorConflict
	}
	return seller, nil

}

func (r *SQLrepository) CheckCLocality(cid string) (models.Seller, error) {
	seller := models.Seller{}
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
}
