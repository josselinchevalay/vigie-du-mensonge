package env

import (
	"fmt"
	"time"
)

type SecurityConfig struct {
	AccessTokenSecret []byte
	AccessTokenTTL    time.Duration
	RefreshTokenTTL   time.Duration

	EmailVerificationTokenSecret []byte
	EmailVerificationTokenTTL    time.Duration

	PasswordUpdateTokenSecret []byte
	PasswordUpdateTokenTTL    time.Duration

	CsrfTokenSecret []byte
	CsrfTokenTTL    time.Duration
}

func loadSecurityConfig() (SecurityConfig, error) {
	accessTokenTTL, err := time.ParseDuration(getEnv("ACCESS_TOKEN_TTL", "45s"))
	if err != nil {
		return SecurityConfig{}, fmt.Errorf("failed to parse ACCESS_TOKEN_TTL: %v", err)
	}

	refreshTokenTTL, err := time.ParseDuration(getEnv("REFRESH_TOKEN_TTL", "60s"))
	if err != nil {
		return SecurityConfig{}, fmt.Errorf("failed to parse REFRESH_TOKEN_TTL: %v", err)
	}

	emailVerificationTokenTTL, err := time.ParseDuration(getEnv("EMAIL_VERIFICATION_TOKEN_TTL", "45s"))
	if err != nil {
		return SecurityConfig{}, fmt.Errorf("failed to parse EMAIL_VERIFICATION_TOKEN_TTL: %v", err)
	}

	passwordUpdateTokenTTL, err := time.ParseDuration(getEnv("PASSWORD_UPDATE_TOKEN_TTL", "45s"))
	if err != nil {
		return SecurityConfig{}, fmt.Errorf("failed to parse PASSWORD_UPDATE_TOKEN_TTL: %v", err)
	}

	return SecurityConfig{
		AccessTokenSecret:            []byte(getEnv("ACCESS_TOKEN_SECRET", "")),
		AccessTokenTTL:               accessTokenTTL,
		RefreshTokenTTL:              refreshTokenTTL,
		EmailVerificationTokenSecret: []byte(getEnv("EMAIL_VERIFICATION_TOKEN_SECRET", "")),
		EmailVerificationTokenTTL:    emailVerificationTokenTTL,
		PasswordUpdateTokenSecret:    []byte(getEnv("PASSWORD_UPDATE_TOKEN_SECRET", "")),
		PasswordUpdateTokenTTL:       passwordUpdateTokenTTL,
	}, nil
}
