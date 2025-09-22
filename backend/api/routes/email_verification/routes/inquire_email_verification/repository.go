package inquire_email_verification

import (
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	createUserToken(usrTok *models.UserToken) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) createUserToken(usrTok *models.UserToken) error {
	return r.db.Create(usrTok).Error
}
