package models

type Employee struct {
	ID           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WareHouseId  int    `json:"warehouse_id"`
}
type Buyers struct {
	Buyer []Buyer `json:"buyers"`
}
type Buyer struct {
	ID           int    `json:"id"`
	CardNumberID int    `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}
type Warehouses struct {
	Warehouse []Warehouse `json:"warehouses"`
}

type Warehouse struct {
	ID                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}
