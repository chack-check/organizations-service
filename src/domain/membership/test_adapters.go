package membership

import (
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

type MockRolesAdapter struct{}

func (adapter MockRolesAdapter) Save(role membershipModels.Role, organizationId int) (*membershipModels.Role, error) {
	return &role, nil
}

func (adapter MockRolesAdapter) GetByOrganization(organization organizationsModels.Organization) []membershipModels.Role {
	return []membershipModels.Role{
		membershipModels.NewRole(1, "member", "Member", []membershipModels.Permission{}),
	}
}
