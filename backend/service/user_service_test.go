package service

import (
	"backend/model"

	"github.com/stretchr/testify/mock"
)

// ✅ Mock, реализующий UserRepositoryInterface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user model.User) (int64, error) {
	args := m.Called(user)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int64) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetAllUsers() ([]model.UserDTO, error) {
	args := m.Called()
	return args.Get(0).([]model.UserDTO), args.Error(1)
}

func (m *MockUserRepository) DeleteUserByID(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) CreateDeckAllSettings(initialID int64, gender, preferredGender string, cityID, age int64) ([]model.User, error) {
	args := m.Called(initialID, gender, preferredGender, cityID, age)
	return args.Get(0).([]model.User), args.Error(1)
}
