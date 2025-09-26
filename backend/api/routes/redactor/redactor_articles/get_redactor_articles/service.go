package get_redactor_articles

import (
	"fmt"

	"github.com/google/uuid"
)

type Service interface {
	getAndMapRedactorArticles(redactorID uuid.UUID) (ResponseDTO, error)
}

type service struct {
	repo Repository
}

func (s *service) getAndMapRedactorArticles(redactorID uuid.UUID) (ResponseDTO, error) {
	articles, err := s.repo.getArticleByRedactorID(redactorID)
	if err != nil {
		return ResponseDTO{}, fmt.Errorf("failed to get articles: %v", err)
	}

	respDTO := make(ResponseDTO, len(articles))
	for i := range articles {
		respDTO[i] = newArticleDTO(articles[i])
	}
	return respDTO, nil
}
