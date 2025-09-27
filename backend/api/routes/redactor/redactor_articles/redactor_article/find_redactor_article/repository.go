package find_redactor_article

import (
	"fmt"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findRedactorArticlesByReference(redactorID, reference uuid.UUID) (models.Article, []models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findRedactorArticlesByReference(redactorID, reference uuid.UUID) (models.Article, []models.Article, error) {
	var articles []models.Article

	if err := r.db.Where("redactor_id = ? AND reference = ?", redactorID, reference).
		Order("created_at DESC").
		Preload("Sources").
		Preload("Tags").
		Preload("Politicians", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name")
		}).
		Find(&articles).Error; err != nil {
		return models.Article{}, nil, err
	}

	if len(articles) == 0 {
		return models.Article{}, nil, fmt.Errorf("article not found")
	}

	if len(articles) > 1 {
		return articles[0], articles[1:], nil
	}

	return articles[0], nil, nil
}
