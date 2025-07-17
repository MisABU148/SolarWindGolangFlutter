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
			name TEXT
		)`,
		`
		CREATE TABLE IF NOT EXISTS sports (
			id SERIAL PRIMARY KEY,
			name TEXT
		)`,
		`
		CREATE TABLE IF NOT EXISTS users (
		id BIGINT PRIMARY KEY,
		telegram_id BIGINT UNIQUE,
		username TEXT UNIQUE,
		description TEXT,
		age INTEGER,
		gender TEXT CHECK (gender IN ('male', 'female', 'other')),
		preferred_gender TEXT CHECK (gender IN ('male', 'female', 'other')),
		city_id INTEGER REFERENCES cities(id),
		preferred_gym_time INTEGER,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	)`,
		`
		CREATE TABLE IF NOT EXISTS user_sport (
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		sport_id INTEGER REFERENCES sports(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, sport_id)
	)`,
		`	CREATE TABLE IF NOT EXISTS likes (
		liker_id BIGINT NOT NULL,
		liked_id BIGINT NOT NULL,
		is_first_likes BOOLEAN NOT NULL,
		is_second_likes BOOLEAN,
		PRIMARY KEY (liker_id, liked_id)
	)`,
	}

	createIndexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_users_telegram_id ON users(telegram_id)",
		"CREATE INDEX IF NOT EXISTS idx_users_city_id ON users(city_id)",
		"CREATE INDEX IF NOT EXISTS idx_user_sport_user_id ON user_sport(user_id)",
		"CREATE INDEX IF NOT EXISTS idx_user_sport_sport_id ON user_sport(sport_id)",
	}

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
