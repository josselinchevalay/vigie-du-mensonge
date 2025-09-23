package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Government represents the governments table

type Government struct {
	ID              uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PrimeMinisterID uuid.UUID    `gorm:"column:prime_minister_id;type:uuid;not null"`
	Reference       int16        `gorm:"column:reference;not null;unique"`
	StartDate       time.Time    `gorm:"column:start_date;not null"`
	EndDate         sql.NullTime `gorm:"column:end_date"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Government) TableName() string { return "governments" }
