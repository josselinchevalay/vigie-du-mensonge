package process_password_update

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

func newServiceWithMock(ctrl *gomock.Controller, repo *MockRepository) *service {
	return &service{
		passwordUpdateTokenSecret: []byte("secret"),
		passwordUpdateTokenTTL:    time.Minute,
		repo:                      repo,
	}
}

func TestService_ProcessPasswordUpdate_InvalidToken_ReturnsUnauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	// repo should not be called
	repo.EXPECT().updateUserPassword(gomock.Any(), gomock.Any()).Times(0)

	svc := newServiceWithMock(ctrl, repo)

	err := svc.processPasswordUpdate("this-is-not-a-valid-jwt", "NewPass123!")

	var ferr *fiber.Error
	if assert.Error(t, err) && assert.ErrorAs(t, err, &ferr) {
		assert.Equal(t, fiber.StatusUnauthorized, ferr.Code)
	}
}

func TestService_ProcessPasswordUpdate_RepositoryFailure_PropagatesError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	svc := newServiceWithMock(ctrl, repo)

	// user to embed in token
	user := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}
	// token for the user
	token, err := jwt_utils.GenerateJWT(user, svc.passwordUpdateTokenSecret, time.Now().Add(svc.passwordUpdateTokenTTL))
	if err != nil {
		t.Fatal(err)
	}

	// Simulate repo failure
	repoErr := errors.New("db update failed")
	repo.EXPECT().updateUserPassword(user.ID, gomock.Any()).Return(repoErr).Times(1)

	err = svc.processPasswordUpdate(token, "NewPass123!")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "failed to update user password")
	}
}
