package service

import (
	"backend/mapper"
	"backend/model"
	"backend/repository"
	"context"
	"fmt"
)

type MatchNotifier interface {
	NotifyMatch(likerID, likedID int64) error
}

type LikesService struct {
	Repo     *repository.LikesRepository
	Notifier MatchNotifier
}

func (s *LikesService) SaveOrUpdateDecision(ctx context.Context, dto *model.LikeDTO) error {
	// Найдём обратную запись (reverse)
	reverseID := model.LikesCompositePrimaryKey{LikerID: dto.LikedID, LikedID: dto.LikerID}
	reverseEntity, err := s.Repo.FindByID(reverseID)
	if err != nil {
		return err
	}

	// Найдём прямую запись (straight)
	straightID := model.LikesCompositePrimaryKey{LikerID: dto.LikerID, LikedID: dto.LikedID}
	straightEntity, err := s.Repo.FindByID(straightID)
	if err != nil {
		return err
	}

	var decisionEntity *model.LikesEntity

	if straightEntity != nil {
		decisionEntity = straightEntity
	} else if reverseEntity != nil {
		decisionEntity = reverseEntity
	} else {
		// Нет ни прямой, ни обратной записи — создаём новую
		entity := mapper.MapToLikesEntity(dto)
		err = s.Repo.Save(entity)
		return err
	}

	// Обновляем поля в существующей записи:
	if decisionEntity.IsSecondLikes != nil {
		decisionEntity.IsFirstLikes = &dto.IsLiker
	} else {
		decisionEntity.IsSecondLikes = &dto.IsLiker
	}

	err = s.Repo.Save(decisionEntity)
	if err != nil {
		return err
	}

	// Проверяем на матч
	if decisionEntity.IsFirstLikes != nil && decisionEntity.IsSecondLikes != nil &&
		*decisionEntity.IsFirstLikes && *decisionEntity.IsSecondLikes {
		// Есть матч!
		err = s.Notifier.NotifyMatch(dto.LikerID, dto.LikedID)
		if err != nil {
			fmt.Printf("NotifyMatch error: %v\n", err)
		}
	}

	return nil
}
