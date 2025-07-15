package service

import (
	"backend/model"
	"backend/repository"
	"strconv"
)

type UserService struct {
	Repo *repository.UserRepository
}

func intSliceToStringSlice(input []int) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = strconv.Itoa(v)
	}
	return result
}

func (s *UserService) CreateUser(dto model.UserDTO) (int64, error) {
	user := model.User{
		UserName:         dto.Username,
		Alias:            dto.Username,
		Description:      dto.Description,
		Age:              dto.Age,
		Gender:           dto.Gender,
		CityId:           dto.CityId,
		PreferredGymTime: intSliceToStringSlice(dto.PreferredGymTime),
		SportId:          intSliceToStringSlice(dto.SportId),
	}

	return s.Repo.CreateUser(user)
}

func (s *UserService) GetByUserID(id int64) (*model.UserDTO, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return s.ToDTO(*user), nil
}

func (s *UserService) DeleteUserByID(id int64) error {
	return s.Repo.DeleteUserByID(id)
}

func (s *UserService) UpdateUser(dto model.UserDTO) error {
	user := s.FromDTO(dto)
	user.ID = dto.ID
	return s.Repo.UpdateUser(user)
}

func (s *UserService) GetUsers() ([]model.UserDTO, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var result []model.UserDTO
	for _, u := range users {
		result = append(result, *s.ToDTO(u))
	}
	return result, nil
}
