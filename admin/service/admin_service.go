package service

import (
  "habitscape/backend/admin/repo"
)

type AdminService interface {
  DeleteAdmin(habitId int)
}

type AdminServiceImpl struct {
	repo repo.AdminRepo
}

func ProvideAdminService(repo repo.AdminRepo) AdminService {
  return &AdminServiceImpl{repo: repo}
}

func (service *AdminServiceImpl) DeleteAdmin(adminId int) {
  
}