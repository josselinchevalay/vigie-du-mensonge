package claim_moderator_article

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	updateArticleModerator(moderatorID, articleID uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) updateArticleModerator(moderatorID, articleID uuid.UUID) error {
	return r.db.Model(&models.Article{}).
		Where("id = ? AND redactor_id <> ? AND moderator_id IS NULL", articleID, moderatorID).
		Update("moderator_id", moderatorID).Error
}
