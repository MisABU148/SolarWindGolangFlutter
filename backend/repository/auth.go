package repository

import (
	"backend/model"
	"fmt"
)

func (r *UserRepository) GetByTelegramID(tgID int64) (*model.User, error) {
	row := r.DB.QueryRow("SELECT id, username, alias, telegram_id FROM users WHERE telegram_id = $1", tgID)
	var user model.User
	err := row.Scan(&user.ID, &user.UserName, &user.Alias, &user.TelegramID)
	fmt.Print(err)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateTelegramUser(user model.User) (int64, error) {
	var id int64

	err := r.DB.QueryRow(`
        INSERT INTO users (
            username,
            alias,
            telegram_id
        )
        VALUES ($1, $2, $3)
        RETURNING id
    `,
		user.UserName,
		user.Alias,
		user.TelegramID,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create telegram user: %w", err)
	}

	return id, nil
}
