package notice

import (
	"github.com/movie42/ychung-go-server/presenter"
)

type Service interface {
	FetchNotices() (*[]presenter.Notice, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchNotices() (*[]presenter.Notice, error) {
	return s.repository.ReadNotice()
}
