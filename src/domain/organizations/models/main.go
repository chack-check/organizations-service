package models

import (
	"slices"

	filesModels "github.com/chack-check/organizations-service/domain/files/models"
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	"github.com/chack-check/organizations-service/utils"
)

type OrganizationStatuses string

var (
	OrganizationActive      OrganizationStatuses = "active"
	OrganizationDeactivated OrganizationStatuses = "deactivated"
)

type CreateOrganizationData struct {
	title          string
	description    string
	inviteTemplate *string
	avatar         *filesModels.UploadingFile
}

func (data CreateOrganizationData) GetTitle() string {
	return data.title
}

func (data CreateOrganizationData) GetDescription() string {
	return data.description
}

func (data CreateOrganizationData) GetInviteTemplate() *string {
	return data.inviteTemplate
}

func (data CreateOrganizationData) GetAvatar() *filesModels.UploadingFile {
	return data.avatar
}

type UpdateOrganizationData struct {
	title          string
	description    string
	inviteTemplate *string
}

func (data UpdateOrganizationData) GetTitle() string {
	return data.title
}

func (data UpdateOrganizationData) GetDescription() string {
	return data.description
}

func (data UpdateOrganizationData) GetInviteTemplate() *string {
	return data.inviteTemplate
}

type OrganizationConditions struct {
	maxOrganizationsCount int
	maxMembersCount       int
	maxGroupChatsCount    int
}

func (conditions OrganizationConditions) GetMaxOrganizationsCount() int {
	return conditions.maxOrganizationsCount
}

func (conditions OrganizationConditions) GetMaxMembersCount() int {
	return conditions.maxMembersCount
}

func (conditions OrganizationConditions) GetMaxGroupChatsCount() int {
	return conditions.maxGroupChatsCount
}

type Organization struct {
	id                 int
	title              string
	description        string
	maxMembersCount    int
	maxGroupChatsCount int
	status             OrganizationStatuses
	inviteTemplate     *string
	members            []membershipModels.Member
	ownerId            int
	avatar             *filesModels.SavedFile
}

func (organization Organization) GetId() int {
	return organization.id
}

func (organization Organization) GetTitle() string {
	return organization.title
}

func (organization *Organization) SetTitle(newTitle string) {
	organization.title = newTitle
}

func (organization Organization) GetDescription() string {
	return organization.description
}

func (organization *Organization) SetDescription(newDescription string) {
	organization.description = newDescription
}

func (organization Organization) GetMaxMembersCount() int {
	return organization.maxMembersCount
}

func (organization *Organization) SetMaxMembersCount(count int) {
	organization.maxMembersCount = count
}

func (organization Organization) GetStatus() OrganizationStatuses {
	return organization.status
}

func (organization *Organization) SetStatus(status OrganizationStatuses) {
	organization.status = status
}

func (organization Organization) GetMaxGroupChatsCount() int {
	return organization.maxGroupChatsCount
}

func (organization *Organization) SetMaxGroupChatsCount(count int) {
	organization.maxGroupChatsCount = count
}

func (organization Organization) GetInviteTemplate() *string {
	return organization.inviteTemplate
}

func (organization *Organization) SetInviteTemplate(template string) {
	organization.inviteTemplate = &template
}

func (organization Organization) GetMembers() []membershipModels.Member {
	return organization.members
}

func (organization *Organization) AddMembers(members []membershipModels.Member) {
	newMembers := organization.members
	for _, member := range members {
		membersIds := utils.GetArrayFieldValues(organization.members, func(v membershipModels.Member) int { return v.GetUserId() })
		if slices.Contains(membersIds, member.GetUserId()) {
			continue
		}

		newMembers = append(newMembers, member)
	}

	organization.members = newMembers
}

func (organization *Organization) SetMembers(members []membershipModels.Member) {
	organization.members = members
}

func (organization *Organization) RemoveMembers(members []membershipModels.Member) {
	var newMembers []membershipModels.Member
	for _, member := range organization.members {
		deletingMembersIds := utils.GetArrayFieldValues(members, func(item membershipModels.Member) int { return item.GetUserId() })
		if !slices.Contains(deletingMembersIds, member.GetUserId()) {
			newMembers = append(newMembers, member)
		}
	}

	organization.members = newMembers
}

func (organization Organization) GetOwnerId() int {
	return organization.ownerId
}

func (organization Organization) GetAvatar() *filesModels.SavedFile {
	return organization.avatar
}

func (organization *Organization) SetAvatar(avatar *filesModels.SavedFile) {
	organization.avatar = avatar
}

func NewOrganization(id int, title string, description string, maxMembersCount int, maxGroupChatsCount int, inviteTemplate *string, members []membershipModels.Member, ownerId int, avatar *filesModels.SavedFile) Organization {
	return Organization{
		id:                 id,
		title:              title,
		description:        description,
		maxMembersCount:    maxMembersCount,
		maxGroupChatsCount: maxGroupChatsCount,
		inviteTemplate:     inviteTemplate,
		members:            members,
		ownerId:            ownerId,
		avatar:             avatar,
	}
}

func NewOrganizationConditions(maxOrganizationsCount int, maxMembersCount int, maxGroupChatsCount int) OrganizationConditions {
	return OrganizationConditions{
		maxOrganizationsCount: maxOrganizationsCount,
		maxMembersCount:       maxMembersCount,
		maxGroupChatsCount:    maxGroupChatsCount,
	}
}

func NewCreateOrganizationData(
	title string,
	description string,
	inviteTemplate *string,
	avatar *filesModels.UploadingFile,
) CreateOrganizationData {
	return CreateOrganizationData{
		title:          title,
		description:    description,
		inviteTemplate: inviteTemplate,
		avatar:         avatar,
	}
}

func NewUpdateOrganizationData(
	title string,
	description string,
	inviteTemplate *string,
) UpdateOrganizationData {
	return UpdateOrganizationData{
		title:          title,
		description:    description,
		inviteTemplate: inviteTemplate,
	}
}
