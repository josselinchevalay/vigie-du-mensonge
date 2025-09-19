package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Security struct {
	AccessTokenSecret []byte
	AccessTokenTTL    time.Duration
	RefreshTokenTTL   time.Duration

	EmailVerificationTokenSecret []byte
	EmailVerificationTokenTTL    time.Duration

	PasswordUpdateTokenSecret []byte
	PasswordUpdateTokenTTL    time.Duration
}

type Database struct {
	Host            string
	User            string
	Password        string
	Name            string
	Port            string
	SSLMode         string
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}

func (cfg Database) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)
}

type Env struct {
	ActiveProfile string
	AllowOrigins  string
	Database      Database
	Security      Security
}

var Config = mustLoad()

func mustLoad() Env {
	config, err := load()

	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	if err = config.validate(); err != nil {
		log.Fatalf("invalid env: %v", err)
	}

	return config
}

func load() (Env, error) {
	accessTokenTTL, err := time.ParseDuration(getEnv("ACCESS_TOKEN_TTL", "45s"))
	if err != nil {
		return Env{}, fmt.Errorf("failed to parse ACCESS_TOKEN_TTL: %v", err)
	}

	refreshTokenTTL, err := time.ParseDuration(getEnv("REFRESH_TOKEN_TTL", "60s"))
	if err != nil {
		return Env{}, fmt.Errorf("failed to parse REFRESH_TOKEN_TTL: %v", err)
	}

	emailVerificationTokenTTL, err := time.ParseDuration(getEnv("EMAIL_VERIFICATION_TOKEN_TTL", "45s"))
	if err != nil {
		return Env{}, fmt.Errorf("failed to parse EMAIL_VERIFICATION_TOKEN_TTL: %v", err)
	}

	passwordUpdateTokenTTL, err := time.ParseDuration(getEnv("PASSWORD_UPDATE_TOKEN_TTL", "45s"))
	if err != nil {
		return Env{}, fmt.Errorf("failed to parse PASSWORD_UPDATE_TOKEN_TTL: %v", err)
	}

	securityConfig := Security{
		AccessTokenSecret:            []byte(getEnv("ACCESS_TOKEN_SECRET", "")),
		AccessTokenTTL:               accessTokenTTL,
		RefreshTokenTTL:              refreshTokenTTL,
		EmailVerificationTokenSecret: []byte(getEnv("EMAIL_VERIFICATION_TOKEN_SECRET", "")),
		EmailVerificationTokenTTL:    emailVerificationTokenTTL,
		PasswordUpdateTokenSecret:    []byte(getEnv("PASSWORD_UPDATE_TOKEN_SECRET", "")),
		PasswordUpdateTokenTTL:       passwordUpdateTokenTTL,
	}

	dbConnMaxLifetime, err := time.ParseDuration(getEnv("DB_CONN_MAX_LIFETIME", "10m"))
	if err != nil {
		return Env{}, fmt.Errorf("failed to parse DB_CONN_MAX_LIFETIME: %v", err)
	}

	dbConnMaxIdleTime, err := time.ParseDuration(getEnv("DB_CONN_MAX_IDLE_TIME", "5m"))
	if err != nil {
		return Env{}, fmt.Errorf("failed to parse DB_CONN_MAX_IDLE_TIME: %v", err)
	}

	dbMaxOpenConns, err := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "2"))
	if err != nil {
		return Env{}, fmt.Errorf("failed to parse DB_MAX_OPEN_CONNS: %v", err)
	}

	dbMaxIdleConns, err := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "1"))
	if err != nil {
		return Env{}, fmt.Errorf("failed to parse DB_MAX_IDLE_CONNS: %v", err)
	}

	dbConfig := Database{
		Host:            getEnv("DB_HOST", ""),
		User:            getEnv("DB_USER", ""),
		Password:        getEnv("DB_PASSWORD", ""),
		Name:            getEnv("DB_NAME", ""),
		Port:            getEnv("DB_PORT", ""),
		SSLMode:         getEnv("DB_SSL_MODE", ""),
		ConnMaxLifetime: dbConnMaxLifetime,
		ConnMaxIdleTime: dbConnMaxIdleTime,
		MaxOpenConns:    dbMaxOpenConns,
		MaxIdleConns:    dbMaxIdleConns,
	}

	c := Env{
		ActiveProfile: getEnv("ACTIVE_PROFILE", "test"),
		AllowOrigins:  getEnv("ALLOW_ORIGINS", ""),
		Database:      dbConfig,
		Security:      securityConfig,
	}

	return c, nil
}

func (e Env) validate() error {
	if e.Security.AccessTokenTTL <= 0 {
		return fmt.Errorf("ACCESS_TOKEN_TTL must be > 0")
	}
	if e.Security.RefreshTokenTTL <= 0 {
		return fmt.Errorf("REFRESH_TOKEN_TTL must be > 0")
	}
	if e.ActiveProfile == "prod" {
		if len(e.Security.AccessTokenSecret) == 0 {
			return fmt.Errorf("ACCESS_TOKEN_SECRET is required in prod")
		}
		if len(e.Security.EmailVerificationTokenSecret) == 0 {
			return fmt.Errorf("EMAIL_VERIFICATION_TOKEN_SECRET is required in prod")
		}
		if len(e.Security.PasswordUpdateTokenSecret) == 0 {
			return fmt.Errorf("PASSWORD_UPDATE_TOKEN_SECRET is required in prod")
		}
		if e.Database.Host == "" || e.Database.User == "" || e.Database.Name == "" {
			return fmt.Errorf("DB_* vars (host,user,name) are required in prod")
		}
	}
	if e.Database.MaxIdleConns > e.Database.MaxOpenConns {
		return fmt.Errorf("DB_MAX_IDLE_CONNS cannot exceed DB_MAX_OPEN_CONNS")
	}
	return nil
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
