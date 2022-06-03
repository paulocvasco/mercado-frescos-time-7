package model

type products struct {
	Id                               int    `json:"id"`
	Product_code                     string `json:"product_code"`
	Description                      string `json:"description"`
	Width                            int    `json:"width"`
	Height                           int    `json:"height"`
	Lenght                           int    `json:"lenght"`
	Net_weight                       int    `json:"netweight"`
	Expiration_rate                  int    `json:"expiration_rate"`
	Recommended_freezing_temperature int    `json:"recommended_freezing_temperature"`
	Freezing_rate                    string `json:"freezing_rate"`
	Product_type_id                  string `json:"product_type_id" `
	Seller_id                        string `json:"seller_id"`
}
