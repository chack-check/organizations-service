package database

import (
	filesModels "github.com/chack-check/organizations-service/domain/files/models"
	"github.com/chack-check/organizations-service/domain/membership"
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
	"github.com/lib/pq"
)

func DBPermissionToModel(permission string) *membershipModels.Permission {
	for _, perm := range membership.AllPermissions {
		if perm.GetCode() == permission {
			return &perm
		}
	}

	return nil
}

func DBRoleToModel(role DBRole) membershipModels.Role {
	var permissions []membershipModels.Permission
	for _, permission := range role.Permissions {
		perm := DBPermissionToModel(permission)
		if perm != nil {
			permissions = append(permissions, *perm)
		}
	}

	return membershipModels.NewRole(
		int(role.ID),
		role.Code,
		role.Name,
		permissions,
	)
}

func ModelRoleToDB(role membershipModels.Role, organizationId int) DBRole {
	var permissions pq.StringArray
	for _, permission := range role.GetPermissions() {
		permissions = append(permissions, permission.GetCode())
	}

	return DBRole{
		ID:             uint(role.GetId()),
		Code:           role.GetCode(),
		Name:           role.GetName(),
		Permissions:    permissions,
		OrganizationID: organizationId,
	}
}

func DBMemberToModel(member DBMember) membershipModels.Member {
	var permissions []membershipModels.Permission
	for _, permission := range member.Permissions {
		perm := DBPermissionToModel(permission)
		if perm != nil {
			permissions = append(permissions, *perm)
		}
	}

	return membershipModels.NewMember(
		member.UserID,
		DBRoleToModel(member.Role),
		permissions,
	)
}

func ModelMemberToDB(member membershipModels.Member, organizationId int) DBMember {
	var permissions pq.StringArray
	for _, permission := range member.GetPermissions() {
		permissions = append(permissions, permission.GetCode())
	}

	return DBMember{
		UserID:         member.GetUserId(),
		OrganizationID: organizationId,
		Role:           ModelRoleToDB(member.GetRole(), organizationId),
		Permissions:    permissions,
	}
}

func DBSavedFileToModel(file DBSavedFile) filesModels.SavedFile {
	return filesModels.NewSavedFile(
		file.OriginalUrl,
		file.OriginalFilename,
		file.ConvertedUrl,
		file.ConvertedFilename,
	)
}

func DBOrganizationToModel(organization DBOrganization) organizationsModels.Organization {
	var members []membershipModels.Member
	for _, member := range organization.Members {
		model := DBMemberToModel(member)
		members = append(members, model)
	}

	var avatar *filesModels.SavedFile
	if organization.Avatar != nil {
		file := DBSavedFileToModel(*organization.Avatar)
		avatar = &file
	}

	return organizationsModels.NewOrganization(
		int(organization.ID),
		organization.Title,
		organization.Description,
		organization.MaxMembersCount,
		organization.MaxGroupChatsCount,
		organization.InviteTemplate,
		members,
		organization.OwnerID,
		avatar,
	)
}

func ModelOrganizationToDB(organization organizationsModels.Organization, members []DBMember, avatar *DBSavedFile) DBOrganization {
	return DBOrganization{
		ID:                 uint(organization.GetId()),
		Title:              organization.GetTitle(),
		Description:        organization.GetDescription(),
		MaxMembersCount:    organization.GetMaxMembersCount(),
		MaxGroupChatsCount: organization.GetMaxGroupChatsCount(),
		InviteTemplate:     organization.GetInviteTemplate(),
		Members:            members,
		OwnerID:            organization.GetOwnerId(),
		Avatar:             avatar,
	}
}
