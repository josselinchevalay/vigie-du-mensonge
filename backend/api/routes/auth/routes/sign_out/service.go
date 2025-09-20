package sign_out

import (
	"vdm/core/jwt_utils"
	"vdm/core/logger"
)

type Service interface {
	signOut(accessToken string)
}

type service struct {
	repo              Repository
	accessTokenSecret []byte
}

func (s *service) signOut(accessToken string) {
	authedUser, err := jwt_utils.ParseJWT(accessToken, s.accessTokenSecret)
	if err != nil {
		logger.Error("Failed to parse access token", logger.Err(err))
		return
	}

	if err = s.repo.deleteRefreshTokens(authedUser.ID); err != nil {
		logger.Error("Failed to delete refresh tokens", logger.Err(err))
	}
}
