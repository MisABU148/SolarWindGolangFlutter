package service

import (
	"backend/model"
	"backend/repository"
)

type SportService struct {
	Repo *repository.SportRepository
}

func (s *SportService) GetSports() ([]model.Sport, error) {
	return s.Repo.GetAllSports()
}

func (s *SportService) GetPaginatedSports(page, size int) ([]model.Sport, error) {
	return s.Repo.GetSportsPaginated(page, size)
}

func (s *SportService) SearchSports(word string) ([]model.Sport, error) {
	return s.Repo.SearchSports(word)
}
