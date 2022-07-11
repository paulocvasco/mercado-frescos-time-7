package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func init() {
	// path, _ := os.Getwd()
	// err := godotenv.Load(filepath.Join(path, ".env"))
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// dataSource := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("USER_DB"), os.Getenv("PASS_DB"), os.Getenv("PATH_DB"), os.Getenv("NAME_DB"))
	dataSource := "root:12345678@tcp(localhost:3306)/mercado_fresco_db?parseTime=true"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")
}
