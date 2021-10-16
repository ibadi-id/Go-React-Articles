package database

// Package untuk membuat koneksi dengan database

import (
	"database/sql"
	"log"

	"github.com/fatih/color"
)

var Db *sql.DB

func InitDB() {
	var err error

	// membuat koneksi dengan database
	Db, err = sql.Open("mysql", "testuser:tes12345@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	color.Green("Terhubung dengan Database MYSQL")
}
