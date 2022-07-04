package repository

import (
	"database/sql"
	"errors"
	model "mercado-frescos-time-7/go-web/internal/models"
)

type RepositoryMysql interface {
	GetAll() ([]model.Buyer, error)
	// GetId(id int) (*model.Buyer, error)
	// //Create(CardNumberID string, FirstName, LastName string) (*model.Buyer, error)
	// Update(id int, body *model.Buyer) (*model.Buyer, error)
	// Delete(id int) error
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
		return []model.Buyer{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var section model.Buyer
		if err := rows.Scan(
			&section.ID,
			&section.CardNumberID,
			&section.FirstName,
			&section.LastName,
		); err != nil {
			return []model.Buyer{}, err
		}
		allBuyers = append(allBuyers, section)
	}
	return allBuyers, nil
}
func (r repositoryDb) GetId(id int64) (model.Buyer, error) {
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
