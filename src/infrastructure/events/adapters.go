package events

import "github.com/chack-check/organizations-service/domain/organizations/models"

type OrganizationEventsAdapter struct{}

func (adapter OrganizationEventsAdapter) SendOrganizationCreated(organization models.Organization) {}

func (adapter OrganizationEventsAdapter) SendOrganizationChanged(organization models.Organization) {}

func (adapter OrganizationEventsAdapter) SendOrganizationDeactivated(organization models.Organization) {
}

func (adapter OrganizationEventsAdapter) SendOrganizationActivated(organization models.Organization) {
}
