package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

type City struct {
	ID   int
	Name string
}

type Sport struct {
	ID   int
	Name string
}

func SeedInitialData(db *sql.DB, cityFile, sportFile string) error {
	if err := insertCities(db, cityFile); err != nil {
		return err
	}

	if err := insertSports(db, sportFile); err != nil {
		return err
	}

	return nil
}

func insertCities(db *sql.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var cityNames []string
	if err := json.Unmarshal(data, &cityNames); err != nil {
		return err
	}

	for i, name := range cityNames {
		_, err := db.Exec(`
			INSERT INTO cities(id, name) 
			VALUES ($1, $2)
			ON CONFLICT (id) DO NOTHING`,
			i+1, name)
		if err != nil {
			log.Println("Failed to insert city:", name, err)
		}
	}
	return nil
}

func insertSports(db *sql.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var sportNames []string
	if err := json.Unmarshal(data, &sportNames); err != nil {
		return err
	}

	for i, name := range sportNames {
		_, err := db.Exec(`
			INSERT INTO sports(id, name) 
			VALUES ($1, $2)
			ON CONFLICT (id) DO NOTHING`,
			i+1, name)
		if err != nil {
			log.Println("Failed to insert sport:", name, err)
		}
	}
	return nil
}
