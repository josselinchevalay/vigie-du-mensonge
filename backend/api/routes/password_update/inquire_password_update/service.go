package inquire_password_update

import (
	"errors"
	"fmt"
	"time"
	"vdm/core/dependencies/mailer"
	"vdm/core/hmac_utils"
	"vdm/core/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	inquirePasswordUpdate(email string) error
}

type service struct {
	repo        Repository
	mailer      mailer.Mailer
	tokenSecret []byte
	tokenTTL    time.Duration
	clientURL   string
}

func (s *service) inquirePasswordUpdate(email string) error {
	userID, err := s.repo.findUserID(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil //do not let potentially malicious users know if an email exists
		}
		return fmt.Errorf("failed to find user: %v", err)
	}

	token := uuid.New()

	if err = s.repo.createUserToken(&models.UserToken{
		UserID:   userID,
		Expiry:   time.Now().Add(s.tokenTTL),
		Category: models.UserTokenCategoryPassword,
		Hash:     hmac_utils.HashUUID(token, s.tokenSecret),
	}); err != nil {
		return fmt.Errorf("failed to create token: %v", err)
	}

	if err = s.mailer.Send(
		email,
		"Modification de votre mot de passe",
		fmt.Sprintf("Cliquez sur le lien ci-dessous pour modifier le mot de passe de votre compte Vigie du mensonge:\n\n%s?token=%s",
			s.clientURL+"/password-update", token),
	); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
