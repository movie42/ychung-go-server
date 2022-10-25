package service

import (
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/presenter"
)

func (s *service) FetchBlogPosts() (*[]presenter.BlogPost, error) {
	return s.repository.ReadBlogPosts()
}

func (s *service) InsertBlogPost(post *entities.BlogPost) (*entities.BlogPost, error) {
	return s.repository.CreateBlogPost(post)
}

func (s *service) UpdateBlogPost(post *entities.BlogPost) (*entities.BlogPost, error) {
	return s.repository.UpdateBlogPost(post)
}

func (s *service) DeleteBlogPost(ID string) error {
	return s.repository.DeleteBlogPost(ID)
}
