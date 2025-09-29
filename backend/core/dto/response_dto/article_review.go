package response_dto

import "vdm/core/models"

type ArticleReview struct {
	ModeratorTag string               `json:"moderatorTag,omitempty"`
	Decision     models.ArticleStatus `json:"decision,omitempty"`
	Notes        string               `json:"notes,omitempty"`
}

func NewArticleReview(entity models.ArticleReview) ArticleReview {
	dto := ArticleReview{
		Decision: entity.Decision,
		Notes:    entity.Notes,
	}

	if entity.Moderator != nil {
		dto.ModeratorTag = entity.Moderator.Tag
	}

	return dto
}
