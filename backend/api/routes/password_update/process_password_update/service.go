package process_password_update

import (
	"fmt"
	"time"
	"vdm/core/jwt_utils"
	"vdm/core/locals"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	processPasswordUpdate(authedUser locals.AuthedUser, token, newPassword string) error
}

type service struct {
	passwordUpdateTokenSecret []byte
	passwordUpdateTokenTTL    time.Duration
	repo                      Repository
}

func (s *service) processPasswordUpdate(authedUser locals.AuthedUser, token, newPassword string) error {
	// Parse and validate the password update token
	if tokenUser, err := jwt_utils.ParseJWT(token, s.passwordUpdateTokenSecret); err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	} else if tokenUser.ID != authedUser.ID {
		return &fiber.Error{Code: fiber.StatusForbidden, Message: "credentials do not match"}
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Update the user's password
	if err := s.repo.updateUserPassword(authedUser.ID, string(hashedPassword)); err != nil {
		return fmt.Errorf("failed to update user password: %v", err)
	}
	return nil
}
