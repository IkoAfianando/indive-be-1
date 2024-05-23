package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var DB *sql.DB

func Init() {
	var err error
	mysqlUrl := os.Getenv("MYSQL_URL")
	DB, err = sql.Open("mysql", mysqlUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
