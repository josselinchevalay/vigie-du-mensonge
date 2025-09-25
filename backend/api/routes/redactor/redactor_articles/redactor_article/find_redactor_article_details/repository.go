package find_redactor_article_details

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findArticleBodyAndSources(articleID, authorID uuid.UUID) (string, []string, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findArticleBodyAndSources(articleID, authorID uuid.UUID) (string, []string, error) {
	var article models.Article

	if err := r.db.Where("id = ? AND author_id = ?", articleID, authorID).
		Preload("Sources").
		Select("id", "body").
		First(&article).Error; err != nil {
		return "", nil, err
	}

	sources := make([]string, len(article.Sources))
	for i, source := range article.Sources {
		sources[i] = source.URL
	}

	return article.Body, sources, nil
}
