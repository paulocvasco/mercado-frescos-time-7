package repository

import (
	"context"
	"database/sql"
	"fmt"
	"mercado-frescos-time-7/go-web/internal/sections/domain"
	"mercado-frescos-time-7/go-web/pkg/logger"
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
		logger.Save(err.Error())
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
			logger.Save(err.Error())
			return &sections, err
		}

		sections.Sections = append(sections.Sections, section)
	}

	logger.Save(logger.SectionsResquested)
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

	if err != nil {
		logger.Save(err.Error())
		return &section, err
	}

	logger.Save(fmt.Sprintf(logger.SectionRequested, id))
	return &section, nil
}
func (r *repositorySql) Store(ctx context.Context, section *domain.Section) (*domain.Section, error) {

	stmt, err := r.db.Prepare(queryStore)

	if err != nil {
		logger.Save(err.Error())
		return &domain.Section{}, err
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
		logger.Save(err.Error())
		return &domain.Section{}, err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		logger.Save(err.Error())
		return &domain.Section{}, err
	}

	section.Id = int(lastID)

	logger.Save(fmt.Sprintf(logger.SectionCreated, lastID))
	return section, nil

}
func (r *repositorySql) Update(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	stmt, err := r.db.Prepare(queryUpdate)
	if err != nil {
		logger.Save(err.Error())
		return &domain.Section{}, err
	}

	_, err = stmt.ExecContext(
		ctx,
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
		logger.Save(err.Error())
		return &domain.Section{}, err
	}

	logger.Save(fmt.Sprintf(logger.SectionChanged, section.Id))
	return section, nil
}
func (r *repositorySql) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(queryDelete)
	if err != nil {
		logger.Save(err.Error())
		return err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		logger.Save(err.Error())
		return err
	}
	return nil
}

func (r *repositorySql) GetReportProducts(ctx context.Context, id int) (*domain.ProductReports, error) {
	var query string
	if id == 0 {
		query = queryGetAllReportProducts
	} else {
		query = queryReportProductsById
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		logger.Save(err.Error())
		return &domain.ProductReports{}, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		logger.Save(err.Error())
		return &domain.ProductReports{}, err
	}
	reports := domain.ProductReports{ProductReports: []domain.ProductReport{}}
	for rows.Next() {
		report := domain.ProductReport{}
		err := rows.Scan(&report.SectionId, &report.SectionNumber, &report.ProductsCount)
		if err != nil {
			logger.Save(err.Error())
			return &domain.ProductReports{}, err
		}
		reports.ProductReports = append(reports.ProductReports, report)
	}

	logger.Save(fmt.Sprintf(logger.SectionRequested, id))
	return &reports, nil
}
