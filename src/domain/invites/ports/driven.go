package ports

import (
	invitesModels "github.com/chack-check/organizations-service/domain/invites/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

type InvitesPort interface {
	Save(invite invitesModels.Invite) (*invitesModels.Invite, error)
	GetById(inviteId string) *invitesModels.Invite
	GetByIdForUser(inviteId string, userId int) *invitesModels.Invite
	GetAllForOrganization(organization organizationsModels.Organization) []invitesModels.Invite
	GetAllForUser(userId int) []invitesModels.Invite
}

type InviteEventsPort interface {
	SendInviteCreated(invite invitesModels.Invite)
}
