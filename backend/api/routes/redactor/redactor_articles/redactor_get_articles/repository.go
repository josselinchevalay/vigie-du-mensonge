package redactor_get_articles

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	getArticlesByRedactorID(redactorID uuid.UUID) ([]models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) getArticlesByRedactorID(redactorID uuid.UUID) ([]models.Article, error) {
	var articles []models.Article

	if err := r.db.Where("redactor_id = ? AND status <> ?", redactorID, models.ArticleStatusArchived).
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
