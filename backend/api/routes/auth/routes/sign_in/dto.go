package sign_in

import "time"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	AccessTokenExpiry  time.Time `json:"accessTokenExpiry"`
	RefreshTokenExpiry time.Time `json:"refreshTokenExpiry"`
	EmailVerified      bool      `json:"emailVerified"`
	Roles              []string  `json:"roles"`
}

// Backward compatible alias if any file still refers to Response
type Response = SignInResponse
