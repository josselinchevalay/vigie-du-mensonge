package env

import (
	"fmt"
	"time"
)

type SecurityConfig struct {
	EmailVerificationTokenSecret []byte
	EmailVerificationTokenTTL    time.Duration

	PasswordUpdateTokenSecret []byte
	PasswordUpdateTokenTTL    time.Duration

	AccessTokenSecret []byte
	AccessTokenTTL    time.Duration

	RefreshTokenTTL time.Duration

	RefreshCookieName string
	AccessCookieName  string
	CsrfCookieName    string
	CookieSameSite    string
	CookieSecure      bool
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
		EmailVerificationTokenSecret: []byte(getEnv("EMAIL_VERIFICATION_TOKEN_SECRET", "")),
		EmailVerificationTokenTTL:    emailVerificationTokenTTL,
		PasswordUpdateTokenSecret:    []byte(getEnv("PASSWORD_UPDATE_TOKEN_SECRET", "")),
		PasswordUpdateTokenTTL:       passwordUpdateTokenTTL,
		AccessTokenSecret:            []byte(getEnv("ACCESS_TOKEN_SECRET", "")),
		AccessTokenTTL:               accessTokenTTL,
		AccessCookieName:             getEnv("ACCESS_COOKIE_NAME", "__Host-jwt"),
		RefreshCookieName:            getEnv("REFRESH_COOKIE_NAME", "__Host-rft"),
		RefreshTokenTTL:              refreshTokenTTL,
		CookieSameSite:               getEnv("COOKIE_SAME_SITE", "strict"),
		CookieSecure:                 getEnv("COOKIE_SECURE", "true") == "true",
		CsrfCookieName:               getEnv("CSRF_COOKIE_NAME", "__Host-csrf_"),
	}, nil
}
