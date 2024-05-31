package organizations

import (
	"fmt"
	"slices"

	filesModels "github.com/chack-check/organizations-service/domain/files/models"
	filesPorts "github.com/chack-check/organizations-service/domain/files/ports"
	"github.com/chack-check/organizations-service/domain/membership"
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	membershipPorts "github.com/chack-check/organizations-service/domain/membership/ports"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
	organizationsPorts "github.com/chack-check/organizations-service/domain/organizations/ports"
)

var (
	ErrMaxOrganizationsLimitReached = fmt.Errorf("max organizations count limit reached")
	ErrIncorrectFile                = fmt.Errorf("incorrect uploading file")
	ErrSavingOrganization           = fmt.Errorf("error saving organization")
	ErrOrganizationNotFound         = fmt.Errorf("organization not found")
	ErrDeleteMembersPermission      = fmt.Errorf("you have no permission to delete organization members")
	ErrRoleNotFound                 = fmt.Errorf("role not found")
	ErrSavingMember                 = fmt.Errorf("error saving member")
	ErrOrganizationNotOwner         = fmt.Errorf("you are not an owner of organization")
	ErrMemberNotFound               = fmt.Errorf("member with this user id not found")
	ErrUpdateOrganizationPermission = fmt.Errorf("you have no permissions to update organization")
	ErrSavingFile                   = fmt.Errorf("error saving file")
)

type CreateOrganizationHandler struct {
	organizationsPort      organizationsPorts.OrganizationsPort
	organizationEventsPort organizationsPorts.OrganizationEventsPort
	rolesPort              membershipPorts.RolesPort
	filesPort              filesPorts.FilesPort
	subscriptionsPort      organizationsPorts.SubscriptionsPort
}

func (handler CreateOrganizationHandler) saveAvatar(file *filesModels.UploadingFile) (*filesModels.SavedFile, error) {
	var savedAvatar *filesModels.SavedFile
	if file != nil {
		isAvatarValid := handler.filesPort.ValidateUploadingFile(*file)
		if !isAvatarValid {
			return nil, ErrIncorrectFile
		}

		saved, err := handler.filesPort.SaveFile(*file)
		if err != nil {
			return nil, err
		}

		savedAvatar = saved
	}

	return savedAvatar, nil
}

func (handler CreateOrganizationHandler) createMemberRole(organizationId int) (*membershipModels.Role, error) {
	role := membershipModels.NewRole(0, "member", "Member", []membershipModels.Permission{})
	savedRole, err := handler.rolesPort.Save(role, organizationId)
	if err != nil {
		return nil, err
	}

	return savedRole, nil
}

func (handler CreateOrganizationHandler) createAdminRole(organizationId int) (*membershipModels.Role, error) {
	role := membershipModels.NewRole(0, "admin", "Admin", membership.AdminPermissions)
	savedRole, err := handler.rolesPort.Save(role, organizationId)
	if err != nil {
		return nil, err
	}

	return savedRole, nil
}

func (handler CreateOrganizationHandler) createOwnerRole(organizationId int) (*membershipModels.Role, error) {
	role := membershipModels.NewRole(0, "owner", "Owner", membership.AllPermissions)
	savedRole, err := handler.rolesPort.Save(role, organizationId)
	if err != nil {
		return nil, err
	}

	return savedRole, nil
}

func (handler CreateOrganizationHandler) createBaseRoles(organizationId int) (*membershipModels.Role, error) {
	_, err := handler.createMemberRole(organizationId)
	if err != nil {
		return nil, err
	}

	_, err = handler.createAdminRole(organizationId)
	if err != nil {
		return nil, err
	}

	ownerRole, err := handler.createOwnerRole(organizationId)
	if err != nil {
		return nil, err
	}

	return ownerRole, nil
}

func (handler CreateOrganizationHandler) Execute(userId int, createData organizationsModels.CreateOrganizationData) (*organizationsModels.Organization, error) {
	conditions, err := handler.subscriptionsPort.GetUserOrganizationConditions(userId)
	if err != nil {
		return nil, err
	}

	openOrganizationsCount := handler.organizationsPort.GetOpenCountForUser(userId)
	if openOrganizationsCount >= conditions.GetMaxOrganizationsCount() {
		return nil, ErrMaxOrganizationsLimitReached
	}

	savedAvatar, err := handler.saveAvatar(createData.GetAvatar())
	if err != nil {
		return nil, err
	}

	organization := organizationsModels.NewOrganization(
		0,
		createData.GetTitle(),
		createData.GetDescription(),
		conditions.GetMaxMembersCount(),
		conditions.GetMaxGroupChatsCount(),
		createData.GetInviteTemplate(),
		[]membershipModels.Member{},
		userId,
		savedAvatar,
	)
	savedOrganization, err := handler.organizationsPort.Save(organization)
	if err != nil {
		return nil, ErrSavingOrganization
	}

	ownerRole, err := handler.createBaseRoles(savedOrganization.GetId())
	if err != nil {
		return nil, err
	}
	member := membershipModels.NewMember(userId, ownerRole, []membershipModels.Permission{})
	savedOrganization.SetMembers([]membershipModels.Member{member})
	savedOrganization, err = handler.organizationsPort.Save(*savedOrganization)
	if err != nil {
		return nil, err
	}

	handler.organizationEventsPort.SendOrganizationCreated(*savedOrganization)
	return savedOrganization, nil
}

type HasUserOrganizationsHandler struct {
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler HasUserOrganizationsHandler) Execute(userId int) bool {
	activeOrganizations := handler.organizationsPort.GetForUser(userId, false)
	return len(activeOrganizations) > 0
}

type GetUserOrganizationsHandler struct {
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler GetUserOrganizationsHandler) Execute(userId int) []organizationsModels.Organization {
	activeOrganizations := handler.organizationsPort.GetForUser(userId, false)
	return activeOrganizations
}

type UpdateOrganizationHandler struct {
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler UpdateOrganizationHandler) Execute(userId int, organizationId int, updateData organizationsModels.UpdateOrganizationData) (*organizationsModels.Organization, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	organization.SetTitle(updateData.GetTitle())
	organization.SetDescription(updateData.GetDescription())
	if updateData.GetInviteTemplate() != nil {
		organization.SetInviteTemplate(*updateData.GetInviteTemplate())
	}

	savedOrganization, err := handler.organizationsPort.Save(*organization)
	if err != nil {
		return nil, ErrSavingOrganization
	}

	return savedOrganization, nil
}

type DeleteMembersHandler struct {
	membersPort       membershipPorts.MembersPort
	organizationsPort organizationsPorts.OrganizationsPort
}

func (handler DeleteMembersHandler) Execute(userId int, organizationId int, deletingMembers []int) (*organizationsModels.Organization, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(*organization, userId)
	if member == nil {
		return nil, ErrOrganizationNotFound
	}

	if !member.HasPermission(membership.RemoveMembersPermission.GetCode()) && organization.GetOwnerId() != userId {
		return nil, ErrDeleteMembersPermission
	}

	var organizationDeletingMembers []membershipModels.Member
	for _, member := range organization.GetMembers() {
		if member.GetUserId() == userId {
			continue
		}

		if slices.Contains(deletingMembers, member.GetUserId()) {
			organizationDeletingMembers = append(organizationDeletingMembers, member)
		}
	}

	organization.RemoveMembers(organizationDeletingMembers)
	savedOrganization, err := handler.organizationsPort.Save(*organization)
	if err != nil {
		return nil, ErrSavingOrganization
	}

	return savedOrganization, nil
}

type DeactivateOrganizationHandler struct {
	organizationsPort      organizationsPorts.OrganizationsPort
	organizationEventsPort organizationsPorts.OrganizationEventsPort
}

func (handler DeactivateOrganizationHandler) Execute(userId int, organizationId int) error {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return ErrOrganizationNotFound
	}

	if organization.GetOwnerId() != userId {
		return ErrOrganizationNotOwner
	}

	organization.SetStatus(organizationsModels.OrganizationDeactivated)
	savedOrganization, err := handler.organizationsPort.Save(*organization)
	if err != nil {
		return ErrSavingOrganization
	}

	handler.organizationEventsPort.SendOrganizationDeactivated(*savedOrganization)
	return nil
}

type ReactivateOrganization struct {
	organizationsPort      organizationsPorts.OrganizationsPort
	organizationEventsPort organizationsPorts.OrganizationEventsPort
}

func (handler ReactivateOrganization) Execute(userId int, organizationId int) (*organizationsModels.Organization, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, true)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	if organization.GetOwnerId() != userId {
		return nil, ErrOrganizationNotOwner
	}

	organization.SetStatus(organizationsModels.OrganizationActive)
	savedOrganization, err := handler.organizationsPort.Save(*organization)
	if err != nil {
		return nil, ErrSavingOrganization
	}

	handler.organizationEventsPort.SendOrganizationActivated(*savedOrganization)
	return savedOrganization, nil
}

type UpdateOrganizationAvatarHandler struct {
	filesPort              filesPorts.FilesPort
	organizationsPort      organizationsPorts.OrganizationsPort
	membersPort            membershipPorts.MembersPort
	organizationEventsPort organizationsPorts.OrganizationEventsPort
}

func (handler UpdateOrganizationAvatarHandler) Execute(userId int, organizationId int, newAvatar *filesModels.UploadingFile) (*organizationsModels.Organization, error) {
	organization := handler.organizationsPort.GetByIdForUser(organizationId, userId, false)
	if organization == nil {
		return nil, ErrOrganizationNotFound
	}

	member := handler.membersPort.GetByOrganizationAndUserId(*organization, userId)
	if member == nil {
		return nil, ErrMemberNotFound
	}

	if member.HasPermission(membership.EditOrganizationPermission.GetCode()) && organization.GetOwnerId() != userId {
		return nil, ErrUpdateOrganizationPermission
	}

	if newAvatar == nil {
		organization.SetAvatar(nil)
	} else {
		if !handler.filesPort.ValidateUploadingFile(*newAvatar) {
			return nil, ErrIncorrectFile
		}
		savedFile, err := handler.filesPort.SaveFile(*newAvatar)
		if err != nil {
			return nil, ErrSavingFile
		}

		organization.SetAvatar(savedFile)
	}

	savedOrganization, err := handler.organizationsPort.Save(*organization)
	if err != nil {
		return nil, ErrSavingOrganization
	}

	handler.organizationEventsPort.SendOrganizationChanged(*savedOrganization)
	return savedOrganization, nil
}

func NewCreateOrganizationHandler(
	organizationsPort organizationsPorts.OrganizationsPort,
	organizationEventsPort organizationsPorts.OrganizationEventsPort,
	rolesPort membershipPorts.RolesPort,
	filesPort filesPorts.FilesPort,
	subscriptionsPort organizationsPorts.SubscriptionsPort,
) CreateOrganizationHandler {
	return CreateOrganizationHandler{
		organizationsPort:      organizationsPort,
		organizationEventsPort: organizationEventsPort,
		rolesPort:              rolesPort,
		filesPort:              filesPort,
		subscriptionsPort:      subscriptionsPort,
	}
}

func NewHasUserOrganizationsHandler(organizationsPort organizationsPorts.OrganizationsPort) HasUserOrganizationsHandler {
	return HasUserOrganizationsHandler{organizationsPort: organizationsPort}
}

func NewGetUserOrganizationsHandler(organizationsPort organizationsPorts.OrganizationsPort) GetUserOrganizationsHandler {
	return GetUserOrganizationsHandler{organizationsPort: organizationsPort}
}

func NewUpdateOrganizationHandler(organizationsPort organizationsPorts.OrganizationsPort) UpdateOrganizationHandler {
	return UpdateOrganizationHandler{organizationsPort: organizationsPort}
}

func NewDeactivateOrganizationHandler(
	organizationsPort organizationsPorts.OrganizationsPort,
	organizationEventsPort organizationsPorts.OrganizationEventsPort,
) DeactivateOrganizationHandler {
	return DeactivateOrganizationHandler{
		organizationsPort:      organizationsPort,
		organizationEventsPort: organizationEventsPort,
	}
}

func NewReactivateOrganizationHandler(
	organizationsPort organizationsPorts.OrganizationsPort,
	organizationEventsPort organizationsPorts.OrganizationEventsPort,
) ReactivateOrganization {
	return ReactivateOrganization{
		organizationsPort:      organizationsPort,
		organizationEventsPort: organizationEventsPort,
	}
}

func NewUpdateOrganizationAvatarHandler(
	filesPort filesPorts.FilesPort,
	organizationsPort organizationsPorts.OrganizationsPort,
	membersPort membershipPorts.MembersPort,
	organizationEventsPort organizationsPorts.OrganizationEventsPort,
) UpdateOrganizationAvatarHandler {
	return UpdateOrganizationAvatarHandler{
		filesPort:              filesPort,
		organizationsPort:      organizationsPort,
		membersPort:            membersPort,
		organizationEventsPort: organizationEventsPort,
	}
}
