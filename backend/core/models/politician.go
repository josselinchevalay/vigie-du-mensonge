package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Politician represents the politicians table
// CHECK constraints are enforced at DB level; here we map columns and types
// deleted_at is handled as gorm.DeletedAt for soft deletes
// timestamps are in TIMESTAMPTZ in DB, mapped to time.Time
// UUIDs mapped to uuid.UUID
// Text mapped to string

type Politician struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	LastName  string         `gorm:"column:last_name;not null"`
	FirstName string         `gorm:"column:first_name;not null"`
	ImageUrl  sql.NullString `gorm:"column:image_url"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Politician) TableName() string { return "politicians" }
