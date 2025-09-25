package create_redactor_article

import (
	"fmt"
	"vdm/core/models"

	"github.com/google/uuid"
)

type Service interface {
	mapAndCreateArticle(authorID uuid.UUID, dto RequestDTO) (uuid.UUID, error)
}

type service struct {
	repo Repository
}

func (s *service) mapAndCreateArticle(authorID uuid.UUID, dto RequestDTO) (uuid.UUID, error) {
	article := models.Article{
		AuthorID:    authorID,
		Title:       dto.Title,
		Body:        dto.Body,
		Category:    dto.Category,
		Status:      models.ArticleStatusDraft,
		EventDate:   dto.EventDate,
		Major:       0,
		Minor:       0,
		Reference:   uuid.New(),
		Politicians: make([]*models.Politician, len(dto.Politicians)),
		Tags:        make([]*models.ArticleTag, len(dto.Tags)),
		Sources:     make([]*models.ArticleSource, len(dto.Sources)),
	}

	for i := range dto.Politicians {
		article.Politicians[i] = &models.Politician{ID: dto.Politicians[i]}
	}
	for i := range dto.Tags {
		article.Tags[i] = &models.ArticleTag{Tag: dto.Tags[i]}
	}
	for i := range dto.Sources {
		article.Sources[i] = &models.ArticleSource{URL: dto.Sources[i]}
	}

	if err := s.repo.createArticle(&article); err != nil {
		return uuid.Nil, fmt.Errorf("failed to create article: %w", err)
	}

	return article.ID, nil
}
