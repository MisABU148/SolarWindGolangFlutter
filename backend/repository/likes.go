package repository

import (
	"backend/model"
	"database/sql"
	"errors"
)

type LikesRepository struct {
	DB *sql.DB
}

func (r *LikesRepository) FindByID(id model.LikesCompositePrimaryKey) (*model.LikesEntity, error) {
	var entity model.LikesEntity
	entity.ID = id

	row := r.DB.QueryRow(`
		SELECT is_first_likes, is_second_likes
		FROM likes
		WHERE liker_id = $1 AND liked_id = $2
	`, id.LikerID, id.LikedID)

	var firstLikes, secondLikes sql.NullBool
	err := row.Scan(&firstLikes, &secondLikes)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if firstLikes.Valid {
		entity.IsFirstLikes = &firstLikes.Bool
	}
	if secondLikes.Valid {
		entity.IsSecondLikes = &secondLikes.Bool
	}
	return &entity, nil
}

func (r *LikesRepository) Save(entity *model.LikesEntity) error {
	// Upsert — insert or update
	_, err := r.DB.Exec(`
		INSERT INTO likes (liker_id, liked_id, is_first_likes, is_second_likes)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (liker_id, liked_id) DO UPDATE SET
			is_first_likes = EXCLUDED.is_first_likes,
			is_second_likes = EXCLUDED.is_second_likes
	`,
		entity.ID.LikerID,
		entity.ID.LikedID,
		entity.IsFirstLikes,
		entity.IsSecondLikes,
	)
	return err
}
