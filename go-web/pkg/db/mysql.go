package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func init() {
	dataSource := "root:@tcp(localhost)/mercado_db?charset=utf8mb4&parseTime=True&loc=Local"
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
