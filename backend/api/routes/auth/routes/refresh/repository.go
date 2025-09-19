package refresh

import (
	"vdm/core/logger"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findValidRefreshToken(id uuid.UUID) (models.RefreshToken, error)
	createRefreshToken(rft *models.RefreshToken) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findValidRefreshToken(id uuid.UUID) (models.RefreshToken, error) {
	var rft models.RefreshToken
	if err := r.db.Model(&models.RefreshToken{}).
		Where("id = ? AND expiry > NOW()", id).
		Preload("User.Roles").
		First(&rft).Error; err != nil {
		return models.RefreshToken{}, err
	}
	return rft, nil
}

func (r *repository) createRefreshToken(rft *models.RefreshToken) (err error) {
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

	if err = tx.Model(&models.RefreshToken{}).
		Where("user_id = ?", rft.UserID).
		Delete(&models.RefreshToken{}).Error; err != nil {
		return
	}

	if err = tx.Create(rft).Error; err != nil {
		return
	}

	return
}
