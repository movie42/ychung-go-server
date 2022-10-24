package repository

import (
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	DeleteNotice(ID string) error
	CreateNotice(notice *entities.Notice) (*entities.Notice, error)
	ReadNotices() (*[]presenter.Notice, error)
	UpdateNotice(notice *entities.Notice) (*entities.Notice, error)
	ReadWeeklies() (*[]presenter.Weekly, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}
