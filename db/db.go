package db

import (
	"database/sql"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "/data/recipes.db")
	if err != nil {
		return err
	}
	sqlBytes, err := os.ReadFile("init.sql")
	if err == nil {
		DB.Exec(string(sqlBytes))
	}
	return DB.Ping()
}