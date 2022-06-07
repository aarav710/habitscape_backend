package controller

import (
	"habitscape/backend/comment/service"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
  router *gin.Engine
  service service.CommentService
}

func ProvideCommentController(router *gin.Engine, service service.CommentService) {
  controller := CommentController{router: router, service: service}
	controller.router.DELETE("/comments/:commentId", controller.DeleteComment)
	controller.router.GET("/comments", controller.Hello)
}

func (controller *CommentController) DeleteComment(c *gin.Context) {
  
}

func (controller *CommentController) Hello(c *gin.Context) {
	message := controller.service.Hello("stan")
  c.JSON(200, gin.H{
		"message": message,
	})
}