package controller

import (
	"habitscape/backend/cache"
	"habitscape/backend/result"
	DTO "habitscape/backend/user"
	"habitscape/backend/user/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
  router *gin.Engine
  service service.UserService
	cache *cache.Cache
}

func ProvideUserController(router *gin.Engine, service service.UserService, cache *cache.Cache) {
  controller := UserController{router: router, service: service, cache: cache}
	controller.router.GET("/users/:userId", controller.GetUser)
	controller.router.GET("/users", controller.GetUsers)
	controller.router.GET("/habits/:habitId/users", controller.GetHabitUsers)
	controller.router.DELETE("/habits/:habitId/users/:userId", controller.RemoveUserFromHabit)
}

func (controller *UserController) GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter input type.",
		})
	}
  user, requestError := controller.service.GetUser(userId)
  if requestError != nil {
		c.JSON(result.GetStatusCode(requestError), gin.H{
			"error": requestError,
		})
	}
	userResponse := DTO.ConvertToResponseDTO(user)
	c.JSON(200, gin.H{
		"user": userResponse,
	})
}

func (controller *UserController) GetUsers(c *gin.Context) {
  username := c.DefaultQuery("username", "")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "25"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameter input type.",
		})
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameter input type.",
		})
	}
		
  users, requestError := controller.service.GetUsers(username, offset, limit)
  if requestError != nil {
		c.JSON(result.GetStatusCode(requestError), gin.H{
			"error": requestError,
		})
	}	
  var usersResponse []DTO.UserResponse

	for _, user := range users {
    userResponse := DTO.ConvertToResponseDTO(user)
    usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(200, gin.H{
		"users": usersResponse,
	})
}

func (controller *UserController) GetHabitUsers(c *gin.Context) {
	habitId, err := strconv.Atoi(c.Param("habitId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter input type.",
		})
	}
	users, requestError := controller.service.GetHabitUsers(habitId)
	if requestError != nil {
		c.JSON(result.GetStatusCode(requestError), gin.H{
			"error": requestError,
		})
	}	

  var usersResponse []DTO.UserResponse
	for _, user := range users {
    userResponse := DTO.ConvertToResponseDTO(user)
    usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(200, gin.H{
		"users": usersResponse,
	})
}

func (controller *UserController) RemoveUserFromHabit(c *gin.Context) {
	habitId, err := strconv.Atoi(c.Param("habitId"))
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameter input type.",
		})
	}
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameter input type.",
		})
	}
	requestError := controller.service.RemoveUserFromHabit(habitId, userId)
  if requestError != nil {
		c.JSON(result.GetStatusCode(requestError), gin.H{
			"error": requestError,
		})
	}
	c.JSON(200, gin.H{})
}

