package models

import "github.com/google/uuid"

// ArticleTag represents the article_tags table (composite primary key)

type ArticleTag struct {
	ArticleID uuid.UUID `gorm:"column:article_id;type:uuid;primaryKey"`
	Tag       string    `gorm:"column:tag;primaryKey"`
}

func (ArticleTag) TableName() string { return "article_tags" }
