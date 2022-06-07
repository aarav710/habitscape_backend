package admin

import (
	"habitscape/backend/user"
)

type AdminRequest struct {
  AdminBase
	UserId int
}

type AdminBase struct {
	
}

type AdminResponse struct {
	ID int
  User *user.UserResponse
	AdminBase
}

// habits/:habitId/admins