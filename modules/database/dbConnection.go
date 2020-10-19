package dbconnection

import (
	"database/sql"
	"log"

	// for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DbCon - database connection
var DbCon *sql.DB

// DBError - database connection error
var DBError error

// DbConnect - to connect database
func DbConnect() {
	log.Println("====== Starting Database MySQL ======")
	DbCon, DBError = sql.Open("mysql", "root@tcp(127.0.0.1:3307)/go_print")
	if DBError != nil {
		panic(DBError.Error())
	}
}

// DbDisConnect - to disconnect database
func DbDisConnect() {
	log.Println("====== Stopping Database MySQL ======")
	DbCon.Close()
}
