package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents the users table

type User struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email         string         `gorm:"column:email;unique;not null"`
	EmailVerified bool           `gorm:"column:email_verified;not null;default:false"`
	Password      string         `gorm:"column:password;not null"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (User) TableName() string { return "users" }
