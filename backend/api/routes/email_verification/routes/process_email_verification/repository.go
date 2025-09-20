package process_email_verification

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	updateUserEmailVerified(userID uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) updateUserEmailVerified(userID uuid.UUID) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("email_verified", true).Error
}
