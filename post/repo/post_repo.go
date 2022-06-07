package repo

import (
	"context"
	"habitscape/backend/ent"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/post"
	postRequest "habitscape/backend/post"
	"habitscape/backend/result"
)

type PostRepo interface {
  DeletePost(postId int) error
	GetPost(postId int) result.Result[*ent.Post] 
	UpdatePost(updatePost postRequest.PostRequest, postId int) result.Result[*ent.Post]
	GetPostsByHabitId(habitId int) result.Result[[]*ent.Post]
	CreatePost(createPost postRequest.PostRequest, userId, habitId int) result.Result[*ent.Post]
}

type PostRepoImpl struct {
	db *ent.Client
	ctx context.Context
}

func ProvidePostRepo(db *ent.Client, ctx context.Context) PostRepo{
	return &PostRepoImpl{db: db, ctx: ctx}
}

func (repo *PostRepoImpl) DeletePost(postId int) error {
	err := repo.db.Post.DeleteOneID(postId).Exec(repo.ctx)
	return err
}

func (repo *PostRepoImpl) GetPost(postId int) result.Result[*ent.Post] {
  post, err := repo.db.Post.Query().Where(post.ID(postId)).Only(repo.ctx)
	result := result.Result[*ent.Post]{Result: post, Error: err}
	return result
}

func (repo *PostRepoImpl) UpdatePost(updatePost postRequest.PostRequest, postId int) result.Result[*ent.Post] {
  post, err := repo.db.Post.UpdateOneID(postId).SetCaption(updatePost.Caption).Save(repo.ctx)
	result := result.Result[*ent.Post]{Result: post, Error: err}
	return result
}

func (repo *PostRepoImpl) GetPostsByHabitId(habitId int) result.Result[[]*ent.Post] {
	posts, err := repo.db.Habit.Query().Where(habit.ID(habitId)).QueryPosts().All(repo.ctx)
	result := result.Result[[]*ent.Post]{Result: posts, Error: err}
	return result
}

func (repo *PostRepoImpl) CreatePost(createPost postRequest.PostRequest, userId, habitId int) result.Result[*ent.Post] {
	post, err := repo.db.Post.Create().SetCaption(createPost.Caption).SetPhotoURL(createPost.PhotoUrl).
	               SetHabitID(habitId).SetUserID(userId).Save(repo.ctx)
	result := result.Result[*ent.Post]{Result: post, Error: err}
  return result
}