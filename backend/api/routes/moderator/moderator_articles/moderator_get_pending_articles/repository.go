package moderator_get_pending_articles

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	getPendingArticles(moderatorID uuid.UUID) ([]models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) getPendingArticles(moderatorID uuid.UUID) ([]models.Article, error) {
	var articles []models.Article

	if err := r.db.Where("moderator_id IS NULL AND redactor_id <> ? AND status = ?", moderatorID, models.ArticleStatusUnderReview).
		Order("created_at DESC").
		Select("id", "redactor_id", "reference", "title", "event_date", "updated_at", "category").
		Preload("Redactor", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "tag")
		}).
		Preload("Politicians", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name")
		}).
		Preload("Tags").
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}
