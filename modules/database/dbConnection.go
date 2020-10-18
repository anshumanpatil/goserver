package dbConnection

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DbCon *sql.DB
var DBError error


func DbConnect(){
    log.Println("====== Starting Database MySQL ======")
	DbCon, DBError = sql.Open("mysql", "root@tcp(127.0.0.1:3307)/go_print")
	if DBError != nil {
        panic(DBError.Error())
    }
}

func DbDisConnect(){
    log.Println("====== Stopping Database MySQL ======")
    DbCon.Close()
}