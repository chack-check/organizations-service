package ports

import (
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

type MembersPort interface {
	GetByOrganizationAndUserId(organization organizationsModels.Organization, userId int) *membershipModels.Member
	Save(member membershipModels.Member, organizatonId int) (*membershipModels.Member, error)
}

type RolesPort interface {
	Save(role membershipModels.Role, organizationId int) (*membershipModels.Role, error)
	GetByOrganization(organization organizationsModels.Organization) []membershipModels.Role
	GetByOrganizationAndId(organization organizationsModels.Organization, roleId int) *membershipModels.Role
	Delete(role membershipModels.Role) error
}
