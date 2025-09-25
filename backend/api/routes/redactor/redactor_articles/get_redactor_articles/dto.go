package get_redactor_articles

import (
	"time"
	"vdm/core/models"

	"github.com/google/uuid"
)

type ResponseDTO []ArticleDTO

type ArticleDTO struct {
	ID          uuid.UUID              `json:"id"`
	Title       string                 `json:"title"`
	EventDate   time.Time              `json:"eventDate"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	Politicians []PoliticianDTO        `json:"politicians"`
	Tags        []string               `json:"tags"`
	Category    models.ArticleCategory `json:"category"`
	Status      models.ArticleStatus   `json:"status"`
}

func newArticleDTO(entity models.Article) ArticleDTO {
	dto := ArticleDTO{
		ID:          entity.ID,
		Title:       entity.Title,
		EventDate:   entity.EventDate,
		UpdatedAt:   entity.UpdatedAt,
		Politicians: make([]PoliticianDTO, len(entity.Politicians)),
		Tags:        make([]string, len(entity.Tags)),
		Status:      entity.Status,
		Category:    entity.Category,
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
