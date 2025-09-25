package get_redactor_articles

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	getArticlesByAuthorId(authorID uuid.UUID) ([]models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) getArticlesByAuthorId(authorID uuid.UUID) ([]models.Article, error) {
	var articles []models.Article

	if err := r.db.Where("author_id = ?", authorID).
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
