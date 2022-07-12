package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func init() {
	urlDB := "root:@tcp(localhost)/mercado_fresco_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// open connection with db
	StorageDB, err = sql.Open("mysql", urlDB)
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

	// test connection
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}

	// set db options
	StorageDB.SetConnMaxLifetime(time.Minute * 1)
	StorageDB.SetMaxOpenConns(10)
	StorageDB.SetMaxIdleConns(10)

	log.Println("Db connected")
}

func Get() *sql.DB {
	return StorageDB
}
