package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Expense struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}
