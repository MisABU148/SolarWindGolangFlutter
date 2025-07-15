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
		INSERT INTO users(username, password_hash, alias, description, age, preferredgender, gender, city_id, preferred_gym_time, sport_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var userID int64
	err = stmt.QueryRow(
		user.UserName,
		user.PasswordHash,
		user.Alias,
		user.Description,
		user.Age,
		user.PreferredGender,
		user.Gender,
		user.CityId,
		string(preferredTimeJSON),
		string(sportsJSON),
	).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *UserRepository) GetUserByID(id int64) (*model.User, error) {
	row := r.DB.QueryRow(`
		SELECT id, username, password_hash, alias, description, age, preferredgender, gender, city_id, preferred_gym_time, sport_id
		FROM users
		WHERE id = $1
	`, id)

	var user model.User
	var gymTimeJSON, sportsJSON string

	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.PasswordHash,
		&user.Alias,
		&user.Description,
		&user.Age,
		&user.PreferredGender,
		&user.Gender,
		&user.CityId,
		&gymTimeJSON,
		&sportsJSON,
	)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal([]byte(gymTimeJSON), &user.PreferredGymTime)
	_ = json.Unmarshal([]byte(sportsJSON), &user.SportId)

	return &user, nil
}

func (r *UserRepository) DeleteUserByID(id int64) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *UserRepository) UpdateUser(user model.User) error {
	gymTimeJSON, _ := json.Marshal(user.PreferredGymTime)
	sportsJSON, _ := json.Marshal(user.SportId)

	_, err := r.DB.Exec(`
		UPDATE users SET
			username = $1, password_hash = $2, alias = $3, description = $4, age = $5,
			preferredgender = $6, gender = $7, city_id = $8, preferred_gym_time = $9, sport_id = $10
		WHERE id = $11
	`,
		user.UserName,
		user.PasswordHash,
		user.Alias,
		user.Description,
		user.Age,
		user.PreferredGender,
		user.Gender,
		user.CityId,
		string(gymTimeJSON),
		string(sportsJSON),
		user.ID,
	)
	return err
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := r.DB.Query(`
		SELECT id, username, password_hash, alias, description, age, preferredgender, gender, city_id, preferred_gym_time, sport_id
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		var gymTimeJSON, sportsJSON string
		err := rows.Scan(
			&user.ID,
			&user.UserName,
			&user.PasswordHash,
			&user.Alias,
			&user.Description,
			&user.Age,
			&user.PreferredGender,
			&user.Gender,
			&user.CityId,
			&gymTimeJSON,
			&sportsJSON,
		)
		if err != nil {
			continue
		}
		_ = json.Unmarshal([]byte(gymTimeJSON), &user.PreferredGymTime)
		_ = json.Unmarshal([]byte(sportsJSON), &user.SportId)
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	row := r.DB.QueryRow(`
		SELECT id, username, password_hash, alias, description, age, preferredgender, gender, city_id, preferred_gym_time, sport_id
		FROM users
		WHERE username = $1
	`, username)

	var user model.User
	var gymTimeJSON, sportsJSON string

	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.PasswordHash,
		&user.Alias,
		&user.Description,
		&user.Age,
		&user.PreferredGender,
		&user.Gender,
		&user.CityId,
		&gymTimeJSON,
		&sportsJSON,
	)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal([]byte(gymTimeJSON), &user.PreferredGymTime)
	_ = json.Unmarshal([]byte(sportsJSON), &user.SportId)

	return &user, nil
}
