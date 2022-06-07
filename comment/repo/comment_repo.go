package repo

import (
	"context"
	"habitscape/backend/ent"
	"habitscape/backend/ent/comment"
	"habitscape/backend/ent/post"
	"habitscape/backend/ent/user"
	"habitscape/backend/result"
)

type CommentRepo interface {
	DeleteComment(commentId int) error
	GetCommentsByPostId(postId int) result.Result[[]*ent.Comment]
	IsCommentByUser(commentId, userId int) result.Result[bool] 
}

type CommentRepoImpl struct {
	db *ent.Client
	ctx context.Context
}

func ProvideCommentRepo(db *ent.Client, ctx context.Context) CommentRepo{
	return &CommentRepoImpl{db: db, ctx: ctx}
}

func (repo *CommentRepoImpl) DeleteComment(commentId int) error {
  err := repo.db.Comment.DeleteOneID(commentId).Exec(repo.ctx)
	return err
}

func (repo *CommentRepoImpl) GetCommentsByPostId(postId int) result.Result[[]*ent.Comment] {
	comments, err := repo.db.Post.Query().Where(post.ID(postId)).QueryComments().WithUser().All(repo.ctx)
	result := result.Result[[]*ent.Comment]{Result: comments, Error: err}
	return result
}

func (repo *CommentRepoImpl) IsCommentByUser(commentId, userId int) result.Result[bool] {
	isCommentByUser, err := repo.db.Comment.Query().Where(comment.And(comment.ID(commentId), comment.HasUserWith(user.ID(userId)))).Exist(repo.ctx)
	result := result.Result[bool]{Result: isCommentByUser, Error: err}
	return result
}