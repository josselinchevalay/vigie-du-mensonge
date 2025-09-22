package models

import (
	"time"

	"github.com/google/uuid"
)

type UserTokenCategory string

const (
	UserTokenCategoryPassword UserTokenCategory = "PASSWORD"
	UserTokenCategoryEmail    UserTokenCategory = "EMAIL"
	UserTokenCategoryRefresh  UserTokenCategory = "REFRESH"
)

type UserToken struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	UserID uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	User   *User     `gorm:"foreignKey:UserID"`

	Category UserTokenCategory `gorm:"column:category;not null"`
	Hash     string            `gorm:"column:hash;not null"`
	Expiry   time.Time         `gorm:"column:expiry;not null"`
}

func (UserToken) TableName() string { return "user_tokens" }

func (rft UserToken) Expired() bool {
	return time.Now().UTC().After(rft.Expiry)
}
