package handlers

import (
	"encoding/json"
	"net/http"

	"expense-api/database"
	"expense-api/models"
)

// ================= USERS =================

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result, err := database.DB.Exec(
		"INSERT INTO users(name) VALUES(?)",
		user.Name,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)

	json.NewEncoder(w).Encode(user)
}

// ================= EXPENSES =================

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, user_id, title, amount FROM expenses")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var e models.Expense
		rows.Scan(&e.ID, &e.UserID, &e.Title, &e.Amount)
		expenses = append(expenses, e)
	}

	json.NewEncoder(w).Encode(expenses)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var exp models.Expense
	json.NewDecoder(r.Body).Decode(&exp)

	result, err := database.DB.Exec(
		"INSERT INTO expenses(user_id, title, amount) VALUES(?,?,?)",
		exp.UserID,
		exp.Title,
		exp.Amount,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	exp.ID = int(id)

	json.NewEncoder(w).Encode(exp)
}

// ================= UPDATE EXPENSE =================

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var updated models.Expense
	json.NewDecoder(r.Body).Decode(&updated)

	_, err := database.DB.Exec(
		"UPDATE expenses SET title=?, amount=? WHERE id=?",
		updated.Title,
		updated.Amount,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "updated",
	})
}

// ================= DELETE EXPENSE =================

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := database.DB.Exec(
		"DELETE FROM expenses WHERE id=?",
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted",
	})
}

// ================= FILTER USER EXPENSE =================

func GetUserExpenses(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	rows, err := database.DB.Query(
		"SELECT id, user_id, title, amount FROM expenses WHERE user_id=?",
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var e models.Expense
		rows.Scan(&e.ID, &e.UserID, &e.Title, &e.Amount)
		expenses = append(expenses, e)
	}

	json.NewEncoder(w).Encode(expenses)
}
