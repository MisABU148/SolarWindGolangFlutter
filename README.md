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
<img width="596" height="486" alt="image" src="https://github.com/user-attachments/assets/46de6bb7-e999-49df-84e7-b69f2f56e18d" />


## 📸 Screenshots

Dark mode
<img width="590" height="1280" alt="image" src="https://github.com/user-attachments/assets/5dcb83f5-6cc8-40e2-ab5e-282524b897a0" />
<img width="590" height="1280" alt="image" src="https://github.com/user-attachments/assets/8f68df49-e12e-4b1c-9015-cdb412989214" /

Different languages
<img width="590" height="1280" alt="image" src="https://github.com/user-attachments/assets/c6597243-1664-4cb9-82de-753b8e7f4993" />
<img width="590" height="1280" alt="image" src="https://github.com/user-attachments/assets/47d6f749-444b-4342-ae40-f72f13e2b3a6" />

Telegram authentication
<img width="590" height="1280" alt="image" src="https://github.com/user-attachments/assets/c8b01dd8-5d99-4bc0-916c-de7c3c257356" />

Olso we will talk about features iin presentation

📋 Project Management
🗂️ GitHub Projects for tracking issues
It was deleted the project by github, now they in issues [project board](https://github.com/SolarWindApp/SolarWindGolangFlutter/issues?q=is%3Aissue%20state%3Aclosed)

✅ Regular meaningful pull-requests

## 👥 Team:

Backend (Go): Ilyina Maria
Frontend (Flutter): Daria Nikolaeva
