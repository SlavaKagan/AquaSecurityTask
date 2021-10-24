package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB = nil

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./aqua.db")
	checkErr(err)

	database = db
}

func GetDB() *sql.DB {
	if database == nil {
		panic("database not initialized")
	}
	return database
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
