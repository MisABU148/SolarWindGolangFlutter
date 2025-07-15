package db

import (
	"database/sql"
	"log"
)

func InitSchema(db *sql.DB) error {
	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		alias TEXT,
		description TEXT,
		age INTEGER,
		gender TEXT,
		city_id INTEGER,
		preferred_gym_time TEXT,
		sport_id TEXT
	);`

	createCities := `
	CREATE TABLE IF NOT EXISTS cities (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);`

	createSports := `
	CREATE TABLE IF NOT EXISTS sports (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);`

	statements := []string{createUsers, createCities, createSports}
	for _, stmt := range statements {
		if _, err := db.Exec(stmt); err != nil {
			log.Printf("Error executing statement: %s", stmt)
			return err
		}
	}

	return nil
}
