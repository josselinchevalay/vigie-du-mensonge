package find_published_article

import (
	"errors"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findPublishedArticle(articleID uuid.UUID) (*models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findPublishedArticle(articleID uuid.UUID) (*models.Article, error) {
	var article models.Article

	if err := r.db.Where("id = ? AND status = ?", articleID, models.ArticleStatusPublished).
		Preload("Sources").
		Preload("Tags").
		Preload("Politicians", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name")
		}).
		First(&article).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &article, nil
}
