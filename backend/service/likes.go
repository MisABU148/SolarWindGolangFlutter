package service

import (
	"backend/model"
	"backend/notifier"
	"backend/repository"
	"fmt"
)

type LikesService struct {
	Repo     *repository.LikesRepository
	Notifier notifier.MatchNotifier
}

func (s *LikesService) SaveOrUpdateLike(likerID, likedID int64, isLiked bool) error {
	// Попытка найти существующую запись в прямом и обратном направлении
	forward, _ := s.Repo.FindLike(likerID, likedID)
	reverse, _ := s.Repo.FindLike(likedID, likerID)

	var like model.Like

	switch {
	case forward != nil:
		// Пользователь уже ставил лайк, обновляем его решение
		like = *forward
		like.IsFirstLikes = isLiked

	case reverse != nil:
		// Лайк уже есть от другого пользователя, добавляем ответ
		like = *reverse
		like.IsSecondLikes = &isLiked

	default:
		// Новый лайк, создаём новую запись
		like = model.Like{
			LikerID:       likerID,
			LikedID:       likedID,
			IsFirstLikes:  isLiked,
			IsSecondLikes: nil,
		}
	}

	// Сохраняем/обновляем
	if err := s.Repo.SaveLike(like); err != nil {
		return err
	}

	// Проверка на взаимный лайк
	if like.IsSecondLikes != nil && like.IsFirstLikes && *like.IsSecondLikes {
		if err := s.Notifier.NotifyMatch(likerID, likedID); err != nil {
			// Можно логировать ошибку, если важно
			return fmt.Errorf("match notification failed: %w", err)
		}
	}

	return nil
}
