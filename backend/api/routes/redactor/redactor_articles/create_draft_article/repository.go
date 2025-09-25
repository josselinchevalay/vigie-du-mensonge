package create_draft_article

import (
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	createArticle(article *models.Article) error
}

type repostiroy struct {
	db *gorm.DB
}

func (r *repostiroy) createArticle(article *models.Article) error {
	return r.db.Create(article).Error
}
