package process_sign_up

import (
	"vdm/core/logger"
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	createUserAndRefreshToken(user *models.User, rft *models.UserToken) error
	userExistsByTag(tag string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) createUserAndRefreshToken(user *models.User, rft *models.UserToken) (err error) {
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

	var roleRedactor models.Role
	if err = tx.Where("name = ?", models.RoleRedactor).
		First(&roleRedactor).Error; err != nil {
		return
	}

	user.Roles = []*models.Role{&roleRedactor}

	if err = tx.Create(user).Error; err != nil {
		return
	}

	rft.UserID = user.ID
	if err = tx.Create(rft).Error; err != nil {
		return
	}

	return
}

func (r *repository) userExistsByTag(tag string) (bool, error) {
	var exists bool

	if err := r.db.Model(&models.User{}).Where("tag = ?", tag).
		Select("count(*)>0").
		Find(&exists).Error; err != nil {
		return false, err
	}

	return exists, nil
}
