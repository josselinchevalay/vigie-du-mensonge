package save_redactor_article

import (
	"fmt"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findArticle(articleID, redactorID uuid.UUID) (models.Article, error)
	archiveOldVersionAndCreateNew(article *models.Article) error
	createArticle(article *models.Article) error
	updateArticle(article *models.Article) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findArticle(articleID, redactorID uuid.UUID) (models.Article, error) {
	var article models.Article
	if err := r.db.Where("id = ? AND redactor_id = ?", articleID, redactorID).
		Select("id", "redactor_id", "status", "reference", "minor", "major").
		First(&article).Error; err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func (r *repository) archiveOldVersionAndCreateNew(article *models.Article) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Article{}).
			Where("id = ?", article.ID).
			Update("status", models.ArticleStatusArchived).Error; err != nil {
			return fmt.Errorf("failed to archive old version: %v", err)
		}

		article.ID = uuid.New()
		if err := tx.Create(article).Error; err != nil {
			return fmt.Errorf("failed to create new version: %v", err)
		}
		return nil
	})
}

func (r *repository) createArticle(article *models.Article) error {
	return r.db.Create(article).Error
}

func (r *repository) updateArticle(article *models.Article) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Article{}).
			Where("id = ?", article.ID).
			Updates(map[string]any{
				"title":      article.Title,
				"body":       article.Body,
				"event_date": article.EventDate,
				"category":   article.Category,
			}).Error; err != nil {
			return fmt.Errorf("failed to update article: %v", err)
		}

		if err := tx.Where("article_id = ?", article.ID).
			Delete(&models.ArticleTag{}).Error; err != nil {
			return fmt.Errorf("failed to delete old tags: %v", err)
		}
		if len(article.Tags) > 0 {
			if err := tx.Create(article.Tags).Error; err != nil {
				return fmt.Errorf("failed to create new tags: %v", err)
			}
		}

		if err := tx.Where("article_id = ?", article.ID).
			Delete(&models.ArticleSource{}).Error; err != nil {
			return fmt.Errorf("failed to delete old sources: %v", err)
		}
		if len(article.Sources) > 0 {
			if err := tx.Create(article.Sources).Error; err != nil {
				return fmt.Errorf("failed to create new sources: %v", err)
			}
		}

		if err := tx.Where("article_id = ?", article.ID).
			Delete(&models.ArticlePolitician{}).Error; err != nil {
			return fmt.Errorf("failed to delete old article politicians: %v", err)
		}
		if len(article.ArticlePoliticians) > 0 {
			if err := tx.Create(article.ArticlePoliticians).Error; err != nil {
				return fmt.Errorf("failed to create new article politicians: %v", err)
			}
		}

		return nil
	})
}
