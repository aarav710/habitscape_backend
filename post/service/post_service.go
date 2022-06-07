package service

import (
	"habitscape/backend/post/repo"
)

type PostService interface {
  DeletePost(postId int)
}

type PostServiceImpl struct {
	repo repo.PostRepo
}

func ProvidePostService(repo repo.PostRepo) PostService {
  return &PostServiceImpl{repo: repo}
}

func (service *PostServiceImpl) DeletePost(postId int) {
  
}