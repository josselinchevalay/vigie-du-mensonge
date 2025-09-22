package process_email_verification

import (
	"vdm/core/logger"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findEmailToken(userID uuid.UUID, hash string) (models.UserToken, error)
	updateUserEmailVerifiedAndDeleteTokens(userID uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findEmailToken(userID uuid.UUID, hash string) (models.UserToken, error) {
	var usrTok models.UserToken
	if err := r.db.Where("user_id = ? AND hash = ? AND category = ?", userID, hash, models.UserTokenCategoryEmail).
		First(&usrTok).Error; err != nil {
		return models.UserToken{}, err
	}
	return usrTok, nil
}

func (r *repository) updateUserEmailVerifiedAndDeleteTokens(userID uuid.UUID) (err error) {
	tx := r.db.Begin()

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback().Error; rbErr != nil {
				logger.Error("failed to rollback transaction", logger.Err(rbErr))
			}
			return
		}

		if cmErr := tx.Commit().Error; cmErr != nil {
			logger.Error("failed to commit transaction", logger.Err(cmErr))
			err = cmErr
		}
	}()

	if err = tx.Model(&models.User{}).Where("id = ?", userID).
		Update("email_verified", true).Error; err != nil {
		return
	}

	err = tx.Model(&models.UserToken{}).
		Where("user_id = ? AND category = ?", userID, models.UserTokenCategoryEmail).
		Delete(&models.UserToken{}).Error

	return
}
