package controller

import (
	"github.com/gin-gonic/gin"
	"habitscape/backend/invitation/service"
)

type InvitationController struct {
  router *gin.Engine
  service service.InvitationService
}

func ProvideInvitationController(router *gin.Engine, service service.InvitationService) {
  controller := InvitationController{router: router, service: service}
	controller.router.DELETE("invitations/:invitationId", controller.DeleteInvitation)
}

func (controller *InvitationController) DeleteInvitation(c *gin.Context) {
  
}