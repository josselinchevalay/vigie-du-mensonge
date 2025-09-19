package refresh

import (
	"errors"
	"testing"
	"vdm/core/models"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestService_ErrCreateRefreshToken(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	repo := NewMockRepository(mockCtrl)
	s := &service{repo: repo}

	rftID := uuid.New()
	userID := uuid.New()

	repo.EXPECT().findValidRefreshToken(rftID).
		DoAndReturn(func(_ uuid.UUID) (models.RefreshToken, error) {
			return models.RefreshToken{UserID: userID, User: &models.User{}}, nil
		})

	repo.EXPECT().createRefreshToken(gomock.Any()).
		DoAndReturn(func(rft *models.RefreshToken) error {
			assert.Equal(t, userID, rft.UserID)
			return errors.New("db error")
		})

	_, _, _, err := s.refresh(rftID)

	if assert.Error(t, err) {
		assert.ErrorContains(t, err, "failed to create refresh token")
	}
}
