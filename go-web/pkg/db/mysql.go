package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func InstanceDB(db *sql.DB) error {
	if db != nil {
		StorageDB = db
	}

	urlDB := "root:@tcp(localhost)/mercado_fresco_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// open connection with db
	StorageDB, err = sql.Open("mysql", urlDB)
	if err != nil {
		return err
	}

	// test connection
	if err = StorageDB.Ping(); err != nil {
		return err
	}

	// set db options
	StorageDB.SetConnMaxLifetime(time.Minute * 1)
	StorageDB.SetMaxOpenConns(10)
	StorageDB.SetMaxIdleConns(10)

	return nil
}

func Get() *sql.DB {
	return StorageDB
}
