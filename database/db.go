package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open("sqlite3", "./expense.db")
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	createTable()
}

func createTable() {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);`

	expenseTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT,
		amount REAL
	);`

	DB.Exec(userTable)
	DB.Exec(expenseTable)
}
