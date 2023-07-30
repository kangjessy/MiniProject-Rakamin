package config

import (
	"database/sql"
)

func DBConnection() (db *sql.DB, err error){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "miniproject"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return 
}