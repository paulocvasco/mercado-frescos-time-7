package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var StorageDB *sql.DB

func init() {
	urlDB := "root:@tcp(localhost)/mercado_fresco_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// open connection with db
	db, err = sql.Open("mysql", urlDB)
	if err != nil {
		panic(err)
	}

	// test connection
	if err = db.Ping(); err != nil {
		panic(err)
	}

	// set db options
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	log.Println("Db connected")
}

func Get() *sql.DB {
	return StorageDB
}
