# 🏋️‍♂️ Solar Wind – Workout Partner Finder

## 📌 Overview

**Solar Wind** — это платформа для поиска спортивных напарников поблизости. Никакого дейтинга — только здоровые привычки, совместные тренировки и мотивация.

🔎 Пользователи могут:
- Зарегистрироваться через Telegram
- Создать и редактировать профиль (интересы, график, город)
- Смотреть ленту подходящих партнёров
- Отправлять лайки (при совпадении – дружба!)
- Пользоваться на любом устройстве: Android, iOS, Web
- Переключаться между тёмной/светлой темой и языками (RU/EN)

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

Backend будет доступен по адресу:
📍 http://localhost:8080

Frontend Setup
```bash
cd flutter-app
flutter pub get
flutter run -d chrome
```
🔐 Authentication
Endpoint	Method	Description
/api/login	POST	Вход через Telegram Bot
/api/me	GET	Получить профиль
/api/me	PUT	Обновить профиль
/api/me	DELETE	Удалить аккаунт

📦 Используйте JWT в заголовке:

```http
Authorization: Bearer YOUR_TOKEN
🌐 API Overview
Method	Endpoint	Description
POST	/api/me	Создать/обновить профиль
GET	/api/getUsers	Получить список пользователей
GET	/api/cities	Список городов
GET	/api/sports	Список видов спорта
POST	/api/match	Лайкнуть пользователя
GET	/api/feed	Лента подходящих людей
```

📘 Swagger-документация: /swagger/index.html

🧪 Testing
✅ Unit-тесты (контроллеры, сервисы)

✅ Интеграционные тесты API

✅ Widget-тесты для Flutter-интерфейса

✅ Покрытие login flow, feed, match

# Backend
```bach
go test ./...
```

# Flutter
flutter test
🧩 Features Summary
Функция	✅ Статус
Telegram Auth	✅
JWT авторизация	✅
Создание/редактирование профиля	✅
Лента подходящих пользователей	✅
Лайки и совпадения	✅
Анимация на лайк	✅
Поддержка RU/EN	✅
Тёмная и светлая тема	✅
Docker Compose	✅
CI/CD (GitHub Actions / Vercel)	✅

## 🗃️ Database Schema
📊 Ниже представлена схема базы данных:

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
🗂️ GitHub Projects для трекинга задач

✅ Регулярные meaningful pull-requests

## 👥 Роли в команде:

Backend (Go): Ilyina Maria
Frontend (Flutter): Daria Nikolaeva
