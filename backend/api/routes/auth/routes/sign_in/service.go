package sign_in

import (
	"fmt"
	"time"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	signIn(req SignInRequest) (models.User, locals.AccessToken, locals.RefreshToken, error)
}

type service struct {
	accessTokenTTL    time.Duration
	refreshTokenTTL   time.Duration
	accessTokenSecret []byte
	repo              Repository
}

func (s *service) signIn(req SignInRequest) (models.User, locals.AccessToken, locals.RefreshToken, error) {
	user, err := s.repo.findUserByEmail(req.Email)
	if err != nil {
		return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, &fiber.Error{Code: fiber.StatusNotFound, Message: "invalid credentials"}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, &fiber.Error{Code: fiber.StatusUnauthorized, Message: "invalid credentials"}
	}

	// create new refresh token for user
	rft := models.RefreshToken{UserID: user.ID, Expiry: time.Now().Add(s.refreshTokenTTL)}
	if err = s.repo.createRefreshToken(&rft); err != nil {
		return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, fmt.Errorf("failed to create refresh token: %v", err)
	}

	jwtExpiry := time.Now().Add(s.accessTokenTTL)
	jwt, err := jwt_utils.GenerateJWT(locals.NewAuthedUser(user), s.accessTokenSecret, jwtExpiry)
	if err != nil {
		return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, fmt.Errorf("failed to generate JWT: %v", err)
	}

	return user,
		locals.AccessToken{Token: jwt, Expiry: jwtExpiry},
		locals.RefreshToken{Token: rft.ID, Expiry: rft.Expiry},
		nil
}
