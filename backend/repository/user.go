package repository

import (
	"backend/model"
	"database/sql"
	"encoding/json"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user model.User) (int64, error) {
	preferredTimeJSON, err := json.Marshal(user.PreferredGymTime)
	if err != nil {
		return 0, err
	}

	sportsJSON, err := json.Marshal(user.SportId)
	if err != nil {
		return 0, err
	}

	stmt, err := r.DB.Prepare(`
		INSERT INTO users(username, alias, description, age, gender, city_id, preferred_gym_time, sport_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		user.UserName,
		user.Alias,
		user.Description,
		user.Age,
		user.Gender,
		user.CityId,
		string(preferredTimeJSON),
		string(sportsJSON),
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *UserRepository) GetUserByID(id int64) (*model.User, error) {
	row := r.DB.QueryRow("SELECT id, username, alias, description, age, gender, city_id, preferred_gym_time, sport_id FROM users WHERE id = ?", id)
	var user model.User
	var gymTimeJSON, sportsJSON string

	err := row.Scan(&user.ID, &user.UserName, &user.Alias, &user.Description, &user.Age, &user.Gender, &user.CityId, &gymTimeJSON, &sportsJSON)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal([]byte(gymTimeJSON), &user.PreferredGymTime)
	_ = json.Unmarshal([]byte(sportsJSON), &user.SportId)

	return &user, nil
}

func (r *UserRepository) DeleteUserByID(id int64) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func (r *UserRepository) UpdateUser(user model.User) error {
	gymTimeJSON, _ := json.Marshal(user.PreferredGymTime)
	sportsJSON, _ := json.Marshal(user.SportId)

	_, err := r.DB.Exec(`
		UPDATE users SET username = ?, alias = ?, description = ?, age = ?, gender = ?, city_id = ?, preferred_gym_time = ?, sport_id = ?
		WHERE id = ?
	`,
		user.UserName, user.Alias, user.Description, user.Age, user.Gender, user.CityId, string(gymTimeJSON), string(sportsJSON), user.ID)
	return err
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := r.DB.Query("SELECT id, username, alias, description, age, gender, city_id, preferred_gym_time, sport_id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		var gymTimeJSON, sportsJSON string
		err := rows.Scan(&user.ID, &user.UserName, &user.Alias, &user.Description, &user.Age, &user.Gender, &user.CityId, &gymTimeJSON, &sportsJSON)
		if err != nil {
			continue
		}
		_ = json.Unmarshal([]byte(gymTimeJSON), &user.PreferredGymTime)
		_ = json.Unmarshal([]byte(sportsJSON), &user.SportId)
		users = append(users, user)
	}
	return users, nil
}
