package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	urlDB := "root:@tcp(localhost)/mercado_db?charset=utf8mb4&parseTime=True&loc=Local"
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
}

func Get() sql.DB {
	return *db
}
