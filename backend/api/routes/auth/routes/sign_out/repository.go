package sign_out

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	deleteRefreshTokens(userID uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func (r repository) deleteRefreshTokens(userID uuid.UUID) error {
	return r.db.Where("user_id = ? AND category = ?", userID, models.UserTokenCategoryRefresh).
		Delete(&models.UserToken{}).Error
}
