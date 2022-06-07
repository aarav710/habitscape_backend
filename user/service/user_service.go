package service

import (
	"errors"
	"habitscape/backend/ent"
	"habitscape/backend/user/repo"
)

type UserService interface {
  DeleteUser(userId int)
  GetUser(userId int) (*ent.User, error)
  GetUsers(username string, offset, limit int) ([]*ent.User, error)
  GetHabitUsers(habitId int) ([]*ent.User, error)
  RemoveUserFromHabit(habitId, userId int) error
}

type UserServiceImpl struct {
	repo repo.UserRepo
}

func ProvideUserService(repo repo.UserRepo) UserService {
  return &UserServiceImpl{repo: repo}
}

func (service *UserServiceImpl) DeleteUser(userId int) {
  
}

func (service *UserServiceImpl) GetUser(userId int) (*ent.User, error) {
  user := service.repo.GetUser(userId)
  return user.Result, user.Error
}

func (service *UserServiceImpl) GetUsers(username string, offset, limit int) ([]*ent.User, error) {
  users := service.repo.GetUsersByUsername(username, offset, limit)
  return users.Result, users.Error
}

func (service *UserServiceImpl) GetHabitUsers(habitId int) ([]*ent.User, error) {
  users := service.repo.GetHabitUsers(habitId)
  return users.Result, users.Error
}

func (service *UserServiceImpl) RemoveUserFromHabit(habitId, userId int) error {
  panic("not implemented yet")
  return errors.New("")
}