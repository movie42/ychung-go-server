package service

import (
	"github.com/movie42/ychung-go-server/presenter"
)

func (s *service) FetchWeeklies() (*[]presenter.Weekly, error) {
	return s.repository.ReadWeeklies()
}
