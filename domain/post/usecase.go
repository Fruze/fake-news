package post

import (
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/request"
)

type Usecase interface {
	FetchPost() ([]*models.Post, error)
	CreatePost(pr request.PostRequest) (*models.Post, error)
}
