package models

import (
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

type Invite struct {
	id           string
	organization organizationsModels.Organization
	userId       int
	role         membershipModels.Role
	status       string
}

func (invite Invite) GetId() string {
	return invite.id
}

func (invite Invite) GetOrganization() organizationsModels.Organization {
	return invite.organization
}

func (invite Invite) GetUserId() int {
	return invite.userId
}

func (invite Invite) GetRole() membershipModels.Role {
	return invite.role
}

func (invite Invite) GetStatus() string {
	return invite.status
}

func (invite *Invite) SetStatus(status string) {
	invite.status = status
}

func NewInvite(id string, organization organizationsModels.Organization, userId int, role membershipModels.Role, status string) Invite {
	return Invite{
		id:           id,
		organization: organization,
		userId:       userId,
		role:         role,
		status:       status,
	}
}
