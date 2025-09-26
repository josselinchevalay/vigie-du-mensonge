package find_redactor_article

import (
	"time"
	"vdm/core/models"

	"github.com/google/uuid"
)

type ArticleDTO struct {
	ID          uuid.UUID              `json:"id"`
	Title       string                 `json:"title"`
	Body        string                 `json:"body"`
	Sources     []string               `json:"sources"`
	EventDate   time.Time              `json:"eventDate"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	Politicians []PoliticianDTO        `json:"politicians"`
	Tags        []string               `json:"tags"`
	Category    models.ArticleCategory `json:"category"`
}

func newArticleDTO(entity models.Article) ArticleDTO {
	dto := ArticleDTO{
		ID:          entity.ID,
		Title:       entity.Title,
		Body:        entity.Body,
		Sources:     make([]string, len(entity.Sources)),
		EventDate:   entity.EventDate,
		UpdatedAt:   entity.UpdatedAt,
		Politicians: make([]PoliticianDTO, len(entity.Politicians)),
		Tags:        make([]string, len(entity.Tags)),
		Category:    entity.Category,
	}

	for i := range entity.Sources {
		dto.Sources[i] = entity.Sources[i].URL
	}
	for i := range entity.Politicians {
		dto.Politicians[i] = newPoliticianDTO(*entity.Politicians[i])
	}
	for i := range entity.Tags {
		dto.Tags[i] = entity.Tags[i].Tag
	}

	return dto
}

type PoliticianDTO struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
}

func newPoliticianDTO(entity models.Politician) PoliticianDTO {
	return PoliticianDTO{
		ID:       entity.ID,
		FullName: entity.FirstName + " " + entity.LastName,
	}
}
