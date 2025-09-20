package process_email_verification

import (
	"errors"
	"testing"
	"time"
	"vdm/core/jwt_utils"
	"vdm/core/locals"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestService_ProcessEmailVerification_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)

	secret := []byte("secret")
	svc := &service{emailVerificationTokenSecret: secret, emailVerificationTokenTTL: time.Minute, repo: repo}

	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}
	// Create a token for the same authed user
	expiry := time.Now().Add(time.Minute)
	token, err := jwt_utils.GenerateJWT(authed, secret, expiry)
	if err != nil {
		t.Fatal(err)
	}

	repo.EXPECT().updateUserEmailVerified(authed.ID).Return(nil).Times(1)

	err = svc.processEmailVerification(authed, token)
	assert.NoError(t, err)
}

func TestService_ProcessEmailVerification_Unauthorized_InvalidToken(t *testing.T) {
	repo := NewMockRepository(gomock.NewController(t))
	defer repo.ctrl.Finish()

	secret := []byte("secret")
	svc := &service{emailVerificationTokenSecret: secret, emailVerificationTokenTTL: time.Minute, repo: repo}

	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}

	// Use an obviously invalid token string to force ParseJWT error
	token := "not-a-valid-jwt"

	repo.EXPECT().updateUserEmailVerified(gomock.Any()).Times(0)

	err := svc.processEmailVerification(authed, token)
	var fe *fiber.Error
	if assert.Error(t, err) && assert.ErrorAs(t, err, &fe) {
		assert.Equal(t, fiber.StatusUnauthorized, fe.Code)
	}
}

func TestService_ProcessEmailVerification_Forbidden_UserMismatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)

	secret := []byte("secret")
	svc := &service{emailVerificationTokenSecret: secret, emailVerificationTokenTTL: time.Minute, repo: repo}

	// Authenticated user
	authed := locals.AuthedUser{ID: uuid.New(), Email: "real@example.com"}
	// Token generated for a different user
	tokenUser := locals.AuthedUser{ID: uuid.New(), Email: "other@example.com"}
	token, err := jwt_utils.GenerateJWT(tokenUser, secret, time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	repo.EXPECT().updateUserEmailVerified(gomock.Any()).Times(0)

	err = svc.processEmailVerification(authed, token)
	var fe *fiber.Error
	if assert.Error(t, err) && assert.ErrorAs(t, err, &fe) {
		assert.Equal(t, fiber.StatusForbidden, fe.Code)
	}
}

func TestService_ProcessEmailVerification_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)

	secret := []byte("secret")
	svc := &service{emailVerificationTokenSecret: secret, emailVerificationTokenTTL: time.Minute, repo: repo}

	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}
	token, err := jwt_utils.GenerateJWT(authed, secret, time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	repo.EXPECT().updateUserEmailVerified(authed.ID).Return(errors.New("db failure")).Times(1)

	err = svc.processEmailVerification(authed, token)
	// Expect a generic error (not a fiber.Error) wrapping repo error
	assert.Error(t, err)
	var fe *fiber.Error
	assert.False(t, errors.As(err, &fe))
	assert.Contains(t, err.Error(), "failed to update user verification status")
}
