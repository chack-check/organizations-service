package models

import "testing"

func TestCreatePermissionCategory(t *testing.T) {
	permissionCategory := NewPermissionCategory(
		"name",
		"code",
	)

	if permissionCategory.GetCode() != "code" {
		t.Fatalf("Error creating permission category: code %s != %s", permissionCategory.GetCode(), "code")
	}
	if permissionCategory.GetName() != "name" {
		t.Fatalf("Error creating permission category: name %s != %s", permissionCategory.GetName(), "name")
	}
}

func TestCreatePermission(t *testing.T) {
	permission := NewPermission(
		"permcode",
		"permname",
		nil,
	)

	if permission.GetCode() != "permcode" {
		t.Fatalf("Error creating permission: code %s != %s", permission.GetCode(), "permcode")
	}
	if permission.GetName() != "permname" {
		t.Fatalf("Error creating permission: name %s != %s", permission.GetName(), "permname")
	}
	if permission.GetCategory() != nil {
		t.Fatalf("Error creating permission: category %+v != %v", permission.GetCategory(), nil)
	}
}

func TestCreatePermissionWithCategory(t *testing.T) {
	permissionCategory := NewPermissionCategory(
		"name",
		"code",
	)
	permission := NewPermission(
		"permcode",
		"permname",
		&permissionCategory,
	)

	if permission.GetCode() != "permcode" {
		t.Fatalf("Error creating permission: code %s != %s", permission.GetCode(), "permcode")
	}
	if permission.GetName() != "permname" {
		t.Fatalf("Error creating permission: name %s != %s", permission.GetName(), "permname")
	}
	if permission.GetCategory().GetCode() != permissionCategory.GetCode() {
		t.Fatalf("Error creating permission: category %+v != %+v", permission.GetCategory(), permissionCategory)
	}
}

func TestCreateRole(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)

	if role.GetCode() != "code" {
		t.Fatalf("Error creating role: code %s != %s", role.GetCode(), "code")
	}

	if role.GetName() != "name" {
		t.Fatalf("Error creating role: name %s != %s", role.GetName(), "name")
	}
	if len(role.GetPermissions()) != 0 {
		t.Fatalf("Error creating role: permissions count %d != %d", len(role.GetPermissions()), 0)
	}
}

func TestCreateRoleWithPermissions(t *testing.T) {
	permissions := []Permission{
		NewPermission("code", "name", nil),
	}
	role := NewRole(
		1,
		"code",
		"name",
		permissions,
	)
	if role.GetCode() != "code" {
		t.Fatalf("Error creating role: code %s != %s", role.GetCode(), "code")
	}

	if role.GetName() != "name" {
		t.Fatalf("Error creating role: name %s != %s", role.GetName(), "name")
	}
	if len(role.GetPermissions()) != 1 {
		t.Fatalf("Error creating role: permissions count %d != %d", len(role.GetPermissions()), 1)
	}
	if role.GetPermissions()[0].GetCode() != permissions[0].GetCode() {
		t.Fatalf("Error creating role: permissions %+v != %+v", role.GetPermissions(), permissions)
	}
}

func TestSetRolePermissions(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	permissions := []Permission{
		NewPermission("code", "name", nil),
	}
	role.SetPermissions(permissions)
	if len(role.GetPermissions()) != 1 {
		t.Fatalf("Error set role permissions: permissions count %d != %d", len(role.GetPermissions()), 1)
	}
	if role.GetPermissions()[0].GetCode() != permissions[0].GetCode() {
		t.Fatalf("Error set role permissions: permissions %+v != %+v", role.GetPermissions(), permissions)
	}
}

func TestSetRolePermissionsEmpty(t *testing.T) {
	permissions := []Permission{
		NewPermission("code", "name", nil),
	}
	role := NewRole(
		1,
		"code",
		"name",
		permissions,
	)
	role.SetPermissions([]Permission{})
	if len(role.GetPermissions()) != 0 {
		t.Fatalf("Error set role permissions: permissions count %d != %d", len(role.GetPermissions()), 0)
	}
}

func TestCreateMember(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	member := NewMember(
		1,
		role,
		[]Permission{},
	)

	if member.GetUserId() != 1 {
		t.Fatalf("Error creating member: user id %d != %d", member.GetUserId(), 1)
	}
	if member.GetRole().GetCode() != "code" {
		t.Fatalf("Error creating member: role %+v != %+v", member.GetRole(), role)
	}
	if len(member.GetPermissions()) != 0 {
		t.Fatalf("Error creating member: permissions count %d != %d", len(member.GetPermissions()), 0)
	}
}

func TestMemberHasPermission(t *testing.T) {
	permissions := []Permission{
		NewPermission("permcode", "permname", nil),
	}
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	member := NewMember(
		1,
		role,
		permissions,
	)

	if !member.HasPermission("permcode") {
		t.Fatalf("Error checking has member permission: permission = %s, member permissions: %+v", "permcode", member.GetPermissions())
	}
}

func TestMemberHasPermissionWithRolePermissions(t *testing.T) {
	permissions := []Permission{
		NewPermission("permcode", "permname", nil),
	}
	role := NewRole(
		1,
		"code",
		"name",
		permissions,
	)
	member := NewMember(
		1,
		role,
		[]Permission{},
	)

	if !member.HasPermission("permcode") {
		t.Fatalf("Error checking has member permission: permission = %s, role permissions: %+v", "permcode", role.GetPermissions())
	}
}

func TestMemberHasPermissionWithAnotherPermission(t *testing.T) {
	permissions := []Permission{
		NewPermission("permcode", "permname", nil),
	}
	role := NewRole(
		1,
		"code",
		"name",
		permissions,
	)
	member := NewMember(
		1,
		role,
		[]Permission{},
	)

	if member.HasPermission("anotherperm") {
		t.Fatalf("Error checking has member permission: permission = %s, role permissions: %+v", "anotherperm", role.GetPermissions())
	}
}

func TestSetMemberRole(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	member := NewMember(
		1,
		role,
		[]Permission{},
	)
	newRole := NewRole(
		2,
		"newcode",
		"newname",
		[]Permission{},
	)
	member.SetRole(newRole)

	if member.GetRole().GetCode() != newRole.GetCode() {
		t.Fatalf("Error set member role: %+v != %+v", member.GetRole(), newRole)
	}
}

func TestAddPermisions(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	member := NewMember(
		1,
		role,
		[]Permission{},
	)
	addingPermissions := []Permission{
		NewPermission("code", "name", nil),
	}

	member.AddPermissions(addingPermissions)

	if len(member.GetPermissions()) != 1 {
		t.Fatalf("Error adding member permissions: count %d != %d", len(member.GetPermissions()), 1)
	}
	if member.GetPermissions()[0].GetCode() != addingPermissions[0].GetCode() {
		t.Fatalf("Error adding member permissions: %+v != %+v", member.GetPermissions(), addingPermissions)
	}
}

func TestAddPermisionsWithExisting(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	permissions := []Permission{
		NewPermission("code", "name", nil),
	}
	member := NewMember(
		1,
		role,
		permissions,
	)

	addingPermissions := []Permission{
		NewPermission("code1", "name1", nil),
	}
	member.AddPermissions(addingPermissions)

	if len(member.GetPermissions()) != 2 {
		t.Fatalf("Error adding member permissions: count %d != %d", len(member.GetPermissions()), 2)
	}
	if member.GetPermissions()[1].GetCode() != addingPermissions[0].GetCode() {
		t.Fatalf("Error adding member permissions: %+v, adding permissions: %+v", member.GetPermissions(), addingPermissions)
	}
}

func TestAddPermisionsWithSame(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	permissions := []Permission{
		NewPermission("code", "name", nil),
	}
	member := NewMember(
		1,
		role,
		permissions,
	)

	addingPermissions := []Permission{
		NewPermission("code", "name", nil),
	}
	member.AddPermissions(addingPermissions)

	if len(member.GetPermissions()) != 1 {
		t.Fatalf("Error adding member permissions: count %d != %d", len(member.GetPermissions()), 1)
	}
}

func TestSetPermissions(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	member := NewMember(
		1,
		role,
		[]Permission{},
	)

	newPermissions := []Permission{
		NewPermission("code1", "name1", nil),
	}
	member.SetPermissions(newPermissions)

	if len(member.GetPermissions()) != 1 {
		t.Fatalf("Error adding member permissions: count %d != %d", len(member.GetPermissions()), 1)
	}
	if member.GetPermissions()[0].GetCode() != newPermissions[0].GetCode() {
		t.Fatalf("Error adding member permissions: %+v, adding permissions: %+v", member.GetPermissions(), newPermissions)
	}
}

func TestSetPermissionsEmpty(t *testing.T) {
	role := NewRole(
		1,
		"code",
		"name",
		[]Permission{},
	)
	permissions := []Permission{
		NewPermission("code", "name", nil),
	}
	member := NewMember(
		1,
		role,
		permissions,
	)

	member.SetPermissions([]Permission{})

	if len(member.GetPermissions()) != 0 {
		t.Fatalf("Error adding member permissions: count %d != %d", len(member.GetPermissions()), 0)
	}
}
