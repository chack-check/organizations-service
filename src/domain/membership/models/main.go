package models

import (
	"slices"

	"github.com/chack-check/organizations-service/utils"
)

type PermissionCategory struct {
	name string
	code string
}

func (category PermissionCategory) GetName() string {
	return category.name
}

func (category PermissionCategory) GetCode() string {
	return category.code
}

type Permission struct {
	code     string
	name     string
	category *PermissionCategory
}

func (permission Permission) GetCode() string {
	return permission.code
}

func (permission Permission) GetName() string {
	return permission.name
}

func (permission Permission) GetCategory() *PermissionCategory {
	return permission.category
}

type Role struct {
	id          int
	code        string
	name        string
	permissions []Permission
}

func (role Role) GetId() int {
	return role.id
}

func (role Role) GetCode() string {
	return role.code
}

func (role *Role) SetCode(code string) {
	role.code = code
}

func (role Role) GetName() string {
	return role.name
}

func (role *Role) SetName(name string) {
	role.name = name
}

func (role Role) GetPermissions() []Permission {
	return role.permissions
}

func (role *Role) SetPermissions(permissions []Permission) {
	role.permissions = permissions
}

type Member struct {
	userId      int
	role        *Role
	permissions []Permission
}

func (member Member) GetUserId() int {
	return member.userId
}

func (member Member) GetRole() *Role {
	return member.role
}

func (member Member) GetPermissions() []Permission {
	return member.permissions
}

func (member *Member) SetRole(role *Role) {
	member.role = role
}

func (member *Member) HasPermission(permissionCode string) bool {
	for _, perm := range member.GetPermissions() {
		if perm.GetCode() == permissionCode {
			return true
		}
	}

	for _, perm := range member.GetRole().GetPermissions() {
		if perm.GetCode() == permissionCode {
			return true
		}
	}

	return false
}

func (member *Member) AddPermissions(permissions []Permission) {
	for _, permission := range permissions {
		addingPermissionsCodes := utils.GetArrayFieldValues(member.permissions, func(perm Permission) string { return perm.GetCode() })
		if slices.Contains(addingPermissionsCodes, permission.GetCode()) {
			continue
		}

		member.permissions = append(member.permissions, permission)
	}
}

func (member *Member) SetPermissions(permissions []Permission) {
	member.permissions = permissions
}

func (member *Member) RemovePermissions(permissions []Permission) {
	var newPermissions []Permission
	for _, permission := range newPermissions {
		deletingPermissionsCodes := utils.GetArrayFieldValues(permissions, func(perm Permission) string { return perm.GetCode() })
		if slices.Contains(deletingPermissionsCodes, permission.GetCode()) {
			continue
		}

		newPermissions = append(newPermissions, permission)
	}

	member.permissions = newPermissions
}

type CreateRoleData struct {
	code        string
	name        string
	permissions []Permission
}

func (model CreateRoleData) GetCode() string {
	return model.code
}

func (model CreateRoleData) GetName() string {
	return model.name
}

func (model CreateRoleData) GetPermissions() []Permission {
	return model.permissions
}

type UpdateRoleData struct {
	code        string
	name        string
	permissions []Permission
}

func (model UpdateRoleData) GetCode() string {
	return model.code
}

func (model UpdateRoleData) GetName() string {
	return model.name
}

func (model UpdateRoleData) GetPermissions() []Permission {
	return model.permissions
}

func NewPermissionCategory(name string, code string) PermissionCategory {
	return PermissionCategory{
		name: name,
		code: code,
	}
}

func NewPermission(code string, name string, category *PermissionCategory) Permission {
	return Permission{
		code:     code,
		name:     name,
		category: category,
	}
}

func NewRole(id int, code string, name string, permissions []Permission) Role {
	return Role{
		id:          id,
		code:        code,
		name:        name,
		permissions: permissions,
	}
}

func NewMember(userId int, role *Role, permissions []Permission) Member {
	return Member{
		userId:      userId,
		role:        role,
		permissions: permissions,
	}
}

func NewCreateRoleData(code string, name string, permissions []Permission) CreateRoleData {
	return CreateRoleData{
		code:        code,
		name:        name,
		permissions: permissions,
	}
}
