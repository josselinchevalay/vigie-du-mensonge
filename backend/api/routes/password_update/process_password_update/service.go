package process_password_update

import (
	"fmt"
	"vdm/core/hmac_utils"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	processPasswordUpdate(token uuid.UUID, newPassword string) error
}

type service struct {
	tokenSecret []byte
	repo        Repository
}

func (s *service) processPasswordUpdate(token uuid.UUID, newPassword string) error {
	var user models.User

	if usrTok, err := s.repo.findUserToken(hmac_utils.HashUUID(token, s.tokenSecret)); err != nil {
		return fmt.Errorf("failed to find token: %v", err)
	} else if usrTok.Expired() {
		return &fiber.Error{Code: fiber.StatusGone, Message: "token expired"}
	} else {
		user = *usrTok.User
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Update the user's password
	if err := s.repo.updateUserPasswordAndDeleteTokens(user.ID, string(hashedPassword)); err != nil {
		return fmt.Errorf("failed to update user password: %v", err)
	}

	return nil
}
