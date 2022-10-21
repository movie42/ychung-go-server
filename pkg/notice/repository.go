package notice

import (
	"context"

	"github.com/movie42/ychung-go-server/presenter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	ReadNotice() (*[]presenter.Notice, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) ReadNotice() (*[]presenter.Notice, error) {
	var notices []presenter.Notice
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var notice presenter.Notice
		_ = cursor.Decode(&notice)
		notices = append(notices, notice)
	}
	return &notices, nil
}
