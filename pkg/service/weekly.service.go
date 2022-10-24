package service

import (
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
)

func (s *service) FetchWeeklies() (*[]presenter.Weekly, error) {
	return s.repository.ReadWeeklies()
}

func (s *service) InsertWeekly(weekly *entities.Weekly) (*entities.Weekly, error) {
	return s.repository.CreateWeekly(weekly)
}
func (s *service) UpdateWeekly(weekly *entities.Weekly) (*entities.Weekly, error) {
	return s.repository.UpdateWeekly(weekly)
}

func (s *service) DeleteWeekly(ID string) error {
	return s.repository.DeleteWeekly(ID)
}
