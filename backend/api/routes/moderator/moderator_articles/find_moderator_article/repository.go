package find_moderator_article

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findArticlesByReference(reference uuid.UUID) ([]models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findArticlesByReference(reference uuid.UUID) ([]models.Article, error) {
	var articles []models.Article

	if err := r.db.Where("reference = ?", reference).
		Order("created_at DESC").
		Preload("Sources").
		Preload("Tags").
		Preload("Review", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Moderator", func(db *gorm.DB) *gorm.DB {
				return db.Select("id", "tag")
			}).Select("moderator_id", "article_id", "notes", "decision")
		}).
		Preload("Moderator", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "tag")
		}).
		Preload("Politicians", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "first_name", "last_name")
		}).
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}
