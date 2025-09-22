package process_email_verification

import (
	"fmt"
	"vdm/core/hmac_utils"
	"vdm/core/locals"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	validateTokenAndVerifyEmail(authedUser locals.AuthedUser, token uuid.UUID) error
}

type service struct {
	tokenSecret []byte
	repo        Repository
}

func (s *service) validateTokenAndVerifyEmail(authedUser locals.AuthedUser, token uuid.UUID) error {
	if usrTok, err := s.repo.findEmailToken(authedUser.ID, hmac_utils.HashUUID(token, s.tokenSecret)); err != nil {
		return fmt.Errorf("failed to find token: %v", err)
	} else if usrTok.Expired() {
		return &fiber.Error{Code: fiber.StatusGone, Message: "token expired"}
	}

	if err := s.repo.updateUserEmailVerifiedAndDeleteTokens(authedUser.ID); err != nil {
		return fmt.Errorf("failed to update user verification status: %v", err)
	}
	return nil
}
