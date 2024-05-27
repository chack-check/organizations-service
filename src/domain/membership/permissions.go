package membership

import (
	"github.com/chack-check/organizations-service/domain/membership/models"
)

var OrganizationCategory = models.NewPermissionCategory(
	"Organization",
	"organization",
)

var EditOrganizationPermission = models.NewPermission(
	"edit_organization",
	"Edit organization",
	&OrganizationCategory,
)

var CloseOrganizationPermission = models.NewPermission(
	"close_organization",
	"Close organization",
	&OrganizationCategory,
)

var MembersCategory = models.NewPermissionCategory(
	"Members",
	"members",
)

var ViewInvitesPermission = models.NewPermission(
	"view_invites",
	"View invites",
	&MembersCategory,
)

var CloseInvitesPermission = models.NewPermission(
	"close_invites",
	"Close invites",
	&MembersCategory,
)

var InviteMembersPermission = models.NewPermission(
	"invite_members",
	"Invite members",
	&MembersCategory,
)

var RemoveMembersPermission = models.NewPermission(
	"remove_members",
	"Remove members",
	&MembersCategory,
)

var SetMembersRolesPermission = models.NewPermission(
	"set_members_roles",
	"Set members roles",
	&MembersCategory,
)

var SetMembersPermissions = models.NewPermission(
	"set_members_permissions",
	"Set members permissions",
	&MembersCategory,
)

var RolesCategory = models.NewPermissionCategory(
	"Roles",
	"roles",
)

var CreateRolesPermission = models.NewPermission(
	"create_roles",
	"Create roles",
	&RolesCategory,
)

var DeleteRolesPermission = models.NewPermission(
	"delete_roles",
	"Delete roles",
	&RolesCategory,
)

var EditRolesPermission = models.NewPermission(
	"edit_roles",
	"Edit roles",
	&RolesCategory,
)

var AllPermissions = []models.Permission{
	EditOrganizationPermission,
	CloseOrganizationPermission,
	InviteMembersPermission,
	RemoveMembersPermission,
	SetMembersRolesPermission,
	SetMembersPermissions,
	CreateRolesPermission,
	DeleteRolesPermission,
	EditRolesPermission,
}

var AdminPermissions = []models.Permission{
	InviteMembersPermission,
	RemoveMembersPermission,
	SetMembersRolesPermission,
	SetMembersPermissions,
	CreateRolesPermission,
	DeleteRolesPermission,
	EditRolesPermission,
}
