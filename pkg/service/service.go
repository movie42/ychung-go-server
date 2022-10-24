package service

import (
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/pkg/repository"
	"github.com/movie42/ychung-go-server/presenter"
)

type Service interface {
	FetchWeeklies() (*[]presenter.Weekly, error)
	InsertWeekly(weekly *entities.Weekly) (*entities.Weekly, error)
	UpdateWeekly(weekly *entities.Weekly) (*entities.Weekly, error)
	DeleteWeekly(ID string) error
	FetchNotices() (*[]presenter.Notice, error)
	InsertNotice(notice *entities.Notice) (*entities.Notice, error)
	UpdateNotice(notice *entities.Notice) (*entities.Notice, error)
	DeleteNotice(ID string) error
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}
