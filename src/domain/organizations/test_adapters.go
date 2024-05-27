package organizations

import (
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

type MockOrganizationsAdapter struct{}

func (adapter MockOrganizationsAdapter) Save(organization organizationsModels.Organization) (*organizationsModels.Organization, error) {
	return &organization, nil
}

func (adapter MockOrganizationsAdapter) GetForUser(userId int, includeNotActive bool) []organizationsModels.Organization {
	role := membershipModels.NewRole(1, "member", "Member", []membershipModels.Permission{})
	member := membershipModels.NewMember(userId, &role, []membershipModels.Permission{})
	members := []membershipModels.Member{member}
	return []organizationsModels.Organization{
		organizationsModels.NewOrganization(
			1,
			"title",
			"description",
			5,
			5,
			nil,
			members,
			2,
			nil,
		),
	}
}

func (adapter MockOrganizationsAdapter) GetByIdForUser(id int, userId int, includeNotActive bool) *organizationsModels.Organization {
	role := membershipModels.NewRole(1, "member", "Member", []membershipModels.Permission{})
	member := membershipModels.NewMember(userId, &role, []membershipModels.Permission{})
	members := []membershipModels.Member{member}
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		5,
		nil,
		members,
		2,
		nil,
	)
	return &organization
}

func (adapter MockOrganizationsAdapter) GetOpenCountForUser(userId int) int {
	return 1
}

type MockOrganizationEventsAdapter struct{}

func (adapter MockOrganizationEventsAdapter) SendOrganizationCreated(organization organizationsModels.Organization) {
}

func (adapter MockOrganizationEventsAdapter) SendOrganizationChanged(organization organizationsModels.Organization) {
}

func (adapter MockOrganizationEventsAdapter) SendOrganizationDeactivated(organization organizationsModels.Organization) {
}

func (adapter MockOrganizationEventsAdapter) SendOrganizationActivated(organization organizationsModels.Organization) {
}

type MockSubscriptionsAdapter struct{}

func (adapter MockSubscriptionsAdapter) GetUserOrganizationConditions(userId int) (*organizationsModels.OrganizationConditions, error) {
	conditions := organizationsModels.NewOrganizationConditions(5, 5, 5)
	return &conditions, nil
}
