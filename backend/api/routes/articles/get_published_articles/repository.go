package get_published_articles

import (
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	getPublishedArticles() ([]models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) getPublishedArticles() ([]models.Article, error) {
	var articles []models.Article

	if err := r.db.Where("status = ?", models.ArticleStatusPublished).
		Order("created_at DESC").
		Select("id", "title", "event_date", "updated_at", "category").
		Preload("Politicians", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name")
		}).
		Preload("Tags").
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}
