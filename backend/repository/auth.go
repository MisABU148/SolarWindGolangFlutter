package repository

import (
	"backend/model"
	"fmt"
)

func (r *UserRepository) GetByTelegramID(tgID int64) (*model.UserDTO, error) {
	row := r.DB.QueryRow(`
		SELECT id
		FROM users
		WHERE telegram_id = $1
	`, tgID)

	var user model.UserDTO

	err := row.Scan(
		&user.ID,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) getSportsByUserID(userID int64) ([]int, error) {
	rows, err := r.DB.Query(`SELECT sport_id FROM user_sport WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sports []int
	for rows.Next() {
		var sportID int
		if err := rows.Scan(&sportID); err != nil {
			continue
		}
		sports = append(sports, sportID)
	}

	return sports, nil
}

func (r *UserRepository) CreateTelegramUser(user model.User) (int64, error) {
	fmt.Print("creating user")
	// используем TelegramID как ID
	_, err := r.DB.Exec(`
        INSERT INTO users (
            id,
            username,
            telegram_id,
            description,
            age,
            gender,
            preferred_gender,
            city_id,
            preferred_gym_time
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `,
		user.TelegramID,
		user.UserName,
		user.TelegramID,
		user.Description,
		user.Age,
		"male",
		"male",
		1,
		user.PreferredGymTime,
	)

	if err != nil {
		return 0, fmt.Errorf("failed to create telegram user: %w", err)
	}

	return user.TelegramID, nil
}
