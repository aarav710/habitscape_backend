package comment

import (
	"time"
	"habitscape/backend/user"
)


type CommentRequest struct {
  CommentBase
}

type CommentBase struct {
  Text string `json:"text"`
}

type CommentResponse struct {
	ID int `json:"ID"`
  CommentBase
	DateCreated time.Time `json:"date_created"`
	User user.UserResponse  `json:"user"`
}