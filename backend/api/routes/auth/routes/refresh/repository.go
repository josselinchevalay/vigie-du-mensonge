package refresh

import (
	"vdm/core/logger"
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	findValidRefreshToken(hash string) (models.UserToken, error)
	createRefreshToken(rft *models.UserToken) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findValidRefreshToken(hash string) (models.UserToken, error) {
	var rft models.UserToken
	if err := r.db.Model(&models.UserToken{}).
		Where("hash = ? AND category = ? AND expiry > NOW()", hash, models.UserTokenCategoryRefresh).
		Preload("User.Roles").
		First(&rft).Error; err != nil {
		return models.UserToken{}, err
	}
	return rft, nil
}

func (r *repository) createRefreshToken(rft *models.UserToken) (err error) {
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

	if err = tx.Model(&models.UserToken{}).
		Where("user_id = ? AND category = ?", rft.UserID, models.UserTokenCategoryRefresh).
		Delete(&models.UserToken{}).Error; err != nil {
		return
	}

	if err = tx.Create(rft).Error; err != nil {
		return
	}

	return
}
