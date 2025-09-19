package refresh

import "time"

type Response struct {
	AccessTokenExpiry  time.Time `json:"accessTokenExpiry"`
	RefreshTokenExpiry time.Time `json:"refreshTokenExpiry"`
	EmailVerified      bool      `json:"emailVerified"`
	Roles              []string  `json:"roles"`
}
