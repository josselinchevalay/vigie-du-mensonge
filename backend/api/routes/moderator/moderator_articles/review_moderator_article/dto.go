package review_moderator_article

import "vdm/core/models"

type RequestDTO struct {
	Decision models.ArticleStatus `json:"decision" validate:"required"`
	Notes    string               `json:"notes,omitempty" validate:"max=200"`
}
