package repository

import (
	"database/sql"
	"fmt"
	"github.com/paulocvasco/mercado-frescos-time-7/go-web/internal/models"
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type Repository interface {
	Create(models.Carrier) (models.Carrier, error)
	Get(int) (models.CarriersReport, error)
}

type mysqlDB struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &mysqlDB{
		db: db,
	}
}

func (m *mysqlDB) Create(new models.Carrier) (models.Carrier, error) {
	query := "INSERT INTO carriers(cid, company_name, address, locality_id) VALUES (?, ?, ?, ?)"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return models.Carrier{}, err
	}
	res, err := stmt.Exec(new.Cid, new.Company, new.Address, new.LocalityID)
	if err != nil {
		return models.Carrier{}, err
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return models.Carrier{}, err
	}
	if ra == 0 {
		return models.Carrier{}, customerrors.ErrorStoreFailed
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return models.Carrier{}, err
	}

	new.ID = int(lastID)
	return new, nil
}

func (m *mysqlDB) Get(id int) (models.CarriersReport, error) {
	var query string
	if id == 0 {
		query = "SELECT c.locality_id, l.locality_name, COUNT(*) AS carriers_count FROM carriers c INNER JOIN localities l ON c.locality_id = l.id GROUP BY c.locality_id"
	} else {
		query = fmt.Sprintf("SELECT c.locality_id, l.locality_name, COUNT(*) AS carriers_count FROM carriers c INNER JOIN localities l ON c.locality_id = l.id WHERE c.locality_id = %d GROUP BY c.locality_id", id)
	}

	res, err := m.db.Query(query)
	if err != nil {
		return models.CarriersReport{}, err
	}

	var report models.CarriersReport
	for res.Next() {
		var r models.CarrierInfo
		err := res.Scan(&r.LocalityID, &r.LocalityName, &r.CarriersCount)
		if err != nil {
			return models.CarriersReport{}, err
		}
		report.Data = append(report.Data, r)
	}

	if report.Data == nil {
		return models.CarriersReport{}, customerrors.ErrorItemNotFound
	}
	return report, nil
}
