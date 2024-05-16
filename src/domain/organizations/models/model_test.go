package models_test

import (
	"testing"

	filesModels "github.com/chack-check/organizations-service/domain/files/models"
	membershipModels "github.com/chack-check/organizations-service/domain/membership/models"
	organizationsModels "github.com/chack-check/organizations-service/domain/organizations/models"
)

func TestCreateOrganizationWithNullsNotRequired(t *testing.T) {
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	if organization.GetId() != 1 {
		t.Fatalf("Organization id %d != %d", organization.GetId(), 1)
	}
	if organization.GetTitle() != "title" {
		t.Fatalf("Organization title %s != %s", organization.GetTitle(), "title")
	}
	if organization.GetDescription() != "description" {
		t.Fatalf("Organization description %s != %s", organization.GetDescription(), "description")
	}
	if organization.GetMaxMembersCount() != 5 {
		t.Fatalf("Organization maxMembersCount %d != %d", organization.GetMaxMembersCount(), 5)
	}
	if organization.GetMaxGroupChatsCount() != 7 {
		t.Fatalf("Organization maxGroupChatCount %d != %d", organization.GetMaxGroupChatsCount(), 7)
	}
	if organization.GetMaxGroupChatsCount() != 7 {
		t.Fatalf("Organization maxGroupChatCount %d != %d", organization.GetMaxGroupChatsCount(), 7)
	}
	if organization.GetInviteTemplate() != nil {
		t.Fatalf("Organization inviteTemplate %v != %v", organization.GetInviteTemplate(), nil)
	}
	if len(organization.GetMembers()) != 0 {
		t.Fatalf("Organization members %v != %v", organization.GetMembers(), []membershipModels.Member{})
	}
	if len(organization.GetMembers()) != 0 {
		t.Fatalf("Organization members %v != %v", organization.GetMembers(), []membershipModels.Member{})
	}
	if organization.GetOwnerId() != 2 {
		t.Fatalf("Organization ownerId %d != %d", organization.GetOwnerId(), 2)
	}
	if organization.GetAvatar() != nil {
		t.Fatalf("Organization avatar %v != %v", organization.GetAvatar(), nil)
	}
}

func TestCreateOrganizationWithAllData(t *testing.T) {
	var template = "invite template"
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	var members []membershipModels.Member = []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	avatar := filesModels.NewSavedFile("originalUrl", "originalFilename", nil, nil)
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		&template,
		members,
		2,
		&avatar,
	)
	if organization.GetId() != 1 {
		t.Fatalf("Organization id %d != %d", organization.GetId(), 1)
	}
	if organization.GetTitle() != "title" {
		t.Fatalf("Organization title %s != %s", organization.GetTitle(), "title")
	}
	if organization.GetDescription() != "description" {
		t.Fatalf("Organization description %s != %s", organization.GetDescription(), "description")
	}
	if organization.GetMaxMembersCount() != 5 {
		t.Fatalf("Organization maxMembersCount %d != %d", organization.GetMaxMembersCount(), 5)
	}
	if organization.GetMaxGroupChatsCount() != 7 {
		t.Fatalf("Organization maxGroupChatCount %d != %d", organization.GetMaxGroupChatsCount(), 7)
	}
	if organization.GetMaxGroupChatsCount() != 7 {
		t.Fatalf("Organization maxGroupChatCount %d != %d", organization.GetMaxGroupChatsCount(), 7)
	}
	if organization.GetInviteTemplate() != &template {
		t.Fatalf("Organization inviteTemplate %v != %v", organization.GetInviteTemplate(), &template)
	}
	if len(organization.GetMembers()) != 1 || organization.GetMembers()[0].GetUserId() != members[0].GetUserId() {
		t.Fatalf("Organization members %v != %v", organization.GetMembers(), members)
	}
	if organization.GetOwnerId() != 2 {
		t.Fatalf("Organization ownerId %d != %d", organization.GetOwnerId(), 2)
	}
	if organization.GetAvatar() == nil || organization.GetAvatar().GetOriginalUrl() != avatar.GetOriginalUrl() {
		t.Fatalf("Organization avatar %v != %v", organization.GetAvatar(), &avatar)
	}
}

func TestSetTitle(t *testing.T) {
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	organization.SetTitle("new title")
	if organization.GetTitle() != "new title" {
		t.Fatalf("Set new title error: %s != \"%s\"", organization.GetTitle(), "new title")
	}
}

func TestSetDescription(t *testing.T) {
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	organization.SetDescription("new description")
	if organization.GetDescription() != "new description" {
		t.Fatalf("Set new description error: %s != \"%s\"", organization.GetDescription(), "new description")
	}
}

func TestSetMaxMembersCount(t *testing.T) {
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	organization.SetMaxMembersCount(25)
	if organization.GetMaxMembersCount() != 25 {
		t.Fatalf("Set new maxMembersCount error: %d != %d", organization.GetMaxMembersCount(), 25)
	}
}

func TestSetMaxGroupChatsCount(t *testing.T) {
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	organization.SetMaxGroupChatsCount(27)
	if organization.GetMaxGroupChatsCount() != 27 {
		t.Fatalf("Set new maxGroupChatsCount error: %d != %d", organization.GetMaxGroupChatsCount(), 27)
	}
}

func TestSetInviteTemplate(t *testing.T) {
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	organization.SetInviteTemplate("new invite template")
	if *organization.GetInviteTemplate() != "new invite template" {
		t.Fatalf("Set new inviteTemplate error: %s != %s", *organization.GetInviteTemplate(), "new invite template")
	}
}

func TestAddMembers(t *testing.T) {
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	newMembers := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization.AddMembers(newMembers)
	if len(organization.GetMembers()) != 1 || organization.GetMembers()[0].GetUserId() != newMembers[0].GetUserId() {
		t.Fatalf("Error add new members: %v != %v", organization.GetMembers(), newMembers)
	}
}

func TestAddMembersWithExistingMembers(t *testing.T) {
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	members := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		members,
		2,
		nil,
	)
	newMembers := []membershipModels.Member{
		membershipModels.NewMember(2, role, []membershipModels.Permission{}),
	}
	organization.AddMembers(newMembers)
	if len(organization.GetMembers()) != 2 {
		t.Fatalf("Error add new members: new count %d != %d", len(organization.GetMembers()), 2)
	}
	if organization.GetMembers()[1].GetUserId() != newMembers[0].GetUserId() {
		t.Fatalf("Error add new members: %+v != %+v", organization.GetMembers(), newMembers)
	}
}

func TestAddMembersWithSameMember(t *testing.T) {
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	members := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		members,
		2,
		nil,
	)
	newMembers := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization.AddMembers(newMembers)
	if len(organization.GetMembers()) != 1 {
		t.Fatalf("Error add new members: new count %d != %d", len(organization.GetMembers()), 1)
	}
	if organization.GetMembers()[0].GetUserId() != newMembers[0].GetUserId() {
		t.Fatalf("Error add new members: %+v != %+v", organization.GetMembers(), newMembers)
	}
}

func TestSetMembers(t *testing.T) {
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	newMembers := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization.SetMembers(newMembers)
	if len(organization.GetMembers()) != 1 {
		t.Fatalf("Error set members: count %d != %d", len(organization.GetMembers()), 1)
	}
	if organization.GetMembers()[0].GetUserId() != newMembers[0].GetUserId() {
		t.Fatalf("Error set members: %+v != %+v", organization.GetMembers(), newMembers)
	}
}

func TestSetMembersWithExisting(t *testing.T) {
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	members := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		members,
		2,
		nil,
	)
	newMembers := []membershipModels.Member{
		membershipModels.NewMember(2, role, []membershipModels.Permission{}),
	}
	organization.SetMembers(newMembers)
	if len(organization.GetMembers()) != 1 {
		t.Fatalf("Error set members: count %d != %d", len(organization.GetMembers()), 1)
	}
	if organization.GetMembers()[0].GetUserId() != newMembers[0].GetUserId() {
		t.Fatalf("Error set members: %+v != %+v", organization.GetMembers(), newMembers)
	}
}

func TestSetMembersEmpty(t *testing.T) {
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	members := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		members,
		2,
		nil,
	)
	var newMembers []membershipModels.Member
	organization.SetMembers(newMembers)
	if len(organization.GetMembers()) != 0 {
		t.Fatalf("Error set members: count %d != %d", len(organization.GetMembers()), 0)
	}
}

func TestRemoveMembers(t *testing.T) {
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	members := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		members,
		2,
		nil,
	)
	organization.RemoveMembers(members)
	if len(organization.GetMembers()) != 0 {
		t.Fatalf("Error deleting members: count %d != %d", len(organization.GetMembers()), 0)
	}
}

func TestRemoveMembersNotExisting(t *testing.T) {
	role := membershipModels.NewRole(1, "code", "name", []membershipModels.Permission{})
	members := []membershipModels.Member{
		membershipModels.NewMember(1, role, []membershipModels.Permission{}),
	}
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		members,
		2,
		nil,
	)
	deletingMembers := []membershipModels.Member{
		membershipModels.NewMember(2, role, []membershipModels.Permission{}),
	}
	organization.RemoveMembers(deletingMembers)
	if len(organization.GetMembers()) != 1 {
		t.Fatalf("Error deleting members: count %d != %d", len(organization.GetMembers()), 1)
	}
	if organization.GetMembers()[0].GetUserId() != members[0].GetUserId() {
		t.Fatalf("Error ")
	}
}

func TestSetAvatar(t *testing.T) {
	avatar := filesModels.NewSavedFile(
		"originalUrl",
		"originalFilename",
		nil,
		nil,
	)
	organization := organizationsModels.NewOrganization(
		1,
		"title",
		"description",
		5,
		7,
		nil,
		[]membershipModels.Member{},
		2,
		nil,
	)
	organization.SetAvatar(avatar)
	if organization.GetAvatar() == nil {
		t.Fatalf("Error set avatar: new avatar is nil")
	}
	if organization.GetAvatar().GetOriginalUrl() != avatar.GetOriginalUrl() {
		t.Fatalf("Error set avatar: %+v != %+v", organization.GetAvatar(), avatar)
	}
}
