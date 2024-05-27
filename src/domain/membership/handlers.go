package membership

import (
	"fmt"
	"slices"

	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	membershipPorts "github.com/chack-check/organizations-service/domain/membership/ports"
	organizationsPorts "github.com/chack-check/organizations-service/domain/organizations/ports"
)

var (
	ErrOrganizationNotFound    = fmt.Errorf("organization not found")
	ErrDeleteMembersPermission = fmt.Errorf("you have no permissions to delete members from organization")
	ErrRoleNotFound            = fmt.Errorf("role not found")
	ErrSavingMember            = fmt.Errorf("error saving member")
	ErrSetMemberPermissions    = fmt.Errorf("you have no permissions to set permissions for members in organization")
	ErrCreateRolesPermission   = fmt.Errorf("you have no permissions to create roles in organization")
	ErrSavingRole              = fmt.Errorf("error saving role")
	ErrEditRolesPermission     = fmt.Errorf("you have no perissions to edit roles in organization")
	ErrDeleteRolesPermission   = fmt.Errorf("you have no permissions to delete roles in organization")
	ErrDeletingRole            = fmt.Errorf("error deleting role")
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

type GetOrganizationMembersHandler struct {
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler GetOrganizationMembersHandler) Execute(userId int, organizationId int) (*[]membershipModels.Member, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	response := organization.GetMembers()
	return &response, nil
}

type SetMemberRoleHandler struct {
	membersPort       membershipPorts.MembersPort
	rolesPort         membershipPorts.RolesPort
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler SetMemberRoleHandler) Execute(userId int, organizationId int, memberId int, roleId *int) (*membershipModels.Member, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(*organization, userId)
	if member == nil {
		return nil, ErrOrganizationNotFound
	}

	if !member.HasPermission(RemoveMembersPermission.GetCode()) && organization.GetOwnerId() != userId {
		return nil, ErrDeleteMembersPermission
	}

	if roleId == nil {
		member.SetRole(nil)
	} else {
		role := handler.rolesPort.GetByOrganizationAndId(*organization, *roleId)
		if role == nil {
			return nil, ErrRoleNotFound
		}

		member.SetRole(role)
	}

	savedMember, err := handler.membersPort.Save(*member, organizationId)
	if err != nil {
		return nil, ErrSavingMember
	}

	return savedMember, nil
}

type SetMemberPermissionsHandler struct {
	membersPort       membershipPorts.MembersPort
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler SetMemberPermissionsHandler) Execute(userId int, organizationId int, memberId int, permissions []string) (*membershipModels.Member, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(*organization, userId)
	if member == nil {
		return nil, ErrOrganizationNotFound
	}

	if !member.HasPermission(SetMembersPermissions.GetCode()) && organization.GetOwnerId() != userId {
		return nil, ErrSetMemberPermissions
	}

	var permissionsModels []membershipModels.Permission
	for _, permission := range AllPermissions {
		if slices.Contains(permissions, permission.GetCode()) {
			permissionsModels = append(permissionsModels, permission)
		}
	}

	member.SetPermissions(permissionsModels)
	savedMember, err := handler.membersPort.Save(*member, organizationId)
	if err != nil {
		return nil, err
	}

	return savedMember, nil
}

type CreateRoleHandler struct {
	rolesPort         membershipPorts.RolesPort
	membersPort       membershipPorts.MembersPort
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler CreateRoleHandler) Execute(userId int, organizationId int, createRoleData membershipModels.CreateRoleData) (*membershipModels.Role, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(*organization, userId)
	if member == nil {
		return nil, ErrOrganizationNotFound
	}

	if !member.HasPermission(CreateRolesPermission.GetCode()) && organization.GetOwnerId() != userId {
		return nil, ErrCreateRolesPermission
	}

	role := membershipModels.NewRole(0, createRoleData.GetCode(), createRoleData.GetName(), createRoleData.GetPermissions())
	savedRole, err := handler.rolesPort.Save(role, organizationId)
	if err != nil {
		return nil, ErrSavingRole
	}

	return savedRole, nil
}

type UpdateRoleHandler struct {
	rolesPort         membershipPorts.RolesPort
	membersPort       membershipPorts.MembersPort
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler UpdateRoleHandler) Execute(userId int, organizationId int, roleId int, updateRoleData membershipModels.UpdateRoleData) (*membershipModels.Role, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(*organization, userId)
	if member == nil {
		return nil, ErrOrganizationNotFound
	}

	if !member.HasPermission(EditRolesPermission.GetCode()) && organization.GetOwnerId() != userId {
		return nil, ErrEditRolesPermission
	}

	role := handler.rolesPort.GetByOrganizationAndId(*organization, roleId)
	if role == nil {
		return nil, ErrRoleNotFound
	}

	role.SetPermissions(updateRoleData.GetPermissions())
	role.SetCode(updateRoleData.GetCode())
	role.SetName(updateRoleData.GetName())

	savedRole, err := handler.rolesPort.Save(*role, organizationId)
	if err != nil {
		return nil, ErrSavingRole
	}

	return savedRole, nil
}

type DeleteRoleHandler struct {
	rolesPort         membershipPorts.RolesPort
	membersPort       membershipPorts.MembersPort
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler DeleteRoleHandler) Execute(userId int, organizationId int, roleId int) error {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return ErrOrganizationNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(*organization, userId)
	if member == nil {
		return ErrOrganizationNotFound
	}

	if !member.HasPermission(DeleteRolesPermission.GetCode()) && organization.GetOwnerId() != userId {
		return ErrDeleteRolesPermission
	}

	role := handler.rolesPort.GetByOrganizationAndId(*organization, roleId)
	if role == nil {
		return ErrRoleNotFound
	}

	err := handler.rolesPort.Delete(*role)
	if err != nil {
		return ErrDeletingRole
	}

	return nil
}

type GetPermissionsHandler struct{}

func (handler GetPermissionsHandler) Execute() []membershipModels.Permission {
	return AllPermissions
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

func NewSetMemberRoleHandler(
	membersPort membershipPorts.MembersPort,
	rolesPort membershipPorts.RolesPort,
	organizationsPort organizationsPorts.OrganizationsPort,
) SetMemberRoleHandler {
	return SetMemberRoleHandler{
		membersPort:       membersPort,
		rolesPort:         rolesPort,
		organizationsPort: organizationsPort,
	}
}

func NewSetMemberPermissionsHandler(
	membersPort membershipPorts.MembersPort,
	organizationsPort organizationsPorts.OrganizationsPort,
) SetMemberPermissionsHandler {
	return SetMemberPermissionsHandler{
		membersPort:       membersPort,
		organizationsPort: organizationsPort,
	}
}

func NewCreateRoleHandler(
	rolesPort membershipPorts.RolesPort,
	membersPort membershipPorts.MembersPort,
	organizationsPort organizationsPorts.OrganizationsPort,
) CreateRoleHandler {
	return CreateRoleHandler{
		rolesPort:         rolesPort,
		membersPort:       membersPort,
		organizationsPort: organizationsPort,
	}
}

func NewUpdateRoleHandler(
	rolesPort membershipPorts.RolesPort,
	membersPort membershipPorts.MembersPort,
	organizationsPort organizationsPorts.OrganizationsPort,
) UpdateRoleHandler {
	return UpdateRoleHandler{
		rolesPort:         rolesPort,
		membersPort:       membersPort,
		organizationsPort: organizationsPort,
	}
}

func NewDeleteRoleHandler(
	rolesPort membershipPorts.RolesPort,
	membersPort membershipPorts.MembersPort,
	organizationsPort organizationsPorts.OrganizationsPort,
) DeleteRoleHandler {
	return DeleteRoleHandler{
		rolesPort:         rolesPort,
		membersPort:       membersPort,
		organizationsPort: organizationsPort,
	}
}

func NewGetOrganizationMembersHandler(organizationsPort organizationsPorts.OrganizationsPort) GetOrganizationMembersHandler {
	return GetOrganizationMembersHandler{
		organizationsPort: organizationsPort,
	}
}

func NewGetPermissionsHandler() GetPermissionsHandler {
	return GetPermissionsHandler{}
}
