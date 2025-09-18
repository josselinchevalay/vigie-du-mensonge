package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Occupation represents the occupations table

type Occupation struct {
	ID                    uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PoliticianID          uuid.UUID      `gorm:"column:politician_id;type:uuid;not null"`
	GovernmentID          *uuid.UUID     `gorm:"column:government_id;type:uuid"`
	PresidentialReference *int16         `gorm:"column:presidential_reference;unique"`
	Code                  string         `gorm:"column:code;not null"`
	Title                 string         `gorm:"column:title;not null"`
	StartDate             time.Time      `gorm:"column:start_date;not null"`
	EndDate               sql.NullTime   `gorm:"column:end_date"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Occupation) TableName() string { return "occupations" }
