package ports

import "github.com/chack-check/organizations-service/domain/organizations/models"

type OrganizationsPort interface {
	Save(organization models.Organization) (*models.Organization, error)
	GetForUser(userId int, includeNotActive bool) []models.Organization
	GetByIdForUser(id int, userId int, includeNotActive bool) *models.Organization
	GetOpenCountForUser(userId int) int
}

type OrganizationEventsPort interface {
	SendOrganizationCreated(orgaization models.Organization)
	SendOrganizationChanged(orgaization models.Organization)
	SendOrganizationDeactivated(orgaization models.Organization)
	SendOrganizationActivated(orgaization models.Organization)
}

type SubscriptionsPort interface {
	GetUserOrganizationConditions(userId int) (*models.OrganizationConditions, error)
}
