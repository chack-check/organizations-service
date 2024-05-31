package events

import (
	invitesModels "github.com/chack-check/organizations-service/domain/invites/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

type OrganizationEventsAdapter struct{}

func (adapter OrganizationEventsAdapter) SendOrganizationCreated(organization organizationsModels.Organization) {
}

func (adapter OrganizationEventsAdapter) SendOrganizationChanged(organization organizationsModels.Organization) {
}

func (adapter OrganizationEventsAdapter) SendOrganizationDeactivated(organization organizationsModels.Organization) {
}

func (adapter OrganizationEventsAdapter) SendOrganizationActivated(organization organizationsModels.Organization) {
}

type InviteEventsAdapter struct{}

func (adapter InviteEventsAdapter) SendInviteCreated(invite invitesModels.Invite) {}
