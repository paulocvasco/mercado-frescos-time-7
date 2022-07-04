package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func init() {
	// user := os.Getenv("USERDB")
	// password := os.Getenv("PASSWORDDB")
	// portDB := os.Getenv("PORT")
	// nameDB := os.Getenv("NAMEDB")
	// dataSource := fmt.Sprint("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, portDB, nameDB)
	dataSource := "root:12345678@tcp(localhost:3306)/bootcamp?parseTime=true"
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
