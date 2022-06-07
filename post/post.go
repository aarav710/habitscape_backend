package post

import "time"


type PostRequest struct {
  PostBase	
}

type PostBase struct {
	PhotoUrl string `json:"photo_url"`
  Caption string `json:"caption" binding:"required,max=255"`
}

type PostResponse struct {
	ID int `json:"ID"`
	PostBase 
	DateCreated time.Time `json:"date_created"`
} 