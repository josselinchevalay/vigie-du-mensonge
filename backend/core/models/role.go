package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleName string

const (
	RoleNameAdmin     RoleName = "ADMIN"
	RoleNameModerator RoleName = "MODERATOR"
	RoleNameRedactor  RoleName = "REDACTOR"
)

type Role struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      RoleName       `gorm:"column:name;unique;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Role) TableName() string { return "roles" }
