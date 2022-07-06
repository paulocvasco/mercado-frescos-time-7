package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"mercado-frescos-time-7/go-web/internal/sections/domain"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type repositorySql struct {
	db *sql.DB
}

func NewRepositorySection(db *sql.DB) domain.SectionRepository {
	return &repositorySql{
		db: db,
	}
}

func (r *repositorySql) GetAll(ctx context.Context) (*domain.Sections, error) {
	var sections domain.Sections

	rows, err := r.db.QueryContext(ctx, queryGetAll)
	if err != nil {
		return &domain.Sections{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var section domain.Section

		if err := rows.Scan(
			&section.Id,
			&section.SectionNumber,
			&section.CurrentTemperature,
			&section.MinimumTemperature,
			&section.CurrentCapacity,
			&section.MinimumCapacity,
			&section.MaximumCapacity,
			&section.WarehouseId,
			&section.ProductTypeId,
		); err != nil {
			return &sections, err
		}

		sections.Sections = append(sections.Sections, section)
	}
	return &sections, nil
}
func (r *repositorySql) GetById(ctx context.Context, id int) (*domain.Section, error) {
	row := r.db.QueryRowContext(ctx, queryGetById, id)

	section := domain.Section{}

	err := row.Scan(
		&section.Id,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseId,
		&section.ProductTypeId,
	)

	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return &section, err
	}

	// Other errors
	if err != nil {
		return &section, err
	}

	return &section, nil
}
func (r *repositorySql) Store(ctx context.Context, section *domain.Section) (*domain.Section, error) {

	stmt, err := r.db.Prepare(queryStore)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result

	result, err = r.db.ExecContext(
		ctx,
		queryStore,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseId,
		&section.ProductTypeId,
	)

	if err != nil {
		return &domain.Section{}, err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return &domain.Section{}, err
	}

	section.Id = int(lastID)

	return section, nil

}
func (r *repositorySql) Update(ctx context.Context, section *domain.Section) error {
	result, err := r.db.ExecContext(
		ctx,
		queryUpdate,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseId,
		&section.ProductTypeId,
		&section.Id,
	)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		return customerrors.ErrorInvalidID
	} else if err != nil {
		return customerrors.ErrorInvalidDB
	}
	return nil
}
func (r *repositorySql) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(queryDelete)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	if rowsAffected, err := res.RowsAffected(); rowsAffected == 0 {
		return customerrors.ErrorInvalidID
	} else if err != nil {
		return customerrors.ErrorInvalidDB
	}
	return nil
}

func (r *repositorySql) GetReportProductsById(ctx context.Context, id int) (*domain.ProductReports, error) {
	return &domain.ProductReports{}, nil
}
