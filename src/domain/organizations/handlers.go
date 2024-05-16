package organizations

import (
	"fmt"

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
	member := membershipModels.NewMember(userId, *ownerRole, []membershipModels.Permission{})
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
