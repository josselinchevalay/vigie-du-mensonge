package process_password_update

import (
	"vdm/core/logger"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findUserToken(hash string) (models.UserToken, error)
	updateUserPasswordAndDeleteTokens(userID uuid.UUID, password string) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findUserToken(hash string) (models.UserToken, error) {
	var usrTok models.UserToken

	if err := r.db.Where("hash = ? AND category = ?", hash, models.UserTokenCategoryPassword).
		Preload("User").
		First(&usrTok).Error; err != nil {
		return models.UserToken{}, err
	}

	return usrTok, nil
}

func (r *repository) updateUserPasswordAndDeleteTokens(userID uuid.UUID, password string) (err error) {
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

	if err = tx.Error; err != nil {
		return
	}

	if err = tx.Model(&models.User{}).Where("id = ?", userID).
		Update("password", password).Error; err != nil {
		return
	}

	err = tx.Model(&models.UserToken{}).
		Where("user_id = ? AND category = ?", userID, models.UserTokenCategoryPassword).
		Delete(&models.UserToken{}).Error

	return
}
