package models

type Buyers struct {
	Buyer []Buyer `json:"buyers"`
}
type Buyer struct {
	ID           int    `json:"id"`
	CardNumberID int    `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
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
