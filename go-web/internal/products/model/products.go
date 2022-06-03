package model

type Product struct {
	Id                               int     `json:"id"`
	Product_code                     string  `json:"product_code"`
	Description                      string  `json:"description"`
	Width                            int     `json:"width"`
	Height                           float64 `json:"height"`
	Lenght                           float64 `json:"lenght"`
	Net_weight                       float64 `json:"netweight"`
	Expiration_rate                  int     `json:"expiration_rate"`
	Recommended_freezing_temperature int     `json:"recommended_freezing_temperature"`
	Freezing_rate                    float64 `json:"freezing_rate"`
	Product_type_id                  int     `json:"product_type_id" `
	Seller_id                        int     `json:"seller_id"`
}
