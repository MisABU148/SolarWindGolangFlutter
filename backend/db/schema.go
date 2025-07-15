package db

import (
	"database/sql"
	"log"
)

func InitSchema(db *sql.DB) error {
	// Используем многострочные строки с обратными кавычками для лучшей читаемости
	createTables := []string{
		`
		CREATE TABLE IF NOT EXISTS cities (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		)`,
		`
		CREATE TABLE IF NOT EXISTS sports (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE
		)`,
		`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT,
			telegram_id INTEGER UNIQUE,
			alias TEXT,
			description TEXT,
			age INTEGER CHECK (age >= 18 AND age <= 120),
			preferred_gender TEXT,
			gender TEXT CHECK (gender IN ('male', 'female', 'other')),
			city_id INTEGER REFERENCES cities(id),
			preferred_gym_time TEXT,
			sport_id INTEGER REFERENCES sports(id),
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	// Добавляем индексы для часто используемых полей
	createIndexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_users_telegram_id ON users(telegram_id)",
		"CREATE INDEX IF NOT EXISTS idx_users_city_id ON users(city_id)",
		"CREATE INDEX IF NOT EXISTS idx_users_sport_id ON users(sport_id)",
	}

	// Выполняем все запросы в транзакции
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Создаем таблицы
	for _, stmt := range createTables {
		if _, err = tx.Exec(stmt); err != nil {
			log.Printf("Error creating table: %v\nQuery: %s", err, stmt)
			return err
		}
	}

	// Создаем индексы
	for _, stmt := range createIndexes {
		if _, err = tx.Exec(stmt); err != nil {
			log.Printf("Error creating index: %v\nQuery: %s", err, stmt)
			return err
		}
	}

	return nil
}
