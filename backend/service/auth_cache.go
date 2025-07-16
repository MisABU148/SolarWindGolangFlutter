package service

import (
	"backend/model"
	"backend/repository"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type AuthCache struct {
	data map[int64]int64 // code -> telegramID
	mu   sync.Mutex
	repo *repository.UserRepository
}

type UserResult struct {
	UserID    int64
	IsNewUser bool
}

func NewAuthCache(repo *repository.UserRepository) *AuthCache {
	return &AuthCache{
		data: make(map[int64]int64),
		repo: repo,
	}
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (a *AuthCache) SetCode(code int64, telegramID int64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.data[code] = telegramID

	// Автоудаление через 60 секунд
	go func() {
		time.Sleep(60 * time.Second)
		a.mu.Lock()
		delete(a.data, code)
		a.mu.Unlock()
	}()
}

func (a *AuthCache) GetUserIDByCode(code int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	id, ok := a.data[code]
	if ok {
		delete(a.data, code)
	}
	return id, ok
}

func (s *UserService) CreateOrGetTelegramUser(tgID int64) (*UserResult, error) {
	user, err := s.Repo.GetByTelegramID(tgID)
	fmt.Print("1")
	if err == nil {
		return &UserResult{
			UserID:    user.ID,
			IsNewUser: false,
		}, nil
	}

	userID, err := s.Repo.CreateTelegramUser(model.User{
		TelegramID: tgID,
		UserName:   "tg_" + strconv.FormatInt(tgID, 10),
	})
	if err != nil {
		return nil, err
	}

	return &UserResult{
		UserID:    userID,
		IsNewUser: true,
	}, nil
}
