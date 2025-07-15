# 🏋️‍♂️ Solar Wind – Workout Partner Finder

## 📌 Project Overview

**Solar Wind** is a lightweight matchmaking platform designed to help people find gym or workout partners in their area. Users can:

- Create a personal profile with sports interests, city, and gym schedule
- Browse and filter other users based on preferences
- Connect with like-minded individuals for workouts

> Think of it as a fitness-focused social app — not for dating, but for building healthy training habits together.

🎨 **Design prototype**: [Figma Design]([https://www.figma.com/file/YOUR-FIGMA-LINK-HERE](https://www.figma.com/design/si98563MfBSXuDtOfV8655/FitFlame?t=LBPNpHfkUVk9VxQt-0))

---

## 🚀 How to Run the Project

### 1. Clone the repository

```bash
git clone https://github.com/your-username/gymbuddy-backend.git
cd gymbuddy-backend
```
### 2. Install dependencies
Make sure you have Go 1.20+ installed.

```bash
go mod tidy
```

### 3. Run the application
```bash
go run main.go
```
The backend will start on:
📍 http://localhost:8080

### 4. Project Structure
```graphql
backend/
├── main.go                 # Entry point
├── controller/             # HTTP handlers
├── service/                # Business logic
├── repository/             # Data access layer
├── model/                  # Data models & DTOs
├── db/
│   ├── schema.sql          # DB schema
│   └── json/               # Seed data for cities and sports
├── auth/                   # JWT authentication
```
🔐 Authentication
🧾 Registration
```http
POST /api/me
Create a new user
```

Accepts username, password, profile info

🔑 Login
```http
POST /api/login
```
Returns a JWT token

Use the token to access protected endpoints

🔒 Authorization
Include the token in the request header:

```makefile
Authorization: Bearer YOUR_JWT_TOKEN
📡 Key API Endpoints
Method	Endpoint	Description
POST	/api/me	Register a new user
POST	/api/login	Login and receive JWT
GET	/api/me	Get your profile
PUT	/api/me	Update your profile
DELETE	/api/me	Delete your account
GET	/api/getUsers	Get all users
GET	/api/cities	Get list of cities
GET	/api/sports	Get list of sports
GET	/api/cities/search	Search cities
GET	/api/sports/search	Search sports
```
🛠 Tech Stack
Golang (net/http)

SQLite (via modernc.org/sqlite)

JSON-based seeding for cities & sports

JWT authentication
