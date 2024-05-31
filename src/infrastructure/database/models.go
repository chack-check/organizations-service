package database

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type DBSavedFile struct {
	*gorm.Model
	ID                uint    `gorm:"primaryKey" json:"id"`
	OriginalUrl       string  `json:"original_url"`
	OriginalFilename  string  `json:"original_filename"`
	ConvertedUrl      *string `json:"converted_url"`
	ConvertedFilename *string `json:"converted_filename"`
}

type DBRole struct {
	*gorm.Model
	ID             uint           `gorm:"primaryKey" json:"id"`
	Code           string         `json:"code"`
	Name           string         `json:"name"`
	Permissions    pq.StringArray `gorm:"type:text[]" json:"permissions"`
	OrganizationID int            `json:"organization_id"`
	Organization   DBOrganization `gorm:"foreignKey:OrganizationID" json:"organization"`
}

type DBMember struct {
	*gorm.Model
	ID             uint           `gorm:"primaryKey" json:"id"`
	UserID         int            `json:"user_id"`
	RoleID         *int           `json:"role_id"`
	Role           *DBRole        `gorm:"foreignKey:RoleID" json:"role"`
	Permissions    pq.StringArray `gorm:"type:text[]" json:"permissions"`
	OrganizationID int            `json:"organization_id"`
	Organization   DBOrganization `gorm:"foreignKey:OrganizationID" json:"organization"`
}

type DBOrganization struct {
	*gorm.Model
	ID                 uint         `gorm:"primaryKey" json:"id"`
	Title              string       `json:"title"`
	Description        string       `json:"description"`
	MaxMembersCount    int          `json:"max_members_count"`
	MaxGroupChatsCount int          `json:"max_group_chats_count"`
	InviteTemplate     *string      `json:"invite_template"`
	Members            []DBMember   `gorm:"foreignKey:OrganizationID" json:"members"`
	OwnerID            int          `json:"owner_id"`
	AvatarID           *int         `json:"avatar_id"`
	Avatar             *DBSavedFile `gorm:"foreignKey:AvatarID" json:"avatar"`
}

type DBInvite struct {
	*gorm.Model
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()"`
	OrganizationID int            `json:"organization_id"`
	Organization   DBOrganization `gorm:"foreignKey:OrganizationID"`
	UserID         int
	RoleID         int
	Role           DBRole `gorm:"foreignKey:RoleID"`
	Status         string
}
