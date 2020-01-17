package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var DB *sql.DB

func DBConn() (db *sql.DB) {
	dbName := os.Getenv("dbName")
	dbPass := os.Getenv("dbPass")
	dbHost := os.Getenv("dbHost")
	dbUser := os.Getenv("dbUser")

	var err error
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
