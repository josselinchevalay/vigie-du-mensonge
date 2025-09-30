package save_redactor_article

import (
	"fmt"
	"time"
	"vdm/core/models"

	"github.com/google/uuid"
)

type RequestDTO struct {
	ID            uuid.UUID              `json:"id"`
	Title         string                 `json:"title" validate:"required,min=20,max=200"`
	EventDate     time.Time              `json:"eventDate" validate:"required"`
	Category      models.ArticleCategory `json:"category" validate:"required"`
	Body          string                 `json:"body,omitempty"`
	Tags          []string               `json:"tags,omitempty" validate:"max=5"`
	PoliticianIDs []uuid.UUID            `json:"politicianIds,omitempty" validate:"max=5"`
	Sources       []string               `json:"sources,omitempty" validate:"max=5"`
}

func (dto RequestDTO) toArticle(redactorID uuid.UUID) (models.Article, error) {
	article := models.Article{
		ID:         dto.ID,
		RedactorID: redactorID,
		Title:      dto.Title,
		EventDate:  dto.EventDate,
		Body:       dto.Body,
		Category:   dto.Category,
	}

	if !article.Category.Valid() {
		return models.Article{}, fmt.Errorf("invalid category: %s", article.Category)
	}

	seenSources := make(map[string]bool)
	for _, source := range dto.Sources {
		if seenSources[source] {
			return models.Article{}, fmt.Errorf("duplicate source: %s", source)
		}
		seenSources[source] = true
		article.Sources = append(article.Sources, &models.ArticleSource{ArticleID: article.ID, URL: source})
	}

	seenTags := make(map[string]bool)
	for _, tag := range dto.Tags {
		if seenTags[tag] {
			return models.Article{}, fmt.Errorf("duplicate tag: %s", tag)
		}
		seenTags[tag] = true
		article.Tags = append(article.Tags, &models.ArticleTag{ArticleID: article.ID, Tag: tag})
	}

	seenPoliticians := make(map[uuid.UUID]bool)
	for _, polID := range dto.PoliticianIDs {
		if seenPoliticians[polID] {
			return models.Article{}, fmt.Errorf("duplicate politician: %s", polID)
		}
		seenPoliticians[polID] = true
		article.ArticlePoliticians = append(article.ArticlePoliticians, &models.ArticlePolitician{ArticleID: article.ID, PoliticianID: polID})
	}

	return article, nil
}
