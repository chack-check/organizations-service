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
	uploadingFile := filesModels.NewUplaodingFile(
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
