package inquire_password_update

import (
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	findUserID(email string) (uuid.UUID, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findUserID(email string) (uuid.UUID, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).
		Select("id").
		First(&user).Error; err != nil {
		return uuid.UUID{}, err
	}
	return user.ID, nil
}
