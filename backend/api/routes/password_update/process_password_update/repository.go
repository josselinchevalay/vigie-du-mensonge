package process_password_update

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	updateUserPassword(userID uuid.UUID, password string) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) updateUserPassword(userID uuid.UUID, password string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
}
