package repo

import (
	"context"
	"habitscape/backend/ent"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/user"
	"habitscape/backend/result"
)

type UserRepo interface {
  GetUser(userId int) result.Result[*ent.User]
	UpdateUser(updateUser *ent.User, userId int) result.Result[*ent.User]
	GetUsersByUsername(username string, offset, limit int) result.Result[[]*ent.User]
	CreateUser(createUser *ent.User) result.Result[*ent.User]
	RemoveUserFromHabit(habitId, userId int) result.Result[*ent.User]
  GetHabitUsers(habitId int) result.Result[[]*ent.User] 
	IsUserInHabit(habitId, userId int) result.Result[bool]
	GetUserByEmail(email string) result.Result[*ent.User]
	GetUsersInvitedByHabitId(habitId int) result.Result[[]*ent.User]
}

type UserRepoImpl struct {
	db *ent.Client
	ctx context.Context
}

func ProvideUserRepo(db *ent.Client, ctx context.Context) UserRepo{
	return &UserRepoImpl{db: db, ctx: ctx}
}

func (repo *UserRepoImpl) GetUser(userId int) result.Result[*ent.User]{
  user, err := repo.db.User.Query().Where(user.ID(userId)).Only(repo.ctx)
	result := result.Result[*ent.User]{Result: user, Error: err}
	return result
}

func (repo *UserRepoImpl) UpdateUser(updateUser *ent.User, userId int) result.Result[*ent.User] {
  user, err := repo.db.User.UpdateOneID(userId).SetBio(*updateUser.Bio).Save(repo.ctx)
	result := result.Result[*ent.User]{Result: user, Error: err}
	return result
}

func (repo *UserRepoImpl) GetUsersByUsername(username string, offset, limit int) result.Result[[]*ent.User] {
	users, err := repo.db.User.Query().Where(user.UsernameContains(username)).Offset(offset).Limit(limit).All(repo.ctx)
	result := result.Result[[]*ent.User]{Result: users, Error: err}
	return result
}

func (repo *UserRepoImpl) CreateUser(createUser *ent.User) result.Result[*ent.User] {
	user, err := repo.db.User.Create().SetBio(*createUser.Bio).SetEmail(createUser.Email).SetPhotoURL(createUser.PhotoURL).SetUsername(createUser.Username).Save(repo.ctx)
	result := result.Result[*ent.User]{Result: user, Error: err}
	return result
}

func (repo *UserRepoImpl) RemoveUserFromHabit(habitId, userId int) result.Result[*ent.User] {
	user, err := repo.db.User.UpdateOneID(userId).RemoveHabitIDs(habitId).Save(repo.ctx)
	result := result.Result[*ent.User]{Result: user, Error: err}
	return result
}

func (repo *UserRepoImpl) GetHabitUsers(habitId int) result.Result[[]*ent.User] {
	users, err := repo.db.Habit.Query().Where(habit.ID(habitId)).QueryUsers().All(repo.ctx)
	result := result.Result[[]*ent.User]{Result: users, Error: err}
	return result
}

func (repo *UserRepoImpl) IsUserInHabit(habitId, userId int) result.Result[bool] {
	isUserInHabit, err := repo.db.Habit.Query().Where(habit.ID(habitId)).QueryUsers().Where(user.ID(userId)).Exist(repo.ctx)
	result := result.Result[bool]{Result: isUserInHabit, Error: err}
	return result
}

func (repo *UserRepoImpl) GetUserByEmail(email string) result.Result[*ent.User] {
	user, err := repo.db.User.Query().Where(user.Email(email)).Only(repo.ctx)
	result := result.Result[*ent.User]{Result: user, Error: err}
	return result
}

func (repo *UserRepoImpl) GetUsersInvitedByHabitId(habitId int) result.Result[[]*ent.User] {
	users, err := repo.db.Habit.Query().Where(habit.ID(habitId)).QueryUsers().All(repo.ctx)
	result := result.Result[[]*ent.User]{Result: users, Error: err}
	return result
}