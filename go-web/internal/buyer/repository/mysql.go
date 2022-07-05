package repository

import (
	"database/sql"
	"errors"
	"log"
	model "mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type RepositoryMysql interface {
	GetAll() ([]model.Buyer, error)
	GetId(id int) (model.Buyer, error)
	Create(CardNumberID string, FirstName, LastName string) (model.Buyer, error)
	GetCardNumberId(id string) error
	Update(id int, body model.Buyer) (model.Buyer, error)
	Delete(id int) error
}

type repositoryDb struct {
	db *sql.DB
}

func NewRepositoryMySql(db *sql.DB) RepositoryMysql {
	return repositoryDb{db: db}
}

func (r repositoryDb) GetAll() ([]model.Buyer, error) {
	var allBuyers []model.Buyer
	rows, err := r.db.Query("SELECT * FROM buyers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var section model.Buyer
		err := rows.Scan(
			&section.ID,
			&section.CardNumberID,
			&section.FirstName,
			&section.LastName,
		)
		if err != nil {
			return []model.Buyer{}, err
		}

		allBuyers = append(allBuyers, section)
	}
	return allBuyers, nil
}
func (r repositoryDb) GetId(id int) (model.Buyer, error) {
	var section model.Buyer
	rows := r.db.QueryRow("SELECT * FROM buyers WHERE ID = ? ", id)
	err := rows.Scan(
		&section.ID,
		&section.CardNumberID,
		&section.FirstName,
		&section.LastName,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Buyer{}, errors.New("id Not exist")
	}

	if err != nil {
		return model.Buyer{}, err
	}
	return section, nil
}

func (r repositoryDb) Create(CardNumberID string, FirstName, LastName string) (model.Buyer, error) {

	query := `INSERT INTO buyers(id_card_number,first_name,last_name) 
	VALUES (?, ?, ?)	`

	stmt, _ := r.db.Prepare(query)

	defer stmt.Close()

	section := model.Buyer{
		CardNumberID: CardNumberID,
		FirstName:    FirstName,
		LastName:     LastName,
	}
	log.Println(section)

	result, err := stmt.Exec(
		&section.CardNumberID,
		&section.FirstName,
		&section.LastName,
	)
	if err != nil {
		return model.Buyer{}, err
	}
	lastId, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	section.ID = int(lastId)
	return section, nil

}
func (r repositoryDb) GetCardNumberId(cardId string) error {

	var section model.Buyer
	rows := r.db.QueryRow("SELECT * FROM buyers where id_card_number = ? ", cardId)
	err := rows.Scan(
		&section.ID,
		&section.CardNumberID,
		&section.FirstName,
		&section.LastName,
	)

	if err == nil {
		return customerrors.ErrorConflict
	}

	return nil
}
func (r repositoryDb) Delete(id int) error {
	query := `DELETE FROM buyers where id_card_number = ?`

	stmt, _ := r.db.Prepare(query)

	result, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return customerrors.ErrorInvalidID
	}
	if err != nil {
		return err
	}

	return nil
}

func (r repositoryDb) Update(id int, body model.Buyer) (model.Buyer, error) {

	query := `UPDATE buyers SET id_card_number = ?, first_name = ?, last_name = ? WHERE ID = ?`

	stmt, _ := r.db.Prepare(query)

	section := model.Buyer{
		ID:           id,
		CardNumberID: body.CardNumberID,
		FirstName:    body.FirstName,
		LastName:     body.LastName,
	}

	defer stmt.Close()
	_, err := stmt.Exec(
		&section.CardNumberID,
		&section.FirstName,
		&section.LastName,
		id,
	)
	if err != nil {
		return model.Buyer{}, err
	}

	if err != nil {
		log.Fatal(err)
	}
	return section, nil

}
