package update_redactor_article

import (
	"fmt"
	"time"
	"vdm/core/models"
	"vdm/core/validation"

	"github.com/google/uuid"
)

//TODO: add validation to prevent duplicate on politicians/sources/tags

type RequestDTO struct {
	Title       string                 `json:"title" validate:"required,min=20,max=50"`
	Body        string                 `json:"body" validate:"required,min=200,max=2000"`
	EventDate   time.Time              `json:"eventDate"`
	Tags        []string               `json:"tags" validate:"required,min=1,max=10"`
	Politicians []uuid.UUID            `json:"politicians" validate:"required,min=1,max=5"`
	Sources     []string               `json:"sources" validate:"required,min=1,max=5"`
	Category    models.ArticleCategory `json:"category" validate:"required"`
}

func (dto RequestDTO) Validate() error {
	if err := validation.Validate(dto); err != nil {
		return err
	}

	if !dto.Category.Valid() {
		return fmt.Errorf("invalid category")
	}

	return nil
}
