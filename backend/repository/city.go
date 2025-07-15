package repository

import (
	"backend/model"
	"database/sql"
)

type CityRepository struct {
	DB *sql.DB
}

func (r *CityRepository) GetAllCities() ([]model.City, error) {
	rows, err := r.DB.Query("SELECT id, name FROM cities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []model.City
	for rows.Next() {
		var c model.City
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}
	return cities, nil
}

func (r *CityRepository) GetCitiesPaginated(page, size int) ([]model.City, error) {
	offset := (page - 1) * size
	rows, err := r.DB.Query("SELECT id, name FROM cities LIMIT $1 OFFSET $2", size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []model.City
	for rows.Next() {
		var c model.City
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}
	return cities, nil
}

func (r *CityRepository) SearchCities(word string) ([]model.City, error) {
	rows, err := r.DB.Query("SELECT id, name FROM cities WHERE name LIKE $1", "%"+word+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []model.City
	for rows.Next() {
		var c model.City
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}
	return cities, nil
}
