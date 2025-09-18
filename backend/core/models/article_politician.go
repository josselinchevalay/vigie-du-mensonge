package models

import "github.com/google/uuid"

// ArticlePolitician represents the article_politicians table (composite primary key)

type ArticlePolitician struct {
	ArticleID    uuid.UUID `gorm:"column:article_id;type:uuid;primaryKey"`
	PoliticianID uuid.UUID `gorm:"column:politician_id;type:uuid;primaryKey"`
}

func (ArticlePolitician) TableName() string { return "article_politicians" }
