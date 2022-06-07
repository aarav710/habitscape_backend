package service

import (
	"habitscape/backend/ent"
	"habitscape/backend/habit/repo"
)

type HabitService interface {
  DeleteHabit(habitId int)
  GetHabit(habitId int) (*ent.Habit, error) 
}

type HabitServiceImpl struct {
	repo repo.HabitRepo
}

func ProvideHabitService(repo repo.HabitRepo) HabitService {
  return &HabitServiceImpl{repo: repo}
}

func (service *HabitServiceImpl) DeleteHabit(habitId int) {
  
}

func (service *HabitServiceImpl) GetHabit(habitId int) (*ent.Habit, error) {
  habit := service.repo.GetHabit(habitId)    
  return habit.Result, habit.Error
}