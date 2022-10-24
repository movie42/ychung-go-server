package repository

import (
	"context"

	"github.com/movie42/ychung-go-server/presenter"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) ReadWeeklies() (*[]presenter.Weekly, error) {
	var weeklies []presenter.Weekly
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var weekly presenter.Weekly
		_ = cursor.Decode(&weekly)
		weeklies = append(weeklies, weekly)
	}
	return &weeklies, nil
}
