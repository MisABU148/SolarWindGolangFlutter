package repository

import (
	"backend/model"
	"database/sql"
	"fmt"
)

type LikesRepository struct {
	DB *sql.DB
}

func (r *LikesRepository) SaveLike(like model.Like) error {
	query := `
		INSERT INTO likes (liker_id, liked_id, is_first_likes, is_second_likes)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (liker_id, liked_id)
		DO UPDATE SET is_first_likes = $3, is_second_likes = $4
	`

	_, err := r.DB.Exec(query,
		like.LikerID,
		like.LikedID,
		like.IsFirstLikes,
		like.IsSecondLikes,
	)
	return err
}

func (r *LikesRepository) FindLike(likerID, likedID int64) (*model.Like, error) {
	query := `
		SELECT liker_id, liked_id, is_first_likes, is_second_likes
		FROM likes
		WHERE liker_id = $1 AND liked_id = $2
	`

	row := r.DB.QueryRow(query, likerID, likedID)

	var like model.Like
	var isSecondLikes sql.NullBool

	err := row.Scan(
		&like.LikerID,
		&like.LikedID,
		&like.IsFirstLikes,
		&isSecondLikes,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if isSecondLikes.Valid {
		like.IsSecondLikes = &isSecondLikes.Bool
	} else {
		like.IsSecondLikes = nil
	}

	return &like, nil
}

func (r *LikesRepository) GetMutualLikes(userID int64) ([]int64, error) {
	query := `
        SELECT liker_id AS user_id
        FROM likes
        WHERE liked_id = $1 AND is_second_likes = true
        UNION
        SELECT liked_id AS user_id
        FROM likes
        WHERE liker_id = $1 AND is_first_likes = true
    `

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query liked users: %w", err)
	}
	defer rows.Close()
	fmt.Print(rows)

	var userIDs []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan user_id: %w", err)
		}
		userIDs = append(userIDs, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return userIDs, nil
}
