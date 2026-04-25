package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("sqlite", "expense.db")
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")

	createTables()
}

func createTables() {
	userTable := `
CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        password TEXT
);`

	expenseTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT,
		amount REAL
	);`

	_, err := DB.Exec(userTable)
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec(expenseTable)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tables ready!")
}
