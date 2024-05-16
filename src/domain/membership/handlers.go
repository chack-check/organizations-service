package membership

import (
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	membershipPorts "github.com/chack-check/organizations-service/domain/membership/ports"
	organizationsPorts "github.com/chack-check/organizations-service/domain/organizations/ports"
)

type GetOrganizationRolesHandler struct {
	organizationsPort organizationsPorts.OrganizationsPort
	rolesPort         membershipPorts.RolesPort
}

func (handler GetOrganizationRolesHandler) Execute(userId int, organizationId int) []membershipModels.Role {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return []membershipModels.Role{}
	}

	roles := handler.rolesPort.GetByOrganization(*organization)
	return roles
}

func NewGetOrganizationRolesHandler(
	organizationsPort organizationsPorts.OrganizationsPort,
	rolesPort membershipPorts.RolesPort,
) GetOrganizationRolesHandler {
	return GetOrganizationRolesHandler{
		organizationsPort: organizationsPort,
		rolesPort:         rolesPort,
	}
}
