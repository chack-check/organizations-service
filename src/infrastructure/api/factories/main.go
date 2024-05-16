package factories

import (
	filesModels "github.com/chack-check/organizations-service/domain/files/models"
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
	"github.com/chack-check/organizations-service/infrastructure/api/graph/model"
)

func UploadingFileMetaToModel(meta model.UploadingFileMeta) filesModels.UploadingFileMeta {
	return filesModels.NewUploadingFileMeta(
		meta.URL,
		meta.Filename,
		meta.Signature,
		meta.SystemFiletype.String(),
	)
}

func UploadingFileToModel(file model.UploadingFile) filesModels.UploadingFile {
	var converted *filesModels.UploadingFileMeta
	if file.Converted != nil {
		meta := UploadingFileMetaToModel(*file.Converted)
		converted = &meta
	}

	return filesModels.NewUplaodingFile(
		UploadingFileMetaToModel(*file.Original),
		converted,
	)
}

func SavedFileModelToResponse(file filesModels.SavedFile) model.SavedFile {
	return model.SavedFile{
		OriginalURL:       file.GetOriginalUrl(),
		OriginalFilename:  file.GetOriginalFilename(),
		ConvertedURL:      file.GetConvertedUrl(),
		ConvertedFilename: file.GetConvertedFilename(),
	}
}

func CreateOrganizationDataToModel(data model.CreateOrganizationData) organizationsModels.CreateOrganizationData {
	var avatar *filesModels.UploadingFile
	if data.Avatar != nil {
		file := UploadingFileToModel(*data.Avatar)
		avatar = &file
	}

	return organizationsModels.NewCreateOrganizationData(
		data.Title,
		data.Description,
		data.InviteTemplate,
		avatar,
	)
}

func PermissionCategoryModelToResponse(category membershipModels.PermissionCategory) model.PermissionCategory {
	return model.PermissionCategory{
		Code: category.GetCode(),
		Name: category.GetName(),
	}
}

func PermissionModelToResponse(permission membershipModels.Permission) model.Permission {
	var category *model.PermissionCategory
	if categoryModel := permission.GetCategory(); categoryModel != nil {
		response := PermissionCategoryModelToResponse(*categoryModel)
		category = &response
	}

	return model.Permission{
		Code:     permission.GetCode(),
		Name:     permission.GetName(),
		Category: category,
	}
}

func RoleModelToResponse(role membershipModels.Role) model.Role {
	var permissions []*model.Permission
	for _, perm := range role.GetPermissions() {
		response := PermissionModelToResponse(perm)
		permissions = append(permissions, &response)
	}

	return model.Role{
		ID:          role.GetId(),
		Code:        role.GetCode(),
		Name:        role.GetName(),
		Permissions: permissions,
	}
}

func MemberModelToResponse(member membershipModels.Member) model.Member {
	var permissions []*model.Permission
	for _, perm := range member.GetPermissions() {
		response := PermissionModelToResponse(perm)
		permissions = append(permissions, &response)
	}

	role := RoleModelToResponse(member.GetRole())
	return model.Member{
		UserID:      member.GetUserId(),
		Role:        &role,
		Permissions: permissions,
	}
}

func OrganizationModelToResponse(organization organizationsModels.Organization) model.Organization {
	var members []*model.Member
	for _, member := range organization.GetMembers() {
		response := MemberModelToResponse(member)
		members = append(members, &response)
	}

	var avatar *model.SavedFile
	if file := organization.GetAvatar(); file != nil {
		response := SavedFileModelToResponse(*file)
		avatar = &response
	}

	return model.Organization{
		ID:                 organization.GetId(),
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
