package sign_in

import (
	"vdm/core/logger"
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	findUserByEmail(email string) (models.User, error)
	createRefreshToken(rft *models.RefreshToken) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).
		Where("email = ?", email).
		Preload("Roles").
		First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
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
