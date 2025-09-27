package get_pending_articles

import (
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	getPendingArticles() ([]models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) getPendingArticles() ([]models.Article, error) {
	var articles []models.Article

	if err := r.db.Where("moderator_id IS NULL AND status = ?", models.ArticleStatusUnderReview).
		Order("created_at DESC").
		Select("id", "reference", "title", "event_date", "updated_at", "category", "status").
		Preload("Politicians", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name")
		}).
		Preload("Tags").
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}
