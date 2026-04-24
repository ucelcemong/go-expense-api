package main

import (
	"fmt"
	"net/http"

	"expense-api/database"
	"expense-api/handlers"
)

//----------HOME-----------

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Expense API is running ...")
}

func main() {

	database.Connect()

	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/users/create", handlers.CreateUser)

	http.HandleFunc("/expenses", handlers.GetExpenses)
	http.HandleFunc("/expenses/create", handlers.CreateExpense)
	http.HandleFunc("/expenses/delete", handlers.DeleteExpense)
	http.HandleFunc("/expenses/update", handlers.UpdateExpense)
	http.HandleFunc("/user-expenses", handlers.GetUserExpenses)

	fmt.Println("Server running on port 4000...")
	http.ListenAndServe(":4000", nil)
}
