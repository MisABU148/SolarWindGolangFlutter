package repository

import (
	"backend/model"
	"database/sql"

	"github.com/lib/pq"
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

func (r *SportRepository) GetSportByIds(ids []int64) ([]string, error) {
	rows, err := r.DB.Query("SELECT name FROM sports WHERE id = ANY($1)", pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return names, nil
}

func (r *SportRepository) GetSportByUserId(userID int64) ([]string, error) {
	query := `
		SELECT s.name
		FROM user_sport us
		JOIN sports s ON us.sport_id = s.id
		WHERE us.user_id = $1
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string

	for rows.Next() {
		var sportName string
		if err := rows.Scan(&sportName); err != nil {
			return nil, err
		}
		result = append(result, sportName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
