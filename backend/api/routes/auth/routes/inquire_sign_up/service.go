package inquire_sign_up

import (
	"fmt"
	"time"
	"vdm/core/dependencies/mailer"
	"vdm/core/jwt_utils"
	"vdm/core/locals"

	"github.com/google/uuid"
)

type Service interface {
	checkUserAndSendEmail(email string) error
}

type service struct {
	repo             Repository
	mailer           mailer.Mailer
	emailTokenTTL    time.Duration
	emailTokenSecret []byte
	clientURL        string
}

func (s *service) checkUserAndSendEmail(email string) error {
	exists, err := s.repo.userExists(email)
	if err != nil {
		return err
	}

	var emailSubject, emailBody string
	if exists {
		emailSubject = "Votre compte Vigie du mensonge"
		emailBody = "Nous avez reçu une requête de création de compte pour cette adresse email. Cependant, vous possédez déjà un compte Vigie du mensonge.\n" +
			"Si vous avez perdu votre mot de passe, vous pouvez faire une demande de modification à partir de la page de connexion (lien ci-dessous)." +
			"\n\n" + s.clientURL + "/sign-in"
	} else {
		jwt, err := jwt_utils.GenerateJWT(
			locals.AuthedUser{ID: uuid.Nil, Email: email}, // explicity set ID to nil to highlight this user has not signed-up yet
			s.emailTokenSecret,
			time.Now().Add(s.emailTokenTTL),
		)
		if err != nil {
			return err
		}

		emailSubject = "Inscription sur Vigie du mensonge"
		emailBody = "Afin de créer votre compte Vigie du mensonge, veuillez cliquez sur le lien ci-dessous. Celui-ci expirera après 15 minutes.\n\n" +
			s.clientURL + "/sign-up?token=" + jwt
	}

	if err = s.mailer.Send(
		email,
		emailSubject,
		emailBody,
	); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
