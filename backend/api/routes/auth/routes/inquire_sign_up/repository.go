package inquire_sign_up

import (
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	userExists(email string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) userExists(email string) (bool, error) {
	var exists bool

	if err := r.db.Model(&models.User{}).
		Where("email = ?", email).
		Select("count(*) > 0").
		Find(&exists).Error; err != nil {
		return false, err
	}

	return exists, nil
}
