package sections

import (
	"context"
	"mercado-frescos-time-7/go-web/internal/sections/domain"
)

type serviceSection struct {
	repository domain.SectionRepository
}

func NewServiceSection(s domain.SectionRepository) domain.SectionService {
	return &serviceSection{repository: s}
}

func (s *serviceSection) GetAll(ctx context.Context) (*domain.Sections, error) {
	sections, err := s.repository.GetAll(ctx)
	if err != nil {
		return &domain.Sections{}, err
	}

	return sections, nil
}

func (s *serviceSection) GetById(ctx context.Context, id int) (*domain.Section, error) {
	section, err := s.repository.GetById(ctx, id)
	if err != nil {
		return &domain.Section{}, err
	}

	return section, nil
}

func (s *serviceSection) Store(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	section, err := s.repository.Store(ctx, section)
	if err != nil {
		return &domain.Section{}, err
	}

	return section, nil
}

func (s *serviceSection) Update(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	current, err := s.GetById(ctx, section.Id)
	if err != nil {
		return &domain.Section{}, err
	}

	if section.SectionNumber > 0 {
		current.SectionNumber = section.SectionNumber
	}

	if section.CurrentTemperature > 0 || section.CurrentTemperature < 0 {
		current.CurrentTemperature = section.CurrentTemperature
	}

	if section.MinimumTemperature > 0 || section.MinimumTemperature < 0 {
		current.MinimumTemperature = section.MinimumTemperature
	}

	if section.CurrentCapacity > 0 {
		current.CurrentCapacity = section.CurrentCapacity
	}

	if section.MinimumCapacity > 0 {
		current.MinimumCapacity = section.MinimumCapacity
	}

	if section.MaximumCapacity > 0 {
		current.MaximumCapacity = section.MaximumCapacity
	}

	if section.WarehouseId > 0 {
		current.WarehouseId = section.WarehouseId
	}

	if section.ProductTypeId > 0 {
		current.ProductTypeId = section.ProductTypeId
	}

	section, err = s.repository.Update(ctx, current)

	if err != nil {
		return &domain.Section{}, err
	}

	return section, nil
}

func (s *serviceSection) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceSection) GetReportProducts(ctx context.Context, id int) (*domain.ProductReports, error) {
	productBatch, err := s.repository.GetReportProducts(ctx, id)

	if err != nil {
		return &domain.ProductReports{}, err
	}

	return productBatch, nil

}
