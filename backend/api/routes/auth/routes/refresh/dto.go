package refresh

import "time"

type RefreshResponse struct {
	AccessTokenExpiry  time.Time `json:"accessTokenExpiry"`
	RefreshTokenExpiry time.Time `json:"refreshTokenExpiry"`
	EmailVerified      bool      `json:"emailVerified"`
	Roles              []string  `json:"roles"`
}

// Backward compatible alias if handler or other files still refer to Response
// This can be removed once all references are updated to RefreshResponse.
type Response = RefreshResponse
