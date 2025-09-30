package moderator_save_review

import "vdm/core/models"

type RequestDTO struct {
	Decision models.ArticleStatus `json:"decision" validate:"required"`
	Notes    string               `json:"notes,omitempty" validate:"max=200"`
}
