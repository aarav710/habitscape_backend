package invitation

import (
	"habitscape/backend/ent"
	habitRequest "habitscape/backend/habit"
	userRequest "habitscape/backend/user"
)

type InvitationRequest struct {
	UserId int `json:"userId"`
	InvitationBase
}

type InvitationBase struct {

}

type InvitationResponse struct {
	ID int `json:"ID"`
	User userRequest.UserResponse `json:"user"`
	Habit habitRequest.HabitResponse `json:"habit"`
	InvitationBase
}

func convertToEnt(request *InvitationRequest) *ent.Invitation {
  invite := &ent.Invitation{}
  invite.Edges.User = &ent.User{ID: request.UserId}
	return invite
}

func convertToResponseDTO(model *ent.Invitation) InvitationResponse {
	response := InvitationResponse{ID: model.ID}
	response.User.ID = model.Edges.User.ID
	response.User.DateCreated = model.Edges.User.CreatedAt
	response.User.Email = model.Edges.User.Email
	response.User.Bio = model.Edges.User.Bio
	response.User.PhotoUrl = model.Edges.User.PhotoURL
	response.User.Username = model.Edges.User.Username
	response.Habit.ID = model.Edges.Habit.ID
  response.Habit.Title = model.Edges.Habit.Title
	response.Habit.Description = model.Edges.Habit.Description
	response.Habit.PhotoUrl = model.Edges.Habit.PhotoURL
	response.Habit.Frequency = model.Edges.Habit.Frequency
	return response
}
