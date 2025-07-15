package repository

import (
	"backend/model"
	"database/sql"
)

type SportRepository struct {
	DB *sql.DB
}

func (r *SportRepository) GetAllSports() ([]model.Sport, error) {
	rows, err := r.DB.Query("SELECT id, name FROM sports")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sports []model.Sport
	for rows.Next() {
		var s model.Sport
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		sports = append(sports, s)
	}
	return sports, nil
}

func (r *SportRepository) GetSportsPaginated(page, size int) ([]model.Sport, error) {
	offset := (page - 1) * size
	rows, err := r.DB.Query("SELECT id, name FROM sports LIMIT $1 OFFSET $2", size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sports []model.Sport
	for rows.Next() {
		var s model.Sport
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		sports = append(sports, s)
	}
	return sports, nil
}

func (r *SportRepository) SearchSports(word string) ([]model.Sport, error) {
	rows, err := r.DB.Query("SELECT id, name FROM sports WHERE name LIKE $1", "%"+word+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sports []model.Sport
	for rows.Next() {
		var s model.Sport
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		sports = append(sports, s)
	}
	return sports, nil
}
