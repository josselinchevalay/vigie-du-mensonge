package sign_up

import (
	"errors"
	"testing"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestService_ErrConflict(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	repo := NewMockRepository(mockCtrl)
	s := &service{repo: repo}

	req := SignUpRequest{
		Email:    "hello@world.com",
		Password: "HelloWorld123!",
	}

	repo.EXPECT().createUserAndRefreshToken(gomock.Any(), gomock.Any()).
		DoAndReturn(func(user *models.User, rft *models.RefreshToken) error {
			assert.Equal(t, user.Email, req.Email)
			assert.Equal(t, user.ID, rft.UserID)
			return gorm.ErrDuplicatedKey
		})

	_, _, err := s.signUp(req)

	var fiberErr *fiber.Error
	if ok := errors.As(err, &fiberErr); !ok {
		t.Fatal("expected fiber error")
	}

	assert.Equal(t, fiberErr.Code, fiber.StatusConflict)
}
