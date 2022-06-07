package controller

import (
	"github.com/gin-gonic/gin"
	"habitscape/backend/post/service"
)

type PostController struct {
  router *gin.Engine
  service service.PostService
}

func ProvidePostController(router *gin.Engine, service service.PostService) {
  controller := PostController{router: router, service: service}
	controller.router.DELETE("posts/:postId", controller.DeletePost)
}

func (controller *PostController) DeletePost(c *gin.Context) {
  
}