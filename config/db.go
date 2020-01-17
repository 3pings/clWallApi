package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var DB *sql.DB

func init() {
	dbName := os.Getenv("dbName")
	dbPass := os.Getenv("dbPass")
	dbHost := os.Getenv("dbHost")
	dbUser := os.Getenv("dbUser")

	var err error
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("Connected to ", dbHost)
	}
}
