package repository

const (
	queryGetAll = "SELECT * FROM sections;"

	queryGetById = "SELECT * FROM sections WHERE id = ?;"

	queryStore = "INSERT INTO sections" +
		"(section_number, current_temperature, minimum_temperature, current_capacity," +
		"minimum_capacity, maximum_capacity, warehouse_id, product_type_id)" +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?);"

	queryUpdate = "UPDATE sections SET " +
		"section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?," +
		"minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?;"

	queryDelete = "DELETE FROM sections WHERE id=?;"

	queryReportProductsById   = "SELECT s.id, s.section_number, COUNT(*) as products_count from sections s INNER JOIN products_batches pb ON s.id = pb.section_id WHERE s.id = ? GROUP BY(s.id);"
	queryGetAllReportProducts = "SELECT s.id, s.section_number, COUNT(*) as products_count from sections s Inner join products_batches p on s.id = p.section_id WHERE s.id > ? GROUP BY(s.id);"
)
