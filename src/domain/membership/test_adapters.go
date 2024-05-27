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

func (adapter MockRolesAdapter) GetByOrganizationAndId(organization organizationsModels.Organization, roleId int) *membershipModels.Role {
	role := membershipModels.NewRole(roleId, "member", "Member", []membershipModels.Permission{})
	return &role
}

func (adapter MockRolesAdapter) Delete(role membershipModels.Role) error {
	return nil
}

type MockMembersAdapter struct{}

func (adapter MockMembersAdapter) GetByOrganizationAndUserId(organization organizationsModels.Organization, userId int) *membershipModels.Member {
	role := membershipModels.NewRole(
		1,
		"role_code",
		"role_name",
		[]membershipModels.Permission{},
	)
	member := membershipModels.NewMember(
		1,
		&role,
		[]membershipModels.Permission{},
	)
	return &member
}
func (adapter MockMembersAdapter) Save(member membershipModels.Member, organizationId int) (*membershipModels.Member, error) {
	return &member, nil
}
