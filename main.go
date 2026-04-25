package main

import (
	"fmt"
	"net/http"
	"os"

	"expense-api/database"
	"expense-api/handlers"
)

// HOME
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Expense API is running ...")
}

func main() {
	fmt.Println("APP STARTING...")

	// Connect database
	database.Connect()

	fmt.Println("AFTER DB CONNECT")

	// Routes
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/users/create", handlers.CreateUser)

	http.HandleFunc("/expenses", handlers.GetExpenses)
	http.HandleFunc("/expenses/create", handlers.CreateExpense)
	http.HandleFunc("/expenses/update", handlers.UpdateExpense)
	http.HandleFunc("/expenses/delete", handlers.DeleteExpense)
	http.HandleFunc("/user-expenses", handlers.GetUserExpenses)

	// Ambil PORT dari Railway
	port := os.Getenv("PORT")

	// Debug: cek port dari Railway
	fmt.Println("PORT from Railway:", port)

	// WAJIB: Railway harus kasih port
	if port == "" {
		fmt.Println("ERROR: PORT not set")
		return
	}

	fmt.Println("Server running on port", port)

	// Start server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
