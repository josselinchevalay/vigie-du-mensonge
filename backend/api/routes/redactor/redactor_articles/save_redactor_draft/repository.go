package save_redactor_draft

import (
	"fmt"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	saveArticle(article *models.Article) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) saveArticle(article *models.Article) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if article.ID == uuid.Nil {
			return tx.Create(article).Error
		}

		var existing models.Article
		if err := tx.Where("id = ? AND redactor_id = ? AND status = ?", article.ID, article.RedactorID, models.ArticleStatusDraft).
			Select("id", "redactor_id", "status").
			First(&existing).Error; err != nil {
			return fmt.Errorf("failed to find article: %v", err)
		}

		// Update base fields except status
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
		if err := tx.Create(article.Tags).Error; err != nil {
			return fmt.Errorf("failed to create new tags: %v", err)
		}

		if err := tx.Where("article_id = ?", article.ID).
			Delete(&models.ArticleSource{}).Error; err != nil {
			return fmt.Errorf("failed to delete old sources: %v", err)
		}
		if err := tx.Create(article.Sources).Error; err != nil {
			return fmt.Errorf("failed to create new sources: %v", err)
		}

		if err := tx.Where("article_id = ?", article.ID).
			Delete(&models.ArticlePolitician{}).Error; err != nil {
			return fmt.Errorf("failed to delete old article politicians: %v", err)
		}
		if err := tx.Create(article.ArticlePoliticians).Error; err != nil {
			return fmt.Errorf("failed to create new article politicians: %v", err)
		}

		return nil
	})
}
