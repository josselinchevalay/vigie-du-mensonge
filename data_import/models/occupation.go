package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Occupation struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	PoliticianID uuid.UUID
	GovernmentID *uuid.UUID

	PresidentialReference *int

	Code      string
	Title     string
	StartDate time.Time
	EndDate   sql.NullTime

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (Occupation) TableName() string {
	return "occupations"
}
