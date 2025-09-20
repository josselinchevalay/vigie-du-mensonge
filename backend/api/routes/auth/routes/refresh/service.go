package refresh

import (
	"fmt"
	"time"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	refresh(rftID uuid.UUID) (models.User, locals.AccessToken, locals.RefreshToken, error)
}

type service struct {
	accessTokenTTL    time.Duration
	refreshTokenTTL   time.Duration
	accessTokenSecret []byte
	repo              Repository
}

func (s *service) refresh(rftID uuid.UUID) (models.User, locals.AccessToken, locals.RefreshToken, error) {
	rft, err := s.repo.findValidRefreshToken(rftID)
	if err != nil {
		return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, &fiber.Error{Code: fiber.StatusUnauthorized, Message: "invalid refresh token"}
	}

	user := rft.User
	if user == nil {
		return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, &fiber.Error{Code: fiber.StatusInternalServerError, Message: "unexpected nil user"}
	}

	// issue a new refresh token for the same user with a renewed expiry
	rft = models.RefreshToken{
		UserID: rft.UserID,
		Expiry: time.Now().Add(s.refreshTokenTTL),
	}

	if err = s.repo.createRefreshToken(&rft); err != nil {
		return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, fmt.Errorf("failed to create refresh token: %v", err)
	}

	jwtExpiry := time.Now().Add(s.accessTokenTTL)
	jwt, err := jwt_utils.GenerateJWT(locals.NewAuthedUser(*user), s.accessTokenSecret, jwtExpiry)
	if err != nil {
		return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, fmt.Errorf("failed to generate JWT: %v", err)
	}

	return *user,
		locals.AccessToken{
			Token:  jwt,
			Expiry: jwtExpiry,
		}, locals.RefreshToken{
			Token:  rft.ID,
			Expiry: rft.Expiry,
		}, nil
}
