package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ArticleStatus and ArticleCategory are string enums enforced in DB; we keep as strings but provide helper consts

type ArticleStatus string

func (as ArticleStatus) Valid() bool {
	switch as {
	case ArticleStatusDraft, ArticleStatusPublished, ArticleStatusArchived, ArticleStatusUnderReview, ArticleStatusChangeRequested:
		return true
	}
	return false
}

type ArticleCategory string

func (ac ArticleCategory) Valid() bool {
	switch ac {
	case ArticleCategoryLie, ArticleCategoryFalsehood:
		return true
	}
	return false
}

const (
	ArticleStatusDraft           ArticleStatus = "DRAFT"
	ArticleStatusPublished       ArticleStatus = "PUBLISHED"
	ArticleStatusArchived        ArticleStatus = "ARCHIVED"
	ArticleStatusUnderReview     ArticleStatus = "UNDER_REVIEW"
	ArticleStatusChangeRequested ArticleStatus = "CHANGE_REQUESTED"

	ArticleCategoryLie       ArticleCategory = "LIE"
	ArticleCategoryFalsehood ArticleCategory = "FALSEHOOD"
)

// Article represents the articles table

type Article struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	ArticlePoliticians []*ArticlePolitician `gorm:"foreignKey:ArticleID"`
	Politicians        []*Politician        `gorm:"many2many:article_politicians;"`
	Tags               []*ArticleTag        `gorm:"foreignKey:ArticleID"`
	Sources            []*ArticleSource     `gorm:"foreignKey:ArticleID"`
	Review             *ArticleReview       `gorm:"foreignKey:ArticleID"`

	RedactorID  uuid.UUID  `gorm:"column:redactor_id;type:uuid;not null"`
	ModeratorID *uuid.UUID `gorm:"column:moderator_id;type:uuid"`

	Status    ArticleStatus   `gorm:"column:status;type:text;not null"`
	Category  ArticleCategory `gorm:"column:category;type:text;not null"`
	Title     string          `gorm:"column:title;not null"`
	Body      string          `gorm:"column:body;not null"`
	EventDate time.Time       `gorm:"column:event_date;not null"`
	Reference uuid.UUID       `gorm:"column:reference;not null"`
	Major     int16           `gorm:"column:major;not null"`
	Minor     int16           `gorm:"column:minor;not null"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Article) TableName() string { return "articles" }
