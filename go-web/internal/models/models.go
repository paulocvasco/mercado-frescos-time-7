package models

type Section struct {
	ID                 int `json:"id"`
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
	Section []Section `json:"sections"`
}
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

type Warehouse struct {
	ID                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

type Product struct {
	Id                             int     `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"lenght"`
	NetWeight                      float64 `json:"netweight"`
	ExpirationRate                 int     `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductTypeId                  int     `json:"product_type_id" `
	SellerId                       int     `json:"seller_id"`
}

var Products []Product = []Product{
	{
		Id:                             0,
		Description:                    "test",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	},
	{
		Id:                             1,
		Description:                    "test 2",
		ExpirationRate:                 2,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                   "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	},
}

var LastId int
type Seller struct { 
	ID    int  `json:"id"`
	Cid  int  `json:"cid"`
	Company_name string  `json:"company_name"`
	Address string  `json:"address"`
	Telephone  string `json:"telephone"`
}

