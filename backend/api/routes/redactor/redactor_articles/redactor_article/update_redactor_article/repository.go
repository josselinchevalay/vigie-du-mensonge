package update_redactor_article

import (
	"fmt"
	"time"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	updateArticle(redactorID, articleID uuid.UUID, newData updateData) error
}

type repository struct {
	db *gorm.DB
}

type updateData struct {
	Title       string
	Body        string
	EventDate   time.Time
	Category    models.ArticleCategory
	Politicians []uuid.UUID
	Tags        []string
	Sources     []string
}

func (r *repository) updateArticle(redactorID, articleID uuid.UUID, newData updateData) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Load the article ensuring it belongs to the author; load current status to preserve
		var existing models.Article
		if err := tx.Where("id = ? AND redactor_id = ?", articleID, redactorID).
			Select("id", "redactor_id", "status").
			First(&existing).Error; err != nil {
			return err
		}

		// Update base fields except status
		if err := tx.Model(&models.Article{}).
			Where("id = ?", articleID).
			Updates(map[string]any{
				"title":      newData.Title,
				"body":       newData.Body,
				"event_date": newData.EventDate,
				"category":   newData.Category,
			}).Error; err != nil {
			return fmt.Errorf("failed to update article: %w", err)
		}

		// Replace politicians relations
		if err := tx.Where("article_id = ?", articleID).Delete(&models.ArticlePolitician{}).Error; err != nil {
			return fmt.Errorf("failed to clear politicians: %w", err)
		}
		if len(newData.Politicians) > 0 {
			rows := make([]models.ArticlePolitician, 0, len(newData.Politicians))
			seen := make(map[uuid.UUID]struct{}, len(newData.Politicians))
			for _, pid := range newData.Politicians {
				if _, ok := seen[pid]; ok {
					continue
				}
				seen[pid] = struct{}{}
				rows = append(rows, models.ArticlePolitician{ArticleID: articleID, PoliticianID: pid})
			}
			if len(rows) > 0 {
				if err := tx.Create(&rows).Error; err != nil {
					return fmt.Errorf("failed to insert politicians: %w", err)
				}
			}
		}

		// Replace tags (deduplicate)
		if err := tx.Where("article_id = ?", articleID).Delete(&models.ArticleTag{}).Error; err != nil {
			return fmt.Errorf("failed to clear tags: %w", err)
		}
		if len(newData.Tags) > 0 {
			seenTags := make(map[string]struct{}, len(newData.Tags))
			rows := make([]models.ArticleTag, 0, len(newData.Tags))
			for _, t := range newData.Tags {
				if t == "" {
					continue
				}
				if _, ok := seenTags[t]; ok {
					continue
				}
				seenTags[t] = struct{}{}
				rows = append(rows, models.ArticleTag{ArticleID: articleID, Tag: t})
			}
			if len(rows) > 0 {
				if err := tx.Create(&rows).Error; err != nil {
					return fmt.Errorf("failed to insert tags: %w", err)
				}
			}
		}

		// Replace sources (deduplicate)
		if err := tx.Where("article_id = ?", articleID).Delete(&models.ArticleSource{}).Error; err != nil {
			return fmt.Errorf("failed to clear sources: %w", err)
		}
		if len(newData.Sources) > 0 {
			seenSources := make(map[string]struct{}, len(newData.Sources))
			rows := make([]models.ArticleSource, 0, len(newData.Sources))
			for _, s := range newData.Sources {
				if s == "" {
					continue
				}
				if _, ok := seenSources[s]; ok {
					continue
				}
				seenSources[s] = struct{}{}
				rows = append(rows, models.ArticleSource{ArticleID: articleID, URL: s})
			}
			if len(rows) > 0 {
				if err := tx.Create(&rows).Error; err != nil {
					return fmt.Errorf("failed to insert sources: %w", err)
				}
			}
		}

		return nil
	})
}
