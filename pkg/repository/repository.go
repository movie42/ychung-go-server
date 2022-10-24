package repository

import (
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateNotice(notice *entities.Notice) (*entities.Notice, error)
	ReadNotices() (*[]presenter.Notice, error)
	UpdateNotice(notice *entities.Notice) (*entities.Notice, error)
	DeleteNotice(ID string) error
	ReadWeeklies() (*[]presenter.Weekly, error)
	DeleteWeekly(ID string) error
	CreateWeekly(weekly *entities.Weekly) (*entities.Weekly, error)
	UpdateWeekly(weekly *entities.Weekly) (*entities.Weekly, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}
