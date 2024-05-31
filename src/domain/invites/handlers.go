package invites

import (
	"fmt"

	invitesModels "github.com/chack-check/organizations-service/domain/invites/models"
	invitesPorts "github.com/chack-check/organizations-service/domain/invites/ports"
	membership "github.com/chack-check/organizations-service/domain/membership"
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	membershipPorts "github.com/chack-check/organizations-service/domain/membership/ports"
	organizationsPorts "github.com/chack-check/organizations-service/domain/organizations/ports"
	usersPorts "github.com/chack-check/organizations-service/domain/users/ports"
)

var (
	ErrUserNotFound           = fmt.Errorf("user not found")
	ErrOrganizationNotFound   = fmt.Errorf("organization not found")
	ErrRoleNotFound           = fmt.Errorf("role not found")
	ErrSavingInvite           = fmt.Errorf("error saving invite")
	ErrInviteNotFound         = fmt.Errorf("invite not found")
	ErrNotMember              = fmt.Errorf("you are not a member in organization")
	ErrInvitePermission       = fmt.Errorf("you have no permission to invite members")
	ErrDeleteInvitePermission = fmt.Errorf("you have no permission to delete invites")
	ErrViewInvitesPermission  = fmt.Errorf("you have no permissions to view organization invites")
)

type InviteMemberHandler struct {
	usersPort         usersPorts.UsersPort
	rolesPort         membershipPorts.RolesPort
	inviteEventsPort  invitesPorts.InviteEventsPort
	invitesPort       invitesPorts.InvitesPort
	organizationsPort organizationsPorts.OrganizationsPort
	membersPort       membershipPorts.MembersPort
}

func (handler InviteMemberHandler) Execute(currentUserId int, organizationId int, invitingUserId int, roleId int) (*invitesModels.Invite, error) {
	user := handler.usersPort.GetById(invitingUserId)
	if user == nil {
		return nil, ErrUserNotFound
	}

	organization := handler.organizationsPort.GetByIdForUser(organizationId, currentUserId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	role := handler.rolesPort.GetByOrganizationAndId(*organization, roleId)
	if role == nil {
		return nil, ErrRoleNotFound
	}

	currentMember := handler.membersPort.GetByOrganizationAndUserId(*organization, currentUserId)
	if currentMember == nil {
		return nil, ErrNotMember
	}
	if !currentMember.HasPermission(membership.InviteMembersPermission.GetCode()) && organization.GetOwnerId() != currentUserId {
		return nil, ErrInvitePermission
	}

	invite := invitesModels.NewInvite("", *organization, invitingUserId, *role, "open")
	savedInvite, err := handler.invitesPort.Save(invite)
	if err != nil {
		return nil, ErrSavingInvite
	}

	handler.inviteEventsPort.SendInviteCreated(*savedInvite)
	return savedInvite, nil
}

type InviteResponseHandler struct {
	invitesPort       invitesPorts.InvitesPort
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler InviteResponseHandler) Execute(userId int, inviteId string, accept bool) error {
	invite := handler.invitesPort.GetByIdForUser(inviteId, userId)
	if invite == nil {
		return ErrInviteNotFound
	}

	invite.SetStatus("closed")
	handler.invitesPort.Save(*invite)
	if !accept {
		return nil
	}

	role := invite.GetRole()
	member := membershipModels.NewMember(userId, &role, []membershipModels.Permission{})
	organization := invite.GetOrganization()
	organization.AddMembers([]membershipModels.Member{member})
	handler.organizationsPort.Save(organization)
	return nil
}

type CloseInviteHandler struct {
	invitesPort invitesPorts.InvitesPort
	membersPort membershipPorts.MembersPort
}

func (handler CloseInviteHandler) Execute(userId int, inviteId string) error {
	invite := handler.invitesPort.GetById(inviteId)
	if invite == nil {
		return ErrInviteNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(invite.GetOrganization(), userId)
	if member == nil {
		return ErrInviteNotFound
	}

	if !member.HasPermission(membership.CloseInvitesPermission.GetCode()) && invite.GetOrganization().GetOwnerId() != userId {
		return ErrDeleteInvitePermission
	}

	invite.SetStatus("closed")
	handler.invitesPort.Save(*invite)
	return nil
}

type GetActiveUserInvitesHandler struct {
	invitesPort invitesPorts.InvitesPort
}

func (handler GetActiveUserInvitesHandler) Execute(userId int) []invitesModels.Invite {
	invites := handler.invitesPort.GetAllForUser(userId)
	return invites
}

type GetOrganizationInvitesHandler struct {
	invitesPort       invitesPorts.InvitesPort
	membersPort       membershipPorts.MembersPort
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler GetOrganizationInvitesHandler) Execute(organizationId int, userId int) ([]invitesModels.Invite, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return []invitesModels.Invite{}, ErrOrganizationNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(*organization, userId)
	if member.HasPermission(membership.ViewInvitesPermission.GetCode()) && organization.GetOwnerId() != userId {
		return []invitesModels.Invite{}, ErrViewInvitesPermission
	}

	invites := handler.invitesPort.GetAllForOrganization(*organization)
	return invites, nil
}

func NewInviteMemberHandler(
	usersPort usersPorts.UsersPort,
	rolesPort membershipPorts.RolesPort,
	inviteEventsPort invitesPorts.InviteEventsPort,
	invitesPort invitesPorts.InvitesPort,
	organizationsPort organizationsPorts.OrganizationsPort,
	membersPort membershipPorts.MembersPort,
) InviteMemberHandler {
	return InviteMemberHandler{
		usersPort:         usersPort,
		rolesPort:         rolesPort,
		inviteEventsPort:  inviteEventsPort,
		invitesPort:       invitesPort,
		organizationsPort: organizationsPort,
		membersPort:       membersPort,
	}
}

func NewInviteResponseHandler(
	invitesPort invitesPorts.InvitesPort,
	organizationsPort organizationsPorts.OrganizationsPort,
) InviteResponseHandler {
	return InviteResponseHandler{
		invitesPort:       invitesPort,
		organizationsPort: organizationsPort,
	}
}

func NewCloseInviteHandler(
	invitesPort invitesPorts.InvitesPort,
	membersPort membershipPorts.MembersPort,
) CloseInviteHandler {
	return CloseInviteHandler{
		invitesPort: invitesPort,
		membersPort: membersPort,
	}
}

func NewGetActiveUserInvitesHandler(invitesPort invitesPorts.InvitesPort) GetActiveUserInvitesHandler {
	return GetActiveUserInvitesHandler{
		invitesPort: invitesPort,
	}
}

func NewGetOrganizationInvitesHandler(invitesPort invitesPorts.InvitesPort, membersPort membershipPorts.MembersPort, organizationsPort organizationsPorts.OrganizationsPort) GetOrganizationInvitesHandler {
	return GetOrganizationInvitesHandler{
		invitesPort:       invitesPort,
		membersPort:       membersPort,
		organizationsPort: organizationsPort,
	}
}
