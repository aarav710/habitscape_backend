package repo

import (
	"context"
	"habitscape/backend/ent"
	"habitscape/backend/ent/habit"
	"habitscape/backend/result"
	habitRequest "habitscape/backend/habit"
)

type HabitRepo interface {
  DeleteHabit(habitId int, userId int)
	GetHabitUsersCount(habitId int) result.Result[int]
	AddUserToHabit(habitId, userId int) result.Result[*ent.Habit]
	UpdateHabit(habitId int, updateHabit habitRequest.HabitRequest) result.Result[*ent.Habit] 
	GetHabit(habitId int) result.Result[*ent.Habit]
}

type HabitRepoImpl struct {
	db *ent.Client
  ctx context.Context
}

func ProvideHabitRepo(db *ent.Client, ctx context.Context) HabitRepo{
	return &HabitRepoImpl{db: db, ctx: ctx}
}

//implement transcations to remove admins, users etc first because there is a seperate table for that
func (repo *HabitRepoImpl) DeleteHabit(habitId int, userId int) {
  
}

func (repo *HabitRepoImpl) GetHabitUsersCount(habitId int) result.Result[int] {
	count, err := repo.db.Habit.Query().Where(habit.ID(habitId)).QueryUsers().Count(repo.ctx)
	result := result.Result[int]{Result: count, Error: err}
	return result
}

func (repo *HabitRepoImpl) AddUserToHabit(habitId, userId int) result.Result[*ent.Habit] {
	habit, err := repo.db.Habit.UpdateOneID(habitId).AddUserIDs(userId).Save(repo.ctx)
	result := result.Result[*ent.Habit]{Result: habit, Error: err}
	return result
} 

func (repo *HabitRepoImpl) UpdateHabit(habitId int, updateHabit habitRequest.HabitRequest) result.Result[*ent.Habit] {
	habit, err := repo.db.Habit.UpdateOneID(habitId).SetDescription(updateHabit.Description).SetFrequency(updateHabit.Frequency).SetTitle(updateHabit.Title).Save(repo.ctx)
	result := result.Result[*ent.Habit]{Result: habit, Error: err}
	return result
}

func (repo *HabitRepoImpl) GetHabit(habitId int) result.Result[*ent.Habit] {
	habit, err := repo.db.Habit.Query().Where(habit.ID(habitId)).Only(repo.ctx)
	result := result.Result[*ent.Habit]{Result: habit, Error: err}
	return result
}
