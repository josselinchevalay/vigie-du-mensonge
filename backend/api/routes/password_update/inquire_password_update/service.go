package inquire_password_update

import (
	"errors"
	"fmt"
	"time"
	"vdm/core/dependencies/mailer"
	"vdm/core/jwt_utils"
	"vdm/core/locals"

	"gorm.io/gorm"
)

type Service interface {
	inquirePasswordUpdate(email string) error
}

type service struct {
	repo                      Repository
	mailer                    mailer.Mailer
	passwordUpdateTokenSecret []byte
	passwordUpdateTokenTTL    time.Duration
	clientURL                 string
}

func (s *service) inquirePasswordUpdate(email string) error {
	userID, err := s.repo.findUserID(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return fmt.Errorf("failed to find user: %v", err)
	}

	token, err := jwt_utils.GenerateJWT(
		locals.AuthedUser{ID: userID, Email: email},
		s.passwordUpdateTokenSecret,
		time.Now().Add(s.passwordUpdateTokenTTL),
	)
	if err != nil {
		return fmt.Errorf("failed to generate password update token: %v", err)
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
