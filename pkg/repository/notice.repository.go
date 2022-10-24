package repository

import (
	"context"
	"time"

	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) UpdateNotice(notice *entities.Notice) (*entities.Notice, error) {
	notice.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": notice.ID}, bson.M{"$set": notice})
	if err != nil {
		return nil, err
	}
	return notice, nil
}

func (r *repository) DeleteNotice(ID string) error {
	noticeID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": noticeID})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ReadNotices() (*[]presenter.Notice, error) {
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

func (r *repository) CreateNotice(notice *entities.Notice) (*entities.Notice, error) {
	notice.ID = primitive.NewObjectID()
	notice.CreatedAt = time.Now()
	notice.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), notice)
	if err != nil {
		return nil, err
	}
	return notice, nil
}
