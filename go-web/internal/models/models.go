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
