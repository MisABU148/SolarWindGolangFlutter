# 🏋️‍♂️ Solar Wind – Workout Partner Finder

## 📌 Overview

**Solar Wind** is a platform for finding sports partners nearby. No dating, just healthy habits, joint training and motivation.

🔎 Users can:
- Register via Telegram
- Create and edit a profile (interests, schedule, city)
- View the feed of suitable partners
- Send likes (if it matches, friendship!)
- Use on any device: Android, iOS, Web
- Switch between dark/light theme and languages (RU/EN)

🎨 **[Figma дизайн (демо)]([https://www.figma.com/design/si98563MfBSXuDtOfV8655/FitFlame?t=LBPNpHfkUVk9VxQt-0](https://www.figma.com/design/si98563MfBSXuDtOfV8655/FitFlame?node-id=27-388&t=AbeqimJUuoFLRUwz-1))**  
📲 **Деплой (Frontend)**: ([https://solarwind-app.vercel.app](https://misabu148.github.io/SolarWindGolangFlutter/)

---

## 🏗️ Architecture

![Architecture Diagram]
[frintend](https://drive.google.com/file/d/1cFFs8dMEnyk5t88QDQKv6kpkd2mc_QVq/view?usp=sharing)
[backend](https://drive.google.com/file/d/13fDBf8rgqogzYe9W8dYGhZZLVboNwCkv/view?usp=sharing)

- **Backend**: Golang REST API, PostgreSQL, Telegram Auth
- **Frontend**: Flutter (Web + Mobile), CI/CD, i18n
- **Auth**: JWT, Telegram Bot
- **Infra**: Docker Compose, GitHub Actions CI/CD

---

## 📦 Tech Stack

| Layer     | Tech                                             |
|-----------|--------------------------------------------------|
| Backend   | Go (net/http), PostgreSQL, JWT, Telegram API     |
| Frontend  | Flutter (Web, Android, iOS)                      |
| Database  | PostgreSQL                                       |
| Auth      | Telegram + JWT                                   |
| DevOps    | Docker Compose, GitHub Actions, Vercel           |
| Testing   | Unit + Integration (Go), Widget Tests (Flutter)  |
| UI        | Responsive UI, Dark Mode, Custom Animations      |
| I18n      | RU 🇷🇺 / EN 🇬🇧 localization                       |

---

## 🚀 Getting Started

### Backend Setup

```bash
git clone https://github.com/your-org/gymbuddy-backend.git
cd gymbuddy-backend
go mod tidy
docker-compose up --build
```

## Backend available at:
📍 [deploy](https://solar-wind-gymbro.ru/)
## Swagger:
📍 (https://solar-wind-gymbro.ru/deckShuffle/swagger-ui/index.html#/)

Frontend Setup
```bash
cd flutter-app
flutter pub get
flutter run -d chrome
```
🔐 Authentication
```http
Endpoint	Method	Description
/api/login	POST	Вход через Telegram Bot
/api/me	GET	Получить профиль
/api/me	PUT	Обновить профиль
/api/me	DELETE	Удалить аккаунт
```

📦 Use JWT in header:
Authorization: Bearer YOUR_TOKEN

🌐 API Overview
```http
Method	Endpoint	Description
POST	/api/me	Создать/обновить профиль
GET	/api/getUsers	Получить список пользователей
GET	/api/cities	Список городов
GET	/api/sports	Список видов спорта
POST	/api/match	Лайкнуть пользователя
GET	/api/create-deck	Лента подходящих людей
```

🧪 Testing
✅ Unit tests (controllers, services)
✅ API integration tests
✅ Widget tests for the Flutter interface
✅ Login flow, feed, match coverage

# Backend
```bach
go test ./...
```

# Flutter
flutter test
🧩 Features Summary
Telegram Auth | ✅ Implemented
JWT Authentication | ✅ Implemented
Create/Edit Profile | ✅ Implemented
Matching Feed | ✅ Implemented
Likes & Matches | ✅ Implemented
Like Animation | ✅ Implemented
RU/EN Support | ✅ Implemented
Dark/Light Theme | ✅ Implemented
Docker Compose | ✅ Implemented
CI/CD (GitHub Actions/Vercel) | ✅ Implemented

## 🗃️ Database Schema
📊 The database schema is shown below:

([https://drive.google.com/file/d/1582aUNTyGs4boeZuKjFdMgATzkniZMgu/view?usp=sharing](https://drive.google.com/file/d/1582aUNTyGs4boeZuKjFdMgATzkniZMgu/view?usp=sharing))


## 📸 Screenshots

Dark mode
(https://drive.google.com/file/d/1rUrA0LvTAk-p4qhTr9KQ9zXwctN7TPTN/view?usp=sharing)
(https://drive.google.com/file/d/13zPn_W-6SLXOZwo6mgkHfUXmyQMrz9DY/view?usp=sharing)

Different languages
(https://drive.google.com/file/d/1wx7N1y_sst97yux6Q5WILJkv_42EZDRw/view?usp=sharing)
(https://drive.google.com/file/d/1PsMHlrNdQYnBHbyt3H27Xje_CxXMhimc/view?usp=sharing)

Telegram authentication
(https://drive.google.com/file/d/105p0MRWr_530TB5UxI8CnBZADy4zD5ae/view?usp=sharing)


📋 Project Management
🗂️ GitHub Projects for tracking issues

✅ Regular meaningful pull-requests

## 👥 Team:

Backend (Go): Ilyina Maria
Frontend (Flutter): Daria Nikolaeva
