package inquire_password_update

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
	inquirePasswordUpdate(c *fiber.Ctx) error
}

type handler struct {
	passwordUpdateTokenSecret []byte
	passwordUpdateTokenTTL    time.Duration
	clientURL                 string
	mailer                    mailer.Mailer
}

func (h *handler) inquirePasswordUpdate(c *fiber.Ctx) error {
	// Retrieve authenticated user from locals (set by locals_authed_user middleware)
	authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
	if !ok {
		return errors.New("failed to retrieve authenticated user from locals")
	}

	// Generate a short-lived password update token
	expiry := time.Now().Add(h.passwordUpdateTokenTTL)
	token, err := jwt_utils.GenerateJWT(authedUser, h.passwordUpdateTokenSecret, expiry)
	if err != nil {
		return fmt.Errorf("failed to generate password update token: %v", err)
	}

	// Send the email containing the password update link
	subject := "Mise à jour de votre mot de passe"
	body := fmt.Sprintf(
		"Cliquez sur le lien ci-dessous pour mettre à jour le mot de passe de votre compte Vigie du mensonge:\n\n%s?token=%s",
		h.clientURL+"/password-update", token,
	)

	if err := h.mailer.Send(authedUser.Email, subject, body); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
