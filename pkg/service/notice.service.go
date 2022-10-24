package service

import (
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
)

func (s *service) UpdateNotice(notice *entities.Notice) (*entities.Notice, error) {
	return s.repository.UpdateNotice(notice)
}

func (s *service) FetchNotices() (*[]presenter.Notice, error) {
	return s.repository.ReadNotices()
}

func (s *service) InsertNotice(notice *entities.Notice) (*entities.Notice, error) {
	return s.repository.CreateNotice(notice)
}

func (s *service) DeleteNotice(ID string) error {
	return s.repository.DeleteNotice(ID)
}
