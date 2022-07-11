package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var StorageDB *sql.DB

func init() {
	dataSource := "root:@tcp(localhost)/mercado_fresco?charset=utf8mb4&parseTime=True&loc=Local"
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
