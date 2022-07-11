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
	SectionList []Section `json:"sections"`
}

type SectionMetaData struct {
	LastID  int `json:"last_id"`
	Content Sections
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
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

type BuyersMetaData struct {
	LastID  int `json:"last_id"`
	Content Buyers
}

type Warehouse struct {
	ID                 int    `json:"id"`
	Address            string `json:"address" binding:"required"`
	Telephone          string `json:"telephone" binding:"required"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity" binding:"required"`
	MinimunTemperature int    `json:"minimun_temperature" binding:"required"`
	LocalityID         int    `json:"locality_id"`
}

type Warehouses struct {
	Warehouses []Warehouse `json:"warehouses"`
}

type WarehouseMetaData struct {
	LastID  int `json:"last_id"`
	Content Warehouses
}

type PostWarehouse struct {
	Address            string `json:"address" binding:"required"`
	Telephone          string `json:"telephone" binding:"required"`
	WarehouseCode      string `json:"warehouse_code" binding:"required"`
	MinimunCapacity    *int   `json:"minimun_capacity" binding:"required"`
	MinimunTemperature *int   `json:"minimun_temperature" binding:"required"`
	LocalityID         int    `json:"locality_id"`
}

type Products struct {
	Products []Product `json:"products"`
}

type ProductMetaData struct {
	LastID  int `json:"last_id"`
	Content Products
}

type Product struct {
	Id                             int     `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"netweight"`
	ExpirationRate                 int     `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductTypeId                  int     `json:"product_type_id" `
	SellerId                       int     `json:"seller_id"`
}

type Seller struct {
	ID           int    `json:"id"`
	Cid          int    `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	LocalityID   string `json:"locality_id"`
}

type Locality struct {
    Id            string `json:"id"`
    Locality_name string `json:"locality_name"`
    Province_name string `json:"province_name"`
    Country_name  string `json:"country_name"`
}

type ReportSeller struct {
    LocalityID    string `json:"locality_id"`
    Locality_name string `json:"locality_name"`
    SellerCount   string `json:"seller_count"`
}

type Sellers struct {
	Seller []Seller `json:"sellers"`
	LastID int      `json:"lastid"`
}

type Locality struct {
	ID         int    `json:"id"`
	Name       string `json:"locality_name"`
	ProvinceID int    `json:"province_id"`
}

type Carrier struct {
	ID         int    `json:"id"`
	Cid        int    `json:"cid"`
	Company    string `json:"company_name"`
	Address    string `json:"address"`
	Telephone  string `json:"telephone"`
	LocalityID int    `json:"locality_id"`
}

type CarrierRequest struct {
	Cid        *int    `json:"cid" binding:"required"`
	Company    *string `json:"company_name" binding:"required"`
	Address    *string `json:"address" binding:"required"`
	Telephone  string  `json:"telephone" binding:"required"`
	LocalityID *int    `json:"locality_id" binding:"required"`
}

type CarrierInfo struct {
	LocalityID    int    `json:"locality_id"`
	LocalityName  string `json:"locality_name"`
	CarriersCount int    `json:"carriers_count"`
}

type CarriersReport struct {
	Data []CarrierInfo `json:"reports"`
}
