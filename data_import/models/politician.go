package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Politician struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	FirstName string
	LastName  string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (Politician) TableName() string {
	return "politicians"
}
