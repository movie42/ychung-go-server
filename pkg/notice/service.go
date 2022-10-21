package notice

import (
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
)

type Service interface {
	FetchNotices() (*[]presenter.Notice, error)
	InsertNotice(notice *entities.Notice) (*entities.Notice, error)
	DeleteNotice(ID string) error
	UpdateNotice(notice *entities.Notice) (*entities.Notice, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
func (s *service) UpdateNotice(notice *entities.Notice) (*entities.Notice, error) {
	return s.repository.UpdateNotice(notice)
}

func (s *service) FetchNotices() (*[]presenter.Notice, error) {
	return s.repository.ReadNotice()
}

func (s *service) InsertNotice(notice *entities.Notice) (*entities.Notice, error) {
	return s.repository.CreateNotice(notice)
}

func (s *service) DeleteNotice(ID string) error {
	return s.repository.DeleteNotice(ID)
}
