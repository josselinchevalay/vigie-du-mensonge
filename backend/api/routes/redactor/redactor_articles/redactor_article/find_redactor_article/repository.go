package find_redactor_article

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findRedactorArticle(articleID, redactorID uuid.UUID) (models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findRedactorArticle(articleID, redactorID uuid.UUID) (models.Article, error) {
	var article models.Article

	if err := r.db.Where("id = ? AND redactor_id = ?", articleID, redactorID).
		Preload("Sources").
		Preload("Tags").
		Preload("Politicians", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name")
		}).
		First(&article).Error; err != nil {
		return models.Article{}, err
	}

	sources := make([]string, len(article.Sources))
	for i, source := range article.Sources {
		sources[i] = source.URL
	}

	return article, nil
}
