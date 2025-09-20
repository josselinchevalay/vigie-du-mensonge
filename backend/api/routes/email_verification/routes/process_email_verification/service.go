package process_email_verification

import (
	"fmt"
	"time"
	"vdm/core/jwt_utils"
	"vdm/core/locals"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	processEmailVerification(authedUser locals.AuthedUser, token string) error
}

type service struct {
	emailVerificationTokenSecret []byte
	emailVerificationTokenTTL    time.Duration
	repo                         Repository
}

func (s *service) processEmailVerification(authedUser locals.AuthedUser, token string) error {
	// Parse and validate the verification token
	if emailUser, err := jwt_utils.ParseJWT(token, s.emailVerificationTokenSecret); err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	} else if emailUser.ID != authedUser.ID {
		return &fiber.Error{Code: fiber.StatusForbidden, Message: "credentials do not match"}
	}
	// Mark the user's email as verified
	if err := s.repo.updateUserEmailVerified(authedUser.ID); err != nil {
		return fmt.Errorf("failed to update user verification status: %v", err)
	}
	return nil
}
