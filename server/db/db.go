package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		log.Fatal("Cannot open database:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	log.Println("Connected to SQLite database")
}
