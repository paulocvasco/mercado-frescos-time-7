package models

type Seller struct {
	ID    int  `json:"id"`
	Cid  string  `json:"cid"`
	Company_name string  `json:"company_name"`
	Address string  `json:"address"`
	Telephone  string `json:"telephone"`
}