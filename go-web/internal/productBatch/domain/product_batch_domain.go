package domain

import "context"

type ProductBatch struct {
	Id                 int    `json:"id"`
	BatchNumber        int    `json:"batch_number"`
	CurrentQuantity    int    `json:"current_quantity"`
	CurrentTemperature int    `json:"current_temperature"`
	DueDate            string `json:"due_date"`
	InitialQuantity    int    `json:"initial_quantity"`
	ManufacturingDate  string `json:"manufacturing_date"`
	ManufacturingHour  int    `json:"manufacturing_hour"`
	MinimumTemperature int    `json:"minimum_temperature"`
	ProductId          int    `json:"product_id"`
	SectionId          int    `json:"section_id"`
}

type ProductBatchRepository interface {
	CreateProductBatch(context.Context, *ProductBatch) (*ProductBatch, error)
}

type ProductBatchService interface {
	Store(context.Context, *ProductBatch) (*ProductBatch, error)
}
