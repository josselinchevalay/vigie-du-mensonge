package sign_in

import "time"

type RequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResponseDTO struct {
	AccessTokenExpiry  time.Time `json:"accessTokenExpiry"`
	RefreshTokenExpiry time.Time `json:"refreshTokenExpiry"`
	EmailVerified      bool      `json:"emailVerified"`
	Roles              []string  `json:"roles"`
}
