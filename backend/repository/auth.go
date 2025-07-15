package repository

import (
	"backend/model"
	"encoding/json"
)

func (r *UserRepository) GetByTelegramID(tgID int64) (*model.User, error) {
	row := r.DB.QueryRow("SELECT id, username, alias, telegram_id FROM users WHERE telegram_id = ?", tgID)
	var user model.User
	err := row.Scan(&user.ID, &user.UserName, &user.Alias, &user.TelegramID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateTelegramUser(user model.User) (int64, error) {
	gymTimeJSON, err := json.Marshal(user.PreferredGymTime)
	if err != nil {
		return 0, err
	}

	sportIDJSON, err := json.Marshal(user.SportId)
	if err != nil {
		return 0, err
	}

	stmt, err := r.DB.Prepare(`
		INSERT INTO users (
			username,
			password_hash,
			alias,
			description,
			age,
			preferred_gender,
			gender,
			city_id,
			preferred_gym_time,
			sport_id,
			telegram_id
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
		string(gymTimeJSON),
		string(sportIDJSON),
		user.TelegramID,
	).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}
