package service

import (
	"backend/mapper"
	"backend/model"
	"backend/repository"
	"fmt"
)

type UserService struct {
	Repo      *repository.UserRepository
	CityRepo  *repository.CityRepository
	SportRepo *repository.SportRepository
	LikesRepo *repository.LikesRepository
}

func (s *UserService) CreateUser(dto model.UserDTO) (int64, error) {
	user := mapper.MapToUserEntity(dto)
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetByUserID(id int64) (model.ProfileDTO, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return model.ProfileDTO{}, err
	}

	profileDTO, err := s.MapToProfileDTO(*user)
	if err != nil {
		return model.ProfileDTO{}, err
	}

	return profileDTO, nil
}

func (s *UserService) UpdateUser(dto model.UserDTO) error {
	fmt.Print("update \n")
	fmt.Print(dto)
	user := mapper.MapToUserEntity(dto)
	return s.Repo.UpdateUser(user)
}

func (s *UserService) GetUsers() ([]model.ProfileDTO, error) {
	if s.CityRepo == nil {
		fmt.Println("CityRepo is nil!")
	}
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userDTOs []model.ProfileDTO
	fmt.Printf("get users from db \n")
	for _, user := range users {
		dto, err := s.MapToProfileDTO(user)
		fmt.Printf("start map \n")
		if err != nil {
			return nil, fmt.Errorf("failed to map user %d: %w", user.ID, err)
		}
		userDTOs = append(userDTOs, dto)
	}

	return userDTOs, nil
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

func (s *UserService) CreateDeckForUser(initialID int64) ([]model.ProfileDTO, error) {
	user, err := s.Repo.GetUserByID(initialID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	users, err := s.Repo.CreateDeckAllSettings(
		initialID,
		user.Gender,
		user.PreferredGender,
		user.CityID,
		user.Age,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create deck: %w", err)
	}

	var userDTOs []model.ProfileDTO
	for _, u := range users {
		dto, err := s.MapToProfileDTO(u)
		if err != nil {
			return nil, fmt.Errorf("failed to map user %d: %w", u.ID, err)
		}
		userDTOs = append(userDTOs, dto)
	}

	return userDTOs, nil
}

func (s *UserService) GetMutualMatches(id int64) ([]model.ProfileDTO, error) {
	likes, err := s.LikesRepo.GetMutualLikes(id)
	if err != nil {
		return nil, err
	}

	var profiles []model.ProfileDTO
	for userID := range likes {
		profile, err := s.GetByUserID(int64(userID))
		if err != nil {
			continue
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}
