package repository

const (
	sqlGetAll = "SELECT * FROM products"

	sqlGetById = "SELECT * FROM products WHERE id = ?"

	sqlStore = "INSERT INTO products (description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	sqlUpdate = "UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=? WHERE id=?"

	sqlDelete = "DELETE FROM products WHERE id=?"
)
