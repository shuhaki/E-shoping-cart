# E-shoping-cart


# 🛒 Full-Stack Shopping Cart App

This project is a complete **Shopping Cart System** built with:

- **Frontend**: React (React Router, Axios)
- **Backend**: Go (Gin) + MongoDB

---

## 📦 Features

- User Registration & Login
- Item Listing and Add to Cart
- Checkout (Create Orders)
- Order History

---

## 🧑‍💻 Backend Setup (Go + MongoDB)

### 1. Prerequisites

- Go 1.20+
- MongoDB running locally or MongoDB Atlas

### 2. Environment Variables

Create `.env` file:

```
MONGO_URI=mongodb://localhost:27017/shoppingcart
JWT_SECRET=your_jwt_secret
```

### 3. Run the Backend

```
go mod tidy
go run main.go
```

Server runs at: `http://localhost:8080`

---

## 💻 Frontend Setup (React)

### 1. Install Dependencies

```
cd shopping-cart-frontend
npm install
```

### 2. Run React Dev Server

```
npm start
```

React app runs at: `http://localhost:3000`

---

##POSTMAN COLLECTION = https://github.com/shuhaki/E-shoping-cart/blob/main/ShoppingCart.postman_collection.json

## 🔁 API Collection (Postman)

Use the included Postman collection to test:

- `/api/users` → Register or Login
- `/api/items` → Get/Add Items
- `/api/carts` → Add item to cart
- `/api/orders` → Create Orders

Import the file: `ShoppingCart.postman_collection.json` into Postman.

---

## 🔐 Auth Notes

- After login, a `token` and `user_id` are saved to `localStorage`
- All routes use Bearer Token via `Authorization` header

---

## 🛠 Built With

- [React](https://reactjs.org)
- [Axios](https://axios-http.com)
- [React Router](https://reactrouter.com)
- [Gin Web Framework](https://gin-gonic.com)
- [MongoDB Go Driver](https://go.mongodb.org/mongo-driver)
