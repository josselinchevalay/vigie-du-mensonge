package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ArticleReview represents the article_reviews table

type ArticleReview struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ArticleID   uuid.UUID `gorm:"column:article_id;type:uuid;not null"`
	ModeratorID uuid.UUID `gorm:"column:moderator_id;type:uuid;not null"`

	Decision ArticleStatus `gorm:"column:decision;not null"`
	Notes    string        `gorm:"column:notes;not null"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ArticleReview) TableName() string { return "article_reviews" }
