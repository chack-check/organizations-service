package database

import (
	filesModels "github.com/chack-check/organizations-service/domain/files/models"
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

type DatabaseOrganizationsAdapter struct{}

func (adapter DatabaseOrganizationsAdapter) createMember(member membershipModels.Member, organizationId int) DBMember {
	savingMember := ModelMemberToDB(member, organizationId)
	DatabaseConnection.Save(&savingMember)
	return savingMember
}

func (adapter DatabaseOrganizationsAdapter) getOrCreateMember(member membershipModels.Member, organizationId int) DBMember {
	dbMember := &DBMember{UserID: member.GetUserId(), OrganizationID: organizationId}
	DatabaseConnection.First(dbMember)
	dbMember.OrganizationID = organizationId
	if dbMember.ID != 0 {
		return *dbMember
	}

	return adapter.createMember(member, organizationId)
}

func (adapter DatabaseOrganizationsAdapter) getOrCreateAvatar(avatar filesModels.SavedFile) DBSavedFile {
	dbFile := &DBSavedFile{OriginalUrl: avatar.GetOriginalUrl()}
	DatabaseConnection.First(dbFile)
	dbFile.OriginalFilename = avatar.GetOriginalFilename()
	dbFile.ConvertedUrl = avatar.GetConvertedUrl()
	dbFile.ConvertedFilename = avatar.GetConvertedFilename()
	DatabaseConnection.Save(dbFile)
	return *dbFile
}

func (adapter DatabaseOrganizationsAdapter) Save(organization organizationsModels.Organization) (*organizationsModels.Organization, error) {
	var dbAvatar *DBSavedFile
	if avatar := organization.GetAvatar(); avatar != nil {
		file := adapter.getOrCreateAvatar(*avatar)
		dbAvatar = &file
	}

	savingOrganization := ModelOrganizationToDB(organization, []DBMember{}, dbAvatar)
	DatabaseConnection.Save(&savingOrganization)
	var members []DBMember
	for _, member := range organization.GetMembers() {
		if organization.GetId() == 0 {
			members = append(members, adapter.createMember(member, int(savingOrganization.ID)))
		} else {
			members = append(members, adapter.getOrCreateMember(member, int(savingOrganization.ID)))
		}
	}

	savingOrganization.Members = members
	response := DBOrganizationToModel(savingOrganization)
	return &response, nil
}

func (adapter DatabaseOrganizationsAdapter) GetForUser(userId int, includeNotActive bool) []organizationsModels.Organization {
	var organizations []DBOrganization
	DatabaseConnection.Preload("Members", "user_id = ?", userId).Find(&organizations)
	var models []organizationsModels.Organization
	for _, organization := range organizations {
		model := DBOrganizationToModel(organization)
		models = append(models, model)
	}

	return models
}

func (adapter DatabaseOrganizationsAdapter) GetByIdForUser(id int, userId int, includeNotActive bool) *organizationsModels.Organization {
	organization := &DBOrganization{ID: uint(id)}
	DatabaseConnection.Preload("Members", "user_id = ?", userId).First(organization)
	if organization.ID == 0 {
		return nil
	}

	response := DBOrganizationToModel(*organization)
	return &response
}

func (adapter DatabaseOrganizationsAdapter) GetOpenCountForUser(userId int) int {
	var count int64
	DatabaseConnection.Model(&DBOrganization{}).Preload("Members", "user_id = ?", userId).Count(&count)
	if count != 0 {
		return int(count)
	} else {
		return 0
	}
}

type DatabaseMembersAdapter struct{}

func (adapter DatabaseMembersAdapter) GetByUserId(userId int) *membershipModels.Member {
	member := &DBMember{UserID: userId}
	DatabaseConnection.First(member)
	if member.ID == 0 {
		return nil
	}

	response := DBMemberToModel(*member)
	return &response
}

func (adapter DatabaseMembersAdapter) Save(member membershipModels.Member, organizationId int) (*membershipModels.Member, error) {
	dbMember := ModelMemberToDB(member, organizationId)
	DatabaseConnection.Save(&dbMember)
	response := DBMemberToModel(dbMember)
	return &response, nil
}

type DatabaseRolesAdapter struct{}

func (adapter DatabaseRolesAdapter) GetByOrganization(organization organizationsModels.Organization) []membershipModels.Role {
	var roles []DBRole
	DatabaseConnection.Where("organization_id = ?", organization.GetId()).Find(&roles)
	var rolesModels []membershipModels.Role
	for _, role := range roles {
		roleModel := DBRoleToModel(role)
		rolesModels = append(rolesModels, roleModel)
	}

	return rolesModels
}

func (adapter DatabaseRolesAdapter) Save(role membershipModels.Role, organizationId int) (*membershipModels.Role, error) {
	dbRole := ModelRoleToDB(role, organizationId)
	DatabaseConnection.Save(&dbRole)
	response := DBRoleToModel(dbRole)
	return &response, nil
}

type DatabaseFilesAdapter struct{}

func (adapter DatabaseFilesAdapter) SaveFile(file filesModels.UploadingFile) (*filesModels.SavedFile, error) {
	originalFile := file.GetOriginal()
	var convertedUrl *string
	var convertedFilename *string
	if converted := file.GetConverted(); converted != nil {
		url := converted.GetUrl()
		convertedUrl = &url
		filename := converted.GetFilename()
		convertedFilename = &filename
	}

	savingFile := &DBSavedFile{
		OriginalUrl:       originalFile.GetUrl(),
		OriginalFilename:  originalFile.GetFilename(),
		ConvertedUrl:      convertedUrl,
		ConvertedFilename: convertedFilename,
	}
	DatabaseConnection.Save(savingFile)
	response := DBSavedFileToModel(*savingFile)
	return &response, nil
}

func (adapter DatabaseFilesAdapter) ValidateUploadingFile(file filesModels.UploadingFile) bool {
	return true
}
