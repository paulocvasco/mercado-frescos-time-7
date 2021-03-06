package domain

import "context"

type Section struct {
	Id                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseId        int `json:"warehouse_id"`
	ProductTypeId      int `json:"product_type_id"`
}

type Sections struct {
	Sections []Section `json:"sections"`
}

type ProductReport struct {
	SectionId     int `json:"section_id"`
	SectionNumber int `json:"section_number"`
	ProductsCount int `json:"products_count"`
}

type ProductReports struct {
	ProductReports []ProductReport `json:"product_reports"`
}

//go:generate mockery --name=SectionRepository --output=./mock/mockRepository --outpkg=mockRepository
type SectionRepository interface {
	GetAll(context.Context) (*Sections, error)
	GetById(context.Context, int) (*Section, error)
	Store(context.Context, *Section) (*Section, error)
	Update(context.Context, *Section) (*Section, error)
	Delete(context.Context, int) error

	GetReportProducts(context.Context, int) (*ProductReports, error)
}

//go:generate mockery --name=SectionService --output=./mock/mockService --outpkg=mockService
type SectionService interface {
	GetAll(context.Context) (*Sections, error)
	GetById(context.Context, int) (*Section, error)
	Store(context.Context, *Section) (*Section, error)
	Update(context.Context, *Section) (*Section, error)
	Delete(context.Context, int) error

	GetReportProducts(context.Context, int) (*ProductReports, error)
}
