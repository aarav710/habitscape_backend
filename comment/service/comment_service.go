package service

import (
  "habitscape/backend/comment/repo"
)

type CommentService interface {
  DeleteComment(commentId int) 
  Hello(hello string) string
}

type CommentServiceImpl struct {
	repo repo.CommentRepo
}

func ProvideCommentService(repo repo.CommentRepo) CommentService {
  return &CommentServiceImpl{repo: repo}
}

func (service *CommentServiceImpl) DeleteComment(commentId int) {
  
}

func (service *CommentServiceImpl) Hello(hello string) string {
  return hello
}

