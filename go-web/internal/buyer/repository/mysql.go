package repository

import (
	"database/sql"
	"errors"
	"fmt"
	model "mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/logger"
)

type RepositoryMysql interface {
	GetAll() ([]model.Buyer, error)
	GetId(id int) (model.Buyer, error)
	Create(CardNumberID string, FirstName, LastName string) (model.Buyer, error)
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
		logger.Save(err.Error())
		return []model.Buyer{}, err
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
			logger.Save(err.Error())
			return []model.Buyer{}, err
		}

		allBuyers = append(allBuyers, section)
	}

	logger.Save(logger.BuyersResquested)
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
		logger.Save(err.Error())
		return model.Buyer{}, customerrors.ErrorInvalidID
	}

	if err != nil {
		logger.Save(err.Error())
		return model.Buyer{}, err
	}

	logger.Save(fmt.Sprintf(logger.BuyerRequested, id))
	return section, nil
}

func (r repositoryDb) Create(CardNumberID string, FirstName, LastName string) (model.Buyer, error) {

	query := `INSERT INTO buyers(id_card_number,first_name,last_name)
	VALUES (?, ?, ?)`

	stmt, _ := r.db.Prepare(query)

	defer stmt.Close()

	section := model.Buyer{
		CardNumberID: CardNumberID,
		FirstName:    FirstName,
		LastName:     LastName,
	}

	result, err := stmt.Exec(
		&section.CardNumberID,
		&section.FirstName,
		&section.LastName,
	)
	if err != nil {
		logger.Save(err.Error())
		return model.Buyer{}, err
	}
	lastId, err := result.LastInsertId()

	if err != nil {
		logger.Save(err.Error())
		return model.Buyer{}, err
	}
	section.ID = int(lastId)
	logger.Save(fmt.Sprintf(logger.BuyerCreated, lastId))
	return section, nil

}

func (r repositoryDb) Delete(id int) error {
	query := `DELETE FROM buyers where id_card_number = ?`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		logger.Save(err.Error())
		return err
	}

	result, _ := stmt.Exec(id)

	affectedRows, err := result.RowsAffected()
	if err != nil {
		logger.Save(err.Error())
		return err
	}

	if affectedRows == 0 {
		logger.Save(err.Error())
		return customerrors.ErrorInvalidID
	}
	logger.Save(fmt.Sprintf(logger.BuyerDeleted, id))
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
		logger.Save(err.Error())
		return model.Buyer{}, err
	}

	logger.Save(fmt.Sprintf(logger.BuyerChanged, id))
	return section, nil

}
