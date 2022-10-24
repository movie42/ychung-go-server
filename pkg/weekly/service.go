package weekly

import (
	"github.com/movie42/ychung-go-server/presenter"
)

type service struct {
	repository Repository
}

func (s *service) FetchWeeklies() (*[]presenter.Weekly, error) {
	return s.repository.ReadWeekies()
}
