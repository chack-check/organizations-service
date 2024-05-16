package ports

import (
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

type MembersPort interface {
	GetByUserId(userId int) *membershipModels.Member
	Save(member membershipModels.Member) (*membershipModels.Member, error)
}

type RolesPort interface {
	Save(role membershipModels.Role, organizationId int) (*membershipModels.Role, error)
	GetByOrganization(organization organizationsModels.Organization) []membershipModels.Role
}
