package service

import (
	"backend/model"
	"strconv"
)

func (s *UserService) FromDTO(dto model.UserDTO) model.User {
	var gymTimeStr []string
	for _, t := range dto.PreferredGymTime {
		gymTimeStr = append(gymTimeStr, strconv.Itoa(t))
	}

	var sportsStr []string
	for _, s := range dto.SportId {
		sportsStr = append(sportsStr, strconv.Itoa(s))
	}

	return model.User{
		ID:               dto.ID,
		UserName:         dto.Username,
		Alias:            dto.Username,
		Description:      dto.Description,
		Age:              dto.Age,
		Gender:           dto.Gender,
		CityId:           dto.CityId,
		PreferredGymTime: gymTimeStr,
		SportId:          sportsStr,
	}
}

func (s *UserService) ToDTO(user model.User) *model.UserDTO {
	var gymTimeInts []int
	for _, t := range user.PreferredGymTime {
		if val, err := strconv.Atoi(t); err == nil {
			gymTimeInts = append(gymTimeInts, val)
		}
	}

	var sportInts []int
	for _, s := range user.SportId {
		if val, err := strconv.Atoi(s); err == nil {
			sportInts = append(sportInts, val)
		}
	}

	return &model.UserDTO{
		ID:               user.ID,
		Username:         user.UserName,
		Description:      user.Description,
		Age:              user.Age,
		Gender:           user.Gender,
		CityId:           user.CityId,
		PreferredGymTime: gymTimeInts,
		SportId:          sportInts,
	}
}
