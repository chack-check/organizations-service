package organizations

import (
	"testing"

	"github.com/chack-check/organizations-service/domain/files"
	filesModels "github.com/chack-check/organizations-service/domain/files/models"
	"github.com/chack-check/organizations-service/domain/membership"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

func TestCreateOrganization(t *testing.T) {
	handler := NewCreateOrganizationHandler(
		MockOrganizationsAdapter{},
		MockOrganizationEventsAdapter{},
		membership.MockRolesAdapter{},
		files.MockFilesAdapter{},
		MockSubscriptionsAdapter{},
	)
	data := organizationsModels.NewCreateOrganizationData(
		"creating organization",
		"description",
		nil,
		nil,
	)
	savedOrganization, err := handler.Execute(1, data)
	if err != nil {
		t.Fatalf("Error creationg organization: %v", err)
	}
	if savedOrganization.GetTitle() != data.GetTitle() {
		t.Fatalf("Error creating organization: title %s != %s", savedOrganization.GetTitle(), data.GetTitle())
	}
	if savedOrganization.GetDescription() != data.GetDescription() {
		t.Fatalf("Error creating organization: description %s != %s", savedOrganization.GetDescription(), data.GetDescription())
	}
	if savedOrganization.GetInviteTemplate() != data.GetInviteTemplate() {
		t.Fatalf("Error creating organization: invite template %v != %v", savedOrganization.GetInviteTemplate(), data.GetInviteTemplate())
	}
	if savedOrganization.GetAvatar() != nil {
		t.Fatalf("Error creating organization: avatar %+v != %+v", savedOrganization.GetAvatar(), data.GetAvatar())
	}
	if len(savedOrganization.GetMembers()) != 1 {
		t.Fatalf("Error creating organization: members count %d != %d", len(savedOrganization.GetMembers()), 1)
	}
	if savedOrganization.GetMembers()[0].GetUserId() != 1 {
		t.Fatalf("Error creating organization: member id %d != %d", savedOrganization.GetMembers()[0].GetUserId(), 1)
	}
	if savedOrganization.GetOwnerId() != 1 {
		t.Fatalf("Error creating organization: owner id %d != %d", savedOrganization.GetOwnerId(), 1)
	}
}

func TestCreateOrganizationWithoutNils(t *testing.T) {
	handler := NewCreateOrganizationHandler(
		MockOrganizationsAdapter{},
		MockOrganizationEventsAdapter{},
		membership.MockRolesAdapter{},
		files.MockFilesAdapter{},
		MockSubscriptionsAdapter{},
	)
	template := "invite template"
	convertedMeta := filesModels.NewUploadingFileMeta(
		"convertedUrl",
		"convertedFilename",
		"convertedSignature",
		"convertedFiletype",
	)
	uploadingFile := filesModels.NewUploadingFile(
		filesModels.NewUploadingFileMeta(
			"originalUrl",
			"originalFilename",
			"originalSignature",
			"originalFiletype",
		),
		&convertedMeta,
	)
	data := organizationsModels.NewCreateOrganizationData(
		"creating organization",
		"description",
		&template,
		&uploadingFile,
	)
	savedOrganization, err := handler.Execute(1, data)
	if err != nil {
		t.Fatalf("Error creationg organization: %v", err)
	}
	if savedOrganization.GetTitle() != data.GetTitle() {
		t.Fatalf("Error creating organization: title %s != %s", savedOrganization.GetTitle(), data.GetTitle())
	}
	if savedOrganization.GetDescription() != data.GetDescription() {
		t.Fatalf("Error creating organization: description %s != %s", savedOrganization.GetDescription(), data.GetDescription())
	}
	if savedOrganization.GetInviteTemplate() != data.GetInviteTemplate() {
		t.Fatalf("Error creating organization: invite template %v != %v", savedOrganization.GetInviteTemplate(), data.GetInviteTemplate())
	}
	if savedOrganization.GetAvatar() == nil || *savedOrganization.GetAvatar().GetConvertedUrl() != uploadingFile.GetConverted().GetUrl() {
		t.Fatalf("Error creating organization: avatar %+v != %+v", savedOrganization.GetAvatar(), data.GetAvatar())
	}
	if len(savedOrganization.GetMembers()) != 1 {
		t.Fatalf("Error creating organization: members count %d != %d", len(savedOrganization.GetMembers()), 1)
	}
	if savedOrganization.GetMembers()[0].GetUserId() != 1 {
		t.Fatalf("Error creating organization: member id %d != %d", savedOrganization.GetMembers()[0].GetUserId(), 1)
	}
	if savedOrganization.GetOwnerId() != 1 {
		t.Fatalf("Error creating organization: owner id %d != %d", savedOrganization.GetOwnerId(), 1)
	}
}

func TestHasUserOrganizations(t *testing.T) {
	handler := NewHasUserOrganizationsHandler(
		MockOrganizationsAdapter{},
	)
	hasOrganizations := handler.Execute(1)
	if !hasOrganizations {
		t.Fatalf("Error checking has organizations")
	}
}

func TestGetUserOrganizations(t *testing.T) {
	handler := NewGetUserOrganizationsHandler(
		MockOrganizationsAdapter{},
	)
	userOrganizations := handler.Execute(1)
	if len(userOrganizations) != 1 {
		t.Fatalf("Error fetching user organizations: length %d != %d", len(userOrganizations), 1)
	}
}

func TestUpdateOrganization(t *testing.T) {
	handler := NewUpdateOrganizationHandler(
		MockOrganizationsAdapter{},
	)
	template := "new template"
	updateData := organizationsModels.NewUpdateOrganizationData(
		"new title",
		"new description",
		&template,
	)
	savedOrganization, err := handler.Execute(1, 1, updateData)
	if err != nil {
		t.Fatalf("Error updating organization: %v", err)
	}

	if savedOrganization.GetTitle() != "new title" {
		t.Fatalf("Error updating organization: title %s != %s", savedOrganization.GetTitle(), "new title")
	}

	if *savedOrganization.GetInviteTemplate() != template {
		t.Fatalf("Error updating organization: invite template %s != %s", *savedOrganization.GetInviteTemplate(), template)
	}
}

func TestDeactivateOrganization(t *testing.T) {
	handler := NewDeactivateOrganizationHandler(
		MockOrganizationsAdapter{},
		MockOrganizationEventsAdapter{},
	)
	err := handler.Execute(2, 1)
	if err != nil {
		t.Fatalf("Error deactivating organization: %v", err)
	}
}

func TestDeactivateOrganizationWithNotOwner(t *testing.T) {
	handler := NewDeactivateOrganizationHandler(
		MockOrganizationsAdapter{},
		MockOrganizationEventsAdapter{},
	)
	err := handler.Execute(1, 1)
	if err != ErrOrganizationNotOwner {
		t.Fatalf("Error deactivating organization: %v", err)
	}
}

func TestReactivateOrganization(t *testing.T) {
	handler := NewReactivateOrganizationHandler(
		MockOrganizationsAdapter{},
		MockOrganizationEventsAdapter{},
	)
	organization, err := handler.Execute(2, 1)
	if err != nil {
		t.Fatalf("Error reactivating organization: %v", err)
	}

	if organization == nil {
		t.Fatalf("Error reactivating organization: organization is nil")
	}
}

func TestReactivateOrganizationWithNotOwner(t *testing.T) {
	handler := NewReactivateOrganizationHandler(
		MockOrganizationsAdapter{},
		MockOrganizationEventsAdapter{},
	)
	_, err := handler.Execute(1, 1)
	if err != ErrOrganizationNotOwner {
		t.Fatalf("Error reactivating organization: %v", err)
	}
}

func TestUpdateOrganizationAvatarFile(t *testing.T) {
	handler := NewUpdateOrganizationAvatarHandler(
		files.MockFilesAdapter{},
		MockOrganizationsAdapter{},
		membership.MockMembersAdapter{},
		MockOrganizationEventsAdapter{},
	)
	originalMeta := filesModels.NewUploadingFileMeta(
		"url",
		"filename.jpg",
		"signature",
		"avatar",
	)
	convertedMeta := filesModels.NewUploadingFileMeta(
		"convertedUrl",
		"convertedFilename.jpg",
		"signature",
		"avatar",
	)
	uploading := filesModels.NewUploadingFile(originalMeta, &convertedMeta)
	organization, err := handler.Execute(1, 1, &uploading)
	if err != nil {
		t.Fatalf("Error updating organization avatar: %v", err)
	}

	if organization.GetAvatar() == nil {
		t.Fatalf("Error updating organization avatar: avatar is nil")
	}
	if organization.GetAvatar().GetOriginalUrl() != "url" {
		t.Fatalf("Error updating organization avatar: original url %s != %s", organization.GetAvatar().GetOriginalUrl(), "url")
	}
	if organization.GetAvatar().GetOriginalFilename() != "filename.jpg" {
		t.Fatalf("Error updating organization avatar: original filename %s != %s", organization.GetAvatar().GetOriginalFilename(), "filename.jpg")
	}
	if organization.GetAvatar().GetConvertedUrl() == nil {
		t.Fatalf("Error updating organization avatar: converted file url is nil")
	}
	if *organization.GetAvatar().GetConvertedUrl() != "convertedUrl" {
		t.Fatalf("Error updating organization avatar: url %s != %s", *organization.GetAvatar().GetConvertedUrl(), "convertedUrl")
	}
	if *organization.GetAvatar().GetConvertedFilename() != "convertedFilename.jpg" {
		t.Fatalf("Error updating organization avatar: filename %s != %s", *organization.GetAvatar().GetConvertedFilename(), "convertedFilename.jpg")
	}
}
