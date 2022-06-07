package repo

import (
	"context"
	"habitscape/backend/ent"
	"habitscape/backend/ent/admin"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/user"
	"habitscape/backend/result"
)

type AdminRepo interface {
  DeleteAdmin(adminId int) error
	GetAdminsByHabitId(habitId int) result.Result[[]*ent.Admin] 
	IsUserHabitAdmin(userId, habitId int) result.Result[bool]
}

type AdminRepoImpl struct {
	db *ent.Client
	ctx context.Context
}

func ProvideAdminRepo(db *ent.Client, ctx context.Context) AdminRepo{
	return &AdminRepoImpl{db: db, ctx: ctx}
}

func (repo *AdminRepoImpl) DeleteAdmin(adminId int) error {
	err := repo.db.Admin.DeleteOneID(adminId).Exec(repo.ctx)
	return err
}

func (repo *AdminRepoImpl) GetAdminsByHabitId(habitId int) result.Result[[]*ent.Admin] {
  admins, err := repo.db.Habit.Query().Where(habit.ID(habitId)).QueryAdmins().WithUser().All(repo.ctx)
	result := result.Result[[]*ent.Admin] {Result: admins, Error: err}
	return result
}

func (repo *AdminRepoImpl) IsUserHabitAdmin(userId, habitId int) result.Result[bool] {
	isUserAdmin, err := repo.db.Admin.Query().Where(
	                                          	admin.And(
	                                          		admin.HasUserWith(user.ID(userId)),
	                                          		admin.HasHabitWith(habit.ID(habitId)),
	                                          	),
	                                          ).Exist(repo.ctx)
	result := result.Result[bool]{Result: isUserAdmin, Error: err}
	return result
}

