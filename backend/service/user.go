package service

import (
	"backend/mapper"
	"backend/model"
	"backend/repository"
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
	result := make([]model.UserDTO, len(users))
	return result, nil
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
