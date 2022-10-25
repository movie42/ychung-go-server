package repository

import (
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	ReadNotices() (*[]presenter.Notice, error)
	DeleteNotice(ID string) error
	UpdateNotice(notice *entities.Notice) (*entities.Notice, error)
	CreateNotice(notice *entities.Notice) (*entities.Notice, error)
	ReadWeeklies() (*[]presenter.Weekly, error)
	DeleteWeekly(ID string) error
	UpdateWeekly(weekly *entities.Weekly) (*entities.Weekly, error)
	CreateWeekly(weekly *entities.Weekly) (*entities.Weekly, error)
	ReadBlogPosts() (*[]presenter.BlogPost, error)
	DeleteBlogPost(ID string) error
	UpdateBlogPost(post *entities.BlogPost) (*entities.BlogPost, error)
	CreateBlogPost(post *entities.BlogPost) (*entities.BlogPost, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}
