Shopping Cart – Fullstack Assignment

. This project implements a simple e-commerce flow as described in the assignment.
  Tech stack:
. Backend: Go, Gin, GORM, SQLite
. Frontend: React (Create React App)
. Auth: Token-based (single active token per user)
# How to Run the Project
1. Clone the Repository
git clone <https://github.com/prabhanshu451/Shoppingcart_Prabhanshu_Upadhyay_2200971630048.git>
cd shopping-cart

# Backend Setup (Go)
* Inside the /backend folder:

1. Install dependencies
go mod tidy

2. Run the backend
go run .

3. Backend starts at:
http://localhost:8080

# Frontend Setup (React)
* Inside the /frontend folder:

1. Install dependencies
npm install

2. Run the frontend
npm start

3. Frontend starts at:
http://localhost:3000

# API Usage Guide
1️. Signup
POST /users
{
  "username": "prabh",
  "password": "pass"
}

2️. Login (returns token)
POST /users/login
{
  "username": "prabh",
  "password": "pass"
}

3. Response:
{ "token": "<your-token>" }

4. Store this token and send it in the header for cart/order APIs:
Authorization: <token>

# Items
Create Item
1. POST /items
{ "name": "T-Shirt" }

2. List Items
GET /items

# Cart APIs (Requires token)
Add Item to Cart
1. POST /carts
Header: Authorization: <token>
Body:
{ "item_id": 1 }

2. List Carts
GET /carts

# Order APIs (Requires token)
Create Order
1. POST /orders
{
  "cart_id": 1
}

2. List Orders
GET /orders

# Postman Collection

You can find the Postman collection here:
/postman_collection.json

# Project Structure
backend/
  handlers.go
  models.go
  router.go
  middleware.go
  shopping.db

frontend/
  src/
    Login.js
    Items.js
    apis.js
  public/
  package.json

postman_collection.json


******************Thank You*****************
