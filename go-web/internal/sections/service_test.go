package sections_test

import (
	"context"
	"database/sql"
	"mercado-frescos-time-7/go-web/internal/sections"
	"mercado-frescos-time-7/go-web/internal/sections/domain"
	"mercado-frescos-time-7/go-web/internal/sections/mock/mockRepository"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestGetAllOk(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)
	sections := &domain.Sections{
		Sections: []domain.Section{
			{
				Id:                 1,
				SectionNumber:      1,
				CurrentTemperature: 1,
				MinimumTemperature: 1,
				CurrentCapacity:    1,
				MinimumCapacity:    1,
				MaximumCapacity:    1,
				WarehouseId:        1,
				ProductTypeId:      1,
			},
			{
				Id:                 2,
				SectionNumber:      2,
				CurrentTemperature: 2,
				MinimumTemperature: 2,
				CurrentCapacity:    2,
				MinimumCapacity:    2,
				MaximumCapacity:    2,
				WarehouseId:        2,
				ProductTypeId:      2,
			},
		},
	}

	repo.On("GetAll", mock.Anything).Return(sections, nil).Maybe()
	res, _ := serv.GetAll(context.TODO())
	assert.Equal(t, sections, res)

}
func TestGetAllError(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)

	repo.On("GetAll", mock.Anything).Return(&domain.Sections{}, sql.ErrNoRows).Maybe()
	res, err := serv.GetAll(context.TODO())
	assert.Equal(t, &domain.Sections{}, res)
	assert.Equal(t, sql.ErrNoRows, err)

}

func TestGetByIdOk(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)
	section := &domain.Section{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	repo.On("GetById", mock.Anything, mock.Anything).Return(section, nil).Maybe()
	res, _ := serv.GetById(context.TODO(), 1)
	assert.Equal(t, section, res)

}
func TestGetByIdError(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)

	repo.On("GetById", mock.Anything, mock.Anything).Return(&domain.Section{}, sql.ErrNoRows).Maybe()
	res, err := serv.GetById(context.TODO(), 1)
	assert.Equal(t, &domain.Section{}, res)
	assert.Equal(t, sql.ErrNoRows, err)

}

func TestCreateOK(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)
	section := &domain.Section{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	repo.On("Store", mock.Anything, mock.Anything).Return(section, nil).Maybe()
	result, _ := serv.Store(context.TODO(), section)
	assert.Equal(t, section, result)
}
func TestCreateError(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)
	section := &domain.Section{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	repo.On("Store", mock.Anything, mock.Anything).Return(&domain.Section{}, sql.ErrNoRows).Maybe()
	result, _ := serv.Store(context.TODO(), section)
	assert.Equal(t, &domain.Section{}, result)
}

func TestUpdateOK(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)
	section := &domain.Section{
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    2,
		MaximumCapacity:    2,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	repo.On("Update", mock.Anything, mock.Anything).Return(section, nil).Maybe()
	repo.On("GetById", mock.Anything, mock.Anything).Return(section, nil).Maybe()
	result, _ := serv.Update(context.TODO(), section)
	assert.Equal(t, section, result)
}
func TestUpdateErrorGetById(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)
	section := &domain.Section{
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    2,
		MaximumCapacity:    2,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	repo.On("GetById", mock.Anything, mock.Anything).Return(section, sql.ErrNoRows).Maybe()
	repo.On("Update", mock.Anything, mock.Anything).Return(section, nil).Maybe()

	result, err := serv.Update(context.TODO(), section)

	assert.Equal(t, &domain.Section{}, result)
	assert.Equal(t, sql.ErrNoRows, err)
}
func TestUpdateError(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)
	section := &domain.Section{
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    2,
		MaximumCapacity:    2,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	repo.On("GetById", mock.Anything, mock.Anything).Return(section, nil).Maybe()
	repo.On("Update", mock.Anything, mock.Anything).Return(section, sql.ErrNoRows).Maybe()

	result, err := serv.Update(context.TODO(), section)

	assert.Equal(t, &domain.Section{}, result)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestDeleteOk(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)

	repo.On("Delete", mock.Anything, mock.Anything).Return(nil).Maybe()
	err := serv.Delete(context.TODO(), 1)
	assert.Equal(t, err, nil)
}
func TestDeleteError(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)

	repo.On("Delete", mock.Anything, mock.Anything).Return(sql.ErrNoRows).Maybe()
	err := serv.Delete(context.TODO(), 1)
	assert.Equal(t, err, sql.ErrNoRows)
}

func TestReportProductsOk(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)
	reportProducts := &domain.ProductReports{
		ProductReports: []domain.ProductReport{
			{
				SectionId:     1,
				SectionNumber: 1,
				ProductsCount: 1,
			},
			{
				SectionId:     2,
				SectionNumber: 2,
				ProductsCount: 10,
			},
		},
	}

	repo.On("GetReportProducts", mock.Anything, mock.Anything).Return(reportProducts, nil).Maybe()
	res, _ := serv.GetReportProducts(context.TODO(), 0)
	assert.Equal(t, reportProducts, res)

}
func TestReportProductsError(t *testing.T) {
	repo := mockRepository.NewSectionRepository(t)
	serv := sections.NewServiceSection(repo)

	repo.On("GetReportProducts", mock.Anything, mock.Anything).Return(&domain.ProductReports{}, sql.ErrNoRows).Maybe()
	res, err := serv.GetReportProducts(context.TODO(), -1)

	assert.Equal(t, &domain.ProductReports{}, res)
	assert.Equal(t, sql.ErrNoRows, err)
}
