package update_redactor_article

import (
	"fmt"

	"github.com/google/uuid"
)

type Service interface {
	updateRedactorArticle(authorID, articleID uuid.UUID, reqDTO RequestDTO) error
}

type service struct {
	repo Repository
}

func (s *service) updateRedactorArticle(authorID, articleID uuid.UUID, reqDTO RequestDTO) error {
	if err := reqDTO.Validate(); err != nil {
		return err
	}

	data := updateData{
		Title:       reqDTO.Title,
		Body:        reqDTO.Body,
		EventDate:   reqDTO.EventDate,
		Category:    reqDTO.Category,
		Politicians: reqDTO.Politicians,
		Tags:        reqDTO.Tags,
		Sources:     reqDTO.Sources,
	}

	if err := s.repo.updateArticle(authorID, articleID, data); err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}
	return nil
}
