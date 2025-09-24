package create_article

import (
	"time"
	"vdm/core/models"

	"github.com/google/uuid"
)

type RequestDTO struct {
	Title       string                 `json:"title" validate:"required"`
	Body        string                 `json:"body" validate:"required"`
	EventDate   time.Time              `json:"eventDate"`
	Tags        []string               `json:"tags" validate:"required,min=1,max=10"`
	Politicians []uuid.UUID            `json:"politicians" validate:"required,min=1,max=5"`
	Sources     []string               `json:"sources" validate:"required,min=1,max=5"`
	Category    models.ArticleCategory `json:"category"`
}
