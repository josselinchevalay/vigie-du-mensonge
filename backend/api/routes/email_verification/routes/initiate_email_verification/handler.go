package initiate_email_verification

import (
	"errors"
	"fmt"
	"time"
	"vdm/core/dependencies/mailer"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	inquireEmailVerification(c *fiber.Ctx) error
}

type handler struct {
	emailVerificationTokenSecret []byte
	emailVerificationTokenTTL    time.Duration
	fullPath                     string
	mailer                       mailer.Mailer
}

func (h *handler) inquireEmailVerification(c *fiber.Ctx) error {
	// Retrieve authenticated user from locals (set by locals_authed_user middleware)
	authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
	if !ok {
		return errors.New("failed to retrieve authenticated user from locals")
	}

	// Generate a short-lived email verification token
	expiry := time.Now().Add(h.emailVerificationTokenTTL)
	token, err := jwt_utils.GenerateJWT(authedUser, h.emailVerificationTokenSecret, expiry)
	if err != nil {
		return fmt.Errorf("failed to generate email verification token: %v", err)
	}

	// Send the email containing the verification token
	subject := "Vérification de votre adresse e-mail"

	body := fmt.Sprintf(
		"Cliquez sur le lien ci-dessous pour vérifier l'adresse email de votre compte Vigie du mensonge:\n\n%s?token=%s",
		h.fullPath, token,
	)

	if err := h.mailer.Send(authedUser.Email, subject, body); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
