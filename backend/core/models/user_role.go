package models

import (
	"github.com/google/uuid"
)

// UserRole represents the user_roles join table with composite primary key

type UserRole struct {
	UserID uuid.UUID `gorm:"column:user_id;type:uuid;primaryKey"`
	RoleID uuid.UUID `gorm:"column:role_id;type:uuid;primaryKey"`
}

func (UserRole) TableName() string { return "user_roles" }
