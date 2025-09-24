package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
