package response_dto

import (
	"time"
	"vdm/core/models"

	"github.com/google/uuid"
)

type Article struct {
	ID        uuid.UUID `json:"id"`
	Reference uuid.UUID `json:"reference"`

	RedactorTag  string `json:"redactorTag,omitempty"`
	ModeratorTag string `json:"moderatorTag,omitempty"`

	Title string `json:"title"`
	Body  string `json:"body,omitempty"`

	Category models.ArticleCategory `json:"category"`
	Status   models.ArticleStatus   `json:"status,omitempty"`

	EventDate time.Time `json:"eventDate"`
	UpdatedAt time.Time `json:"updatedAt"`

	Minor int16 `json:"minor,omitempty"`
	Major int16 `json:"major,omitempty"`

	Review *ArticleReview `json:"review,omitempty"`

	Sources     []string     `json:"sources,omitempty"`
	Politicians []Politician `json:"politicians,omitempty"`
	Tags        []string     `json:"tags,omitempty"`
}

func NewArticle(entity models.Article) Article {
	dto := Article{
		ID:        entity.ID,
		Reference: entity.Reference,
		Title:     entity.Title,
		Body:      entity.Body,
		Category:  entity.Category,
		Status:    entity.Status,
		EventDate: entity.EventDate,
		UpdatedAt: entity.UpdatedAt,
		Minor:     entity.Minor,
		Major:     entity.Major,
	}

	if entity.Redactor != nil {
		dto.RedactorTag = entity.Redactor.Tag
	}

	if entity.Moderator != nil {
		dto.ModeratorTag = entity.Moderator.Tag
	}

	if entity.Review != nil {
		reviewDTO := NewArticleReview(*entity.Review)
		dto.Review = &reviewDTO
	}

	if len(entity.Sources) > 0 {
		dto.Sources = make([]string, len(entity.Sources))
		for i := range entity.Sources {
			dto.Sources[i] = entity.Sources[i].URL
		}
	}

	if len(entity.Politicians) > 0 {
		dto.Politicians = make([]Politician, len(entity.Politicians))
		for i := range entity.Politicians {
			dto.Politicians[i] = NewPolitician(*entity.Politicians[i])
		}
	}

	if len(entity.Tags) > 0 {
		dto.Tags = make([]string, len(entity.Tags))
		for i := range entity.Tags {
			dto.Tags[i] = entity.Tags[i].Tag
		}
	}

	return dto
}
