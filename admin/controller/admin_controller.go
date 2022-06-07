package controller

import (
	"habitscape/backend/admin/service"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
  router *gin.Engine
  service service.AdminService
}

func ProvideAdminController(router *gin.Engine, service service.AdminService) {
  controller := AdminController{router: router, service: service}
	controller.router.DELETE("/admins/:adminId", controller.DeleteAdmin)
}

func (controller *AdminController) DeleteAdmin(c *gin.Context) {
  
}