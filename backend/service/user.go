package service

import (
	"backend/mapper"
	"backend/model"
	"backend/repository"
	"fmt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) CreateUser(dto model.UserDTO) (int64, error) {
	user := mapper.MapToUserEntity(dto)
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetByUserID(id int64) (model.UserDTO, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return model.UserDTO{}, err
	}
	return *user, nil
}

func (s *UserService) UpdateUser(dto model.UserDTO) error {
	user := mapper.MapToUserEntity(dto)
	return s.Repo.UpdateUser(user)
}

func (s *UserService) GetUsers() ([]model.UserDTO, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) DeleteUserByID(id int64) error {
	return s.Repo.DeleteUserByID(id)
}

func (s *UserService) GetUserByUsername(username string) (model.UserDTO, error) {
	user, err := s.Repo.GetUserByUsername(username)
	if err != nil {
		return model.UserDTO{}, err
	}
	return mapper.MapToUserDTO(*user), nil
}

func (s *UserService) CreateDeckForUser(initialID int64) ([]model.UserDTO, error) {
	// Получаем параметры пользователя по initialID
	user, err := s.Repo.GetUserByID(initialID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Получаем список пользователей по параметрам
	users, err := s.Repo.CreateDeckAllSettings(
		initialID,
		user.Gender,
		user.PreferredGender,
		user.CityID,
		int64(user.Age), // Если Age в model.User int, а здесь int64 — явно приводим
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create deck: %w", err)
	}

	// Преобразуем []model.User в []model.UserDTO (если нужно)
	var userDTOs []model.UserDTO
	for _, u := range users {
		userDTOs = append(userDTOs, mapper.MapToUserDTO(u)) // Предполагается, что есть такой маппер
	}

	return userDTOs, nil
}
