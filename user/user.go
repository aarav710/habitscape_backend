package user

import (
	"habitscape/backend/ent"
	"time"
)


type UserRequest struct {
  UserBase
}

type UserBase struct {
  Bio *string `json:"bio"`
	Email string `json:"email" binding:"required,email"`
	Username string `json:"username"`
	PhotoUrl string `json:"photo_url"`
}

type UserResponse struct {
	ID int `json:"ID"`
	DateCreated time.Time `json:"date_created"`
  UserBase
}

func ConvertToEnt(request *UserRequest) *ent.User {
  user := &ent.User{}
	user.Bio = request.Bio
	user.Email = request.Email
	user.Username = request.Username
	user.PhotoURL = request.PhotoUrl
	return user
}

func ConvertToResponseDTO(user *ent.User) UserResponse {
  userResponse := UserResponse{}
	userResponse.ID = user.ID
	userResponse.DateCreated = user.CreatedAt
	userResponse.Bio = user.Bio
	userResponse.Email = user.Email
	userResponse.Username = user.Username
	userResponse.PhotoUrl = user.PhotoURL
	return userResponse
}