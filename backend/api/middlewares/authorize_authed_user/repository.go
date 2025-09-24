package authorize_authed_user

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	getUserRoles(userID uuid.UUID) ([]*models.Role, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) getUserRoles(userID uuid.UUID) ([]*models.Role, error) {
	var user models.User

	if err := r.db.Where("id = ?", userID).
		Select("id").
		Preload("Roles").
		First(&user).Error; err != nil {
		return nil, err
	}

	return user.Roles, nil
}
