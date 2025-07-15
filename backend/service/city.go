package service

import (
	"backend/model"
	"backend/repository"
)

type CityService struct {
	Repo *repository.CityRepository
}

func (s *CityService) GetCities() ([]model.City, error) {
	return s.Repo.GetAllCities()
}

func (s *CityService) GetPaginatedCities(page, size int) ([]model.City, error) {
	return s.Repo.GetCitiesPaginated(page, size)
}

func (s *CityService) SearchCities(word string) ([]model.City, error) {
	return s.Repo.SearchCities(word)
}
