package repository

import (
	"context"
	"time"

	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *repository) DeleteWeekly(ID string) error {
	weeklyId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": weeklyId})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateWeekly(weekly *entities.Weekly) (*entities.Weekly, error) {
	weekly.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": weekly.ID}, bson.M{"$set": weekly})
	if err != nil {
		return nil, err
	}
	return weekly, nil
}

func (r *repository) CreateWeekly(weekly *entities.Weekly) (*entities.Weekly, error) {
	weekly.ID = primitive.NewObjectID()
	weekly.CreatedAt = time.Now()
	weekly.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), weekly)
	if err != nil {
		return nil, err
	}
	return weekly, nil
}
