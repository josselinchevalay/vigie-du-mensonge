package env

import (
	"fmt"
	"os"
)

type Config struct {
	ActiveProfile string
	AllowOrigins  string
	BaseURL       string
	Database      DatabaseConfig
	Security      SecurityConfig
	Mailer        MailerConfig
}

func LoadConfig() (Config, error) {
	securityConfig, err := loadSecurityConfig()
	if err != nil {
		return Config{}, fmt.Errorf("failed to load security config: %v", err)
	}

	dbConfig, err := loadDatabaseConfig()
	if err != nil {
		return Config{}, fmt.Errorf("failed to load database config: %v", err)
	}

	mailerConfig, err := loadMailerConfig()
	if err != nil {
		return Config{}, fmt.Errorf("failed to load mailer config: %v", err)
	}

	return Config{
		ActiveProfile: getEnv("ACTIVE_PROFILE", "test"),
		AllowOrigins:  getEnv("ALLOW_ORIGINS", ""),
		BaseURL:       getEnv("BASE_URL", ""),
		Database:      dbConfig,
		Security:      securityConfig,
		Mailer:        mailerConfig,
	}, nil
}

func (e Config) Validate() error {
	//TODO: split validation into separate functions
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
