package repo

import (
	"context"
	"habitscape/backend/ent"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/invitation"
	"habitscape/backend/ent/user"
	"habitscape/backend/result"
)

type InvitationRepo interface {
  DeleteInvitation(invitationId int) error
	GetInvitation(invitationId int) result.Result[*ent.Invitation]
	GetInvitationsByHabitId(habitId int) result.Result[[]*ent.Invitation]
	CreateInvitation(adminId, habitId, userId int) result.Result[*ent.Invitation] 
  GetInvitationByHabitIdAndUserId(userId, habitId int) result.Result[*ent.Invitation]
	GetInvitationIdByHabitIdAndUserId(userId, habitId int) result.Result[int]
	IsInvitationByHabitIdAndUserId(userId, habitId int) result.Result[bool]
}

type InvitationRepoImpl struct {
	db *ent.Client
	ctx context.Context
}

func ProvideInvitationRepo(db *ent.Client, ctx context.Context) InvitationRepo{
	return &InvitationRepoImpl{db: db, ctx: ctx}
}

func (repo *InvitationRepoImpl) DeleteInvitation(invitationId int) error {
  err := repo.db.Invitation.DeleteOneID(invitationId).Exec(repo.ctx)
	return err
}

func (repo *InvitationRepoImpl) GetInvitation(invitationId int) result.Result[*ent.Invitation] {
	invitation, err := repo.db.Invitation.Query().Where(invitation.ID(invitationId)).WithUser().Only(repo.ctx)
	result := result.Result[*ent.Invitation]{Result: invitation, Error: err}
	return result
}

func (repo *InvitationRepoImpl) GetInvitationsByHabitId(habitId int) result.Result[[]*ent.Invitation] {
	invitations, err := repo.db.Habit.Query().Where(habit.ID(habitId)).QueryInvitations().WithUser().All(repo.ctx)
	result := result.Result[[]*ent.Invitation]{Result: invitations, Error: err}
	return result
}

func (repo *InvitationRepoImpl) CreateInvitation(adminId, habitId, userId int) result.Result[*ent.Invitation] {
  invitation, err := repo.db.Invitation.Create().SetAdminID(adminId).SetHabitID(habitId).SetUserID(userId).Save(repo.ctx)
	result := result.Result[*ent.Invitation]{Result: invitation, Error: err}
	return result
}

func (repo *InvitationRepoImpl) GetInvitationByHabitIdAndUserId(userId, habitId int) result.Result[*ent.Invitation] {
	invitation, err := repo.db.Invitation.Query().Where(
	                                              	invitation.And(
	                                              		invitation.HasUserWith(user.ID(userId)),
	                                              		invitation.HasHabitWith(habit.ID(habitId)),
	                                              	),
	                                              ).WithUser().Only(repo.ctx)
	result := result.Result[*ent.Invitation]{Result: invitation, Error: err}
	return result
}

func (repo *InvitationRepoImpl) GetInvitationIdByHabitIdAndUserId(userId, habitId int) result.Result[int] {
	invitationId, err := repo.db.Invitation.Query().Where(
	                                                	invitation.And(
	                                                		invitation.HasUserWith(user.ID(userId)),
	                                                		invitation.HasHabitWith(habit.ID(habitId)),
	                                                	),
	                                                ).OnlyID(repo.ctx)
	result := result.Result[int]{Result: invitationId, Error: err}
	return result
}

func (repo *InvitationRepoImpl) IsInvitationByHabitIdAndUserId(userId, habitId int) result.Result[bool] {
	isHabitInvitation, err := repo.db.Invitation.Query().Where(
	                                                      	invitation.And(
	                                                      		invitation.HasUserWith(user.ID(userId)),
	                                                      		invitation.HasHabitWith(habit.ID(habitId)),
	                                                      	),
	                                                     ).Exist(repo.ctx)
	result := result.Result[bool]{Result: isHabitInvitation, Error: err}
	return result
}