package repository

import (
	"context"
	"time"

	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) ReadBlogPosts() (*[]presenter.BlogPost, error) {
	var posts []presenter.BlogPost
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var post presenter.BlogPost
		_ = cursor.Decode(&post)
		posts = append(posts, post)
	}
	return &posts, nil
}

func (r *repository) DeleteBlogPost(ID string) error {
	postId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": postId})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateBlogPost(post *entities.BlogPost) (*entities.BlogPost, error) {
	post.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": post}, bson.M{"$set": post})
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) CreateBlogPost(post *entities.BlogPost) (*entities.BlogPost, error) {
	post.ID = primitive.NewObjectID()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
