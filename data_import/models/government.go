package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Government struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	PrimeMinisterID uuid.UUID

	Reference int
	StartDate time.Time
	EndDate   sql.NullTime

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (Government) TableName() string {
	return "governments"
}
