package service

import "backend/model"

type UserServiceInterface interface {
	CreateUser(dto model.UserDTO) (int64, error)
	GetByUserID(id int64) (model.UserDTO, error)
	DeleteUserByID(id int64) error
	UpdateUser(dto model.UserDTO) error
	GetUsers() ([]model.UserDTO, error)
	CreateDeckForUser(id int64) ([]model.UserDTO, error)
}
