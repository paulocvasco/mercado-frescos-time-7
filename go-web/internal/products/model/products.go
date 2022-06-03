package model

type Product struct {
	Id                               int     `json:"id"`
	Product_code                     string  `json:"product_code"`
	Description                      string  `json:"description"`
	Width                            float64 `json:"width"`
	Height                           float64 `json:"height"`
	Length                           float64 `json:"lenght"`
	Net_weight                       float64 `json:"netweight"`
	Expiration_rate                  int     `json:"expiration_rate"`
	Recommended_freezing_temperature float64 `json:"recommended_freezing_temperature"`
	Freezing_rate                    float64 `json:"freezing_rate"`
	Product_type_id                  int     `json:"product_type_id" `
	Seller_id                        int     `json:"seller_id"`
}

var Products []Product = []Product{
	{
		Id:                               0,
		Description:                      "test",
		Expiration_rate:                  1,
		Freezing_rate:                    2,
		Height:                           6.4,
		Length:                           4.5,
		Net_weight:                       3.4,
		Product_code:                     "ssd",
		Recommended_freezing_temperature: 1.3,
		Width:                            1.2,
		Product_type_id:                  2,
		Seller_id:                        2,
	},
	{
		Id:                               1,
		Description:                      "test 2",
		Expiration_rate:                  2,
		Freezing_rate:                    2,
		Height:                           6.4,
		Length:                           4.5,
		Net_weight:                       3.4,
		Product_code:                     "ssd",
		Recommended_freezing_temperature: 1.3,
		Width:                            1.2,
		Product_type_id:                  2,
		Seller_id:                        2,
	},
}

var LastId int
