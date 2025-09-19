package sign_out

import (
	"testing"
	"time"
	"vdm/core/jwt_utils"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestService_SignOut_ValidToken_DeletesRefreshTokens(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	repo := NewMockRepository(mockCtrl)

	s := &service{repo: repo, accessTokenSecret: []byte("dummy")}

	userID := uuid.New()
	jwt, err := jwt_utils.GenerateJWT(userID, s.accessTokenSecret, time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	repo.EXPECT().deleteRefreshTokens(userID)

	s.signOut(jwt)
}

func TestService_SignOut_InvalidToken_NoDeletion(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	repo := NewMockRepository(mockCtrl)

	s := &service{repo: repo, accessTokenSecret: []byte("dummy")}

	// invalid token string should not trigger repo call
	s.signOut("not-a-jwt")
}
