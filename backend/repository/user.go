package repository

import (
	"backend/constant"
	"backend/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user model.User) (int64, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	fmt.Print("start creating \n")

	stmt := `
		INSERT INTO users (telegram_id, username, description, age, gender, preferred_gender, city_id, preferred_gym_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var userID int64
	err = tx.QueryRow(stmt,
		user.TelegramID,
		user.UserName,
		user.Description,
		user.Age,
		user.Gender,
		user.PreferredGender,
		user.CityID,
		user.PreferredGymTime,
	).Scan(&userID)

	fmt.Printf("make insert user %s\n", err)
	if err != nil {
		return 0, err
	}

	for _, sportID := range user.SportIDs {
		_, err = tx.Exec(`INSERT INTO user_sport (user_id, sport_id) VALUES ($1, $2)`, userID, sportID)
		if err != nil {
			return 0, err
		}
	}
	fmt.Print("make insert sport \n")

	return userID, tx.Commit()
}

func (r *UserRepository) GetUserByID(id int64) (*model.User, error) {
	var user model.User

	err := r.DB.QueryRow(`
		SELECT id, telegram_id, username, description, age, gender, preferred_gender, city_id, preferred_gym_time
		FROM users WHERE id = $1
	`, id).Scan(
		&user.ID,
		&user.TelegramID,
		&user.UserName,
		&user.Description,
		&user.Age,
		&user.Gender,
		&user.PreferredGender,
		&user.CityID,
		&user.PreferredGymTime,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) DeleteUserByID(id int64) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *UserRepository) UpdateUser(user model.User) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE users SET
			telegram_id = $1,
			username = $2,
			description = $3,
			age = $4,
			gender = $5,
			preferred_gender = $6,
			city_id = $7,
			preferred_gym_time = $8
		WHERE id = $9
	`,
		user.TelegramID,
		user.UserName,
		user.Description,
		user.Age,
		user.Gender,
		user.PreferredGender,
		user.CityID,
		user.PreferredGymTime,
		user.ID,
	)
	fmt.Printf("update user error %s", err)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM user_sport WHERE user_id = $1`, user.ID)
	if err != nil {
		return err
	}

	for _, sportID := range user.SportIDs {
		_, err = tx.Exec(`INSERT INTO user_sport (user_id, sport_id) VALUES ($1, $2)`, user.ID, sportID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := r.DB.Query(`
		SELECT id, telegram_id, username, description, age, gender, preferred_gender, city_id, preferred_gym_time
		FROM users
	`)
	fmt.Printf("getall users error %s", err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.TelegramID,
			&user.UserName,
			&user.Description,
			&user.Age,
			&user.Gender,
			&user.PreferredGender,
			&user.CityID,
			&user.PreferredGymTime,
		)
		fmt.Printf("get all users error %s", err)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User

	err := r.DB.QueryRow(`
		SELECT id, telegram_id, username, description, age, gender, preferred_gender, city_id, preferred_gym_time
		FROM users WHERE username = $1
	`, username).Scan(
		&user.ID,
		&user.TelegramID,
		&user.UserName,
		&user.Description,
		&user.Age,
		&user.Gender,
		&user.PreferredGender,
		&user.CityID,
		&user.PreferredGymTime,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateDeckAllSettings(
	initialUserID int64,
	gender string,
	preferredGender string,
	cityID int64,
	age int64,
) ([]model.User, error) {

	rows, err := r.DB.Query(constant.CreateDeckQuery,
		cityID, age, gender, preferredGender, initialUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		err := rows.Scan(
			&u.ID,
			&u.TelegramID,
			&u.UserName,
			&u.Description,
			&u.Age,
			&u.Gender,
			&u.PreferredGender,
			&u.CityID,
			&u.PreferredGymTime,
			new(int), // matched_sports_count
			new(int), // matched_days_count
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
