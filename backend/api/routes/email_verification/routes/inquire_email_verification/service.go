package inquire_email_verification

import (
	"fmt"
	"time"
	"vdm/core/dependencies/mailer"
	"vdm/core/hmac_utils"
	"vdm/core/locals"
	"vdm/core/models"

	"github.com/google/uuid"
)

type Service interface {
	sendEmailAndCreateToken(authedUser locals.AuthedUser) error
}

type service struct {
	repo        Repository
	mailer      mailer.Mailer
	clientURL   string
	tokenTTL    time.Duration
	tokenSecret []byte
}

func (s *service) sendEmailAndCreateToken(authedUser locals.AuthedUser) error {
	rft := uuid.New()

	if err := s.repo.createUserToken(&models.UserToken{
		UserID:   authedUser.ID,
		Hash:     hmac_utils.HashUUID(rft, s.tokenSecret),
		Expiry:   time.Now().Add(s.tokenTTL),
		Category: models.UserTokenCategoryEmail,
	}); err != nil {
		return err
	}

	subject := "Vérification de votre adresse e-mail"

	body := fmt.Sprintf(
		"Cliquez sur le lien ci-dessous pour vérifier l'adresse e-mail de votre compte Vigie du mensonge:\n\n%s?token=%s",
		s.clientURL+"/email-verification", rft,
	)

	if err := s.mailer.Send(authedUser.Email, subject, body); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
