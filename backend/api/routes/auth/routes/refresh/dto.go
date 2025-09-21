package refresh

import "time"

type ResponseDTO struct {
	AccessTokenExpiry  time.Time `json:"accessTokenExpiry"`
	RefreshTokenExpiry time.Time `json:"refreshTokenExpiry"`
	EmailVerified      bool      `json:"emailVerified"`
	Roles              []string  `json:"roles,omitempty"`
}
