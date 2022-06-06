package models

type Section struct {
	ID                  int `json:"id"`
	Section_number      int `json:"section_number"`
	Current_temperature int `json:"current_temperature"`
	Minimum_temperature int `json:"minimum_temperature"`
	Current_capacity    int `json:"current_capacity"`
	Minimum_capacity    int `json:"minimum_capacity"`
	Maximum_capacity    int `json:"Maximum_capacity"`
	Warehouse_id        int `json:"warehouse_id"`
	Product_type_id     int `json:"product_type_id"`
}

type Sections struct {
	Section []Section `json:"sections"`
}
