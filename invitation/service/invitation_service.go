package service

import (
  "habitscape/backend/invitation/repo"
)

type InvitationService interface {
  DeleteInvitation(invitationId int)
}

type InvitationServiceImpl struct {
	repo repo.InvitationRepo
}

func ProvideInvitationService(repo repo.InvitationRepo) InvitationService {
  return &InvitationServiceImpl{repo: repo}
}

func (service *InvitationServiceImpl) DeleteInvitation(invitationId int) {
  
}