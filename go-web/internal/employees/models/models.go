package employees

type Employee struct {
	ID           int    `json:"id"`
	CardNumberId int    `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WareHouseId  int    `json:"warehouse_id"`
}

type Employees struct {
	Employee []Employee `json:"employees"`
}
