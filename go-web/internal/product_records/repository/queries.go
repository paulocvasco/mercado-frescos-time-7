package repository

const (
	sqlStoreRecord = "INSERT INTO product_records (last_update_date, purchase_prince, sale_price, product_id) VALUES (?, ?, ?, ?)"
	
	sqlGetRecordById = "SELECT * FROM product_records WHERE product_id = ?"
)
