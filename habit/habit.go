package habit

import (
	"habitscape/backend/ent"
	"time"
)

type HabitRequest struct {
  HabitBase
}

type HabitBase struct {
  Description string `json:"description" binding:"required,max=255,min=10"`
	Frequency int `json:"frequency" binding:"required,gt=0,lte=7"`
	Title string `json:"title" binding:"required,max=100,min=10"`
  PhotoUrl string `json:"photo_url"`
}

type HabitResponse struct {
	ID int `json:"ID"`
  HabitBase
  DateCreated time.Time `json:"date_created" binding:"required"`
}

func ConvertToResponseDTO(habit *ent.Habit) HabitResponse {
  habitResponse := HabitResponse{}
  habitResponse.Description = habit.Description
  habitResponse.Frequency = habit.Frequency
  habitResponse.Title = habit.Title
  habitResponse.PhotoUrl = habit.PhotoURL
  habitResponse.ID = habit.ID
  habitResponse.DateCreated = habit.DateCreated
  return habitResponse
}

func ConvertToEnt(request *HabitRequest) *ent.Habit {
  habit := &ent.Habit{}
  habit.Description = request.Description
  habit.Frequency = request.Frequency
  habit.Title = request.Title
  habit.PhotoURL = request.PhotoUrl
  return habit
}