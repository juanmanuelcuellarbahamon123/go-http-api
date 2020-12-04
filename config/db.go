package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (db *sql.DB, err error) {
	driver := "mysql"
	user := "root"
	password := ""
	database := "prueba"
	db, err = sql.Open(driver, user+":"+password+"@/"+database)
	return
}
