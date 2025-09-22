package sign_up

import (
	"errors"
	"fmt"
	"time"
	"vdm/core/hmac_utils"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	signUp(req RequestDTO) (locals.AccessToken, locals.RefreshToken, error)
}

type service struct {
	accessTokenTTL     time.Duration
	refreshTokenTTL    time.Duration
	accessTokenSecret  []byte
	refreshTokenSecret []byte
	repo               Repository
}

func (s *service) signUp(req RequestDTO) (locals.AccessToken, locals.RefreshToken, error) {
	user := &models.User{Email: req.Email}

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost); err != nil {
		return locals.AccessToken{}, locals.RefreshToken{}, fmt.Errorf("failed to hash password: %v", err)
	} else {
		user.Password = string(hashedPassword)
	}

	rft := uuid.New()
	usrTok := &models.UserToken{UserID: user.ID, Hash: hmac_utils.HashUUID(rft, s.refreshTokenSecret),
		Expiry: time.Now().Add(s.refreshTokenTTL), Category: models.UserTokenCategoryRefresh}

	if err := s.repo.createUserAndRefreshToken(user, usrTok); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return locals.AccessToken{}, locals.RefreshToken{}, &fiber.Error{Code: fiber.StatusConflict, Message: "user with this email already exists"}
		}
		return locals.AccessToken{}, locals.RefreshToken{}, fmt.Errorf("failed to create user and refresh token: %v", err)
	}

	jwtExpiry := time.Now().Add(s.accessTokenTTL)
	jwt, err := jwt_utils.GenerateJWT(locals.AuthedUser{ID: user.ID, Email: user.Email}, s.accessTokenSecret, jwtExpiry)
	if err != nil {
		return locals.AccessToken{}, locals.RefreshToken{}, fmt.Errorf("failed to generate JWT: %v", err)
	}

	return locals.AccessToken{
			Token:  jwt,
			Expiry: jwtExpiry,
		}, locals.RefreshToken{
			Token:  rft,
			Expiry: usrTok.Expiry,
		}, nil
}
