package models

import "github.com/google/uuid"

// ArticleSource represents the article_sources table (composite primary key)

type ArticleSource struct {
	ArticleID uuid.UUID `gorm:"column:article_id;type:uuid;primaryKey"`
	URL       string    `gorm:"column:url;primaryKey"`
}

func (ArticleSource) TableName() string { return "article_sources" }
