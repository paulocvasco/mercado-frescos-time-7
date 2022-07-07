package repository

const (
	sqlStoreRecord = "INSERT INTO product_records (last_update_date, purchase_prince, sale_price, product_id) VALUES (?, ?, ?, ?)"
	
	sqlGetRecordById = "SELECT pr.product_id, p.description, COUNT(*) AS records_count FROM product_records pr INNER JOIN products p ON pr.product_id = p.id WHERE pr.product_id = ? GROUP BY pr.product_id"

	sqlGetAllRecords = "SELECT pr.product_id, p.description, COUNT(*) AS records_count FROM product_records pr INNER JOIN products p ON pr.product_id = p.id WHERE pr.product_id > ? GROUP BY pr.product_id"
)
