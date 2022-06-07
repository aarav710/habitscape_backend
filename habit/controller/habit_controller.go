package controller

import (
	DTO "habitscape/backend/habit"
	"habitscape/backend/habit/service"
	"habitscape/backend/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HabitController struct {
  router *gin.Engine
  service service.HabitService
}

func ProvideHabitController(router *gin.Engine, service service.HabitService) {
  controller := HabitController{router: router, service: service}
	controller.router.DELETE("habits/:habitId", controller.DeleteHabit)
	controller.router.GET("/habits/:habitId", controller.GetHabit)
}

func (controller *HabitController) DeleteHabit(c *gin.Context) {
  
}

func (controller *HabitController) GetHabit(c *gin.Context) {
	habitId, err := strconv.Atoi(c.Param("habitId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameter input type.",
		})
	}
	habit, requestError := controller.service.GetHabit(habitId)
  if requestError != nil {
		c.JSON(result.GetStatusCode(requestError), gin.H{
			"error": requestError.Error,
		})
	}
	habitResponse := DTO.ConvertToResponseDTO(habit)
  c.JSON(200, gin.H{
		"habit": habitResponse,
	})
}

