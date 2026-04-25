# 💰 Expense API (Go + JWT + SQLite)

A production-style REST API built with Go (Golang) for managing personal expenses with authentication using JWT.

---

## 🚀 Live Demo
https://go-expense-api-production.up.railway.app/

---

## ✨ Features

- User Registration (bcrypt password hashing)
- User Login
- JWT Authentication
- Middleware Protected Routes
- CRUD Expense Management
- SQLite Database (lightweight & fast)
- RESTful API design
- Deployed on Railway (CI/CD via GitHub)

---

## 🧰 Tech Stack

- Go (net/http)
- SQLite (modernc driver)
- JWT (golang-jwt)
- bcrypt (password hashing)
- Railway (deployment)

---

## 🔐 Authentication Flow

1. Register user
2. Login user
3. Receive JWT token
4. Use token in request header:

---

## 📌 API Endpoints

### 🟢 Auth

#### Register

---

## 🟢 POST/register
Request:
```json
{
  "username": "test",
  "password": "123456"
}
```
---

## 🟢 Login
POST /login
``` id="4t9h5f"

Request:
```json
{
  "username": "test",
  "password": "123456"
}
```
---

## Response
```{
  "message": "login success",
  "token": "JWT_TOKEN_HERE"
}
```
---

## 💸 Expenses (Protected)
⚠️ All endpoints require JWT token
Get all expenses

GET /expenses
``` id="8v3x9k"
```
Header:
Authorization: Bearer YOUR_TOKEN
---

### Create expense
POST /expenses/create
Request:
```json
{
  "user_id": 1,
  "title": "Food",
  "amount": 50000
}
```

---

### Update expense
POST /expenses/update?id=1
``` id="9m2z8c"

Request:
```json
{
  "title": "Updated Food",
  "amount": 75000
}
```

---

## Delete expense
POST /expenses/delete?id=1

---

## 🧪 Example Curl
Login
curl -X POST https://go-expense-api-production.up.railway.app/login \
-H "Content-Type: application/json" \
-d '{"username":"test1","password":"123456"}'

---

## Get Expenses
curl https://go-expense-api-production.up.railway.app/expenses \
-H "Authorization: Bearer YOUR_TOKEN"

---

## 📦 Project Structure

```text
expense-api/
├── database/
├── handlers/
│   ├── auth_middleware.go
│   ├── jwt.go
│   ├── handlers.go
├── models/
├── main.go
├── expense.db
├── go.mod
```

## 🎯 Learning Goals
- REST API development in Go
- JWT authentication system
- Password hashing (bcrypt)
- Middleware security
- SQLite integration
- Deployment (GitHub → Railway)```

---

## 👨‍💻 Author
Built by Roja
