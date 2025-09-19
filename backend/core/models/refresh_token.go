package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	UserID uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	User   *User     `gorm:"foreignKey:UserID"`

	Expiry time.Time `gorm:"column:expiry;not null"`
}

func (RefreshToken) TableName() string { return "refresh_tokens" }

func (rft RefreshToken) Expired() bool {
	return time.Now().UTC().After(rft.Expiry)
}
