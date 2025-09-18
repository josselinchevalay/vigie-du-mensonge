package database

import (
	"fmt"
	"time"
	"vdm/core/env"
	"vdm/core/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Connector interface {
	Close() error
	Migrate() error
	GormDB() *gorm.DB
}

type pgConnector struct {
	db *gorm.DB
}

func NewConnector() (Connector, error) {
	dsn := "host=" + env.DatabaseHost() +
		" user=" + env.DatabaseUser() +
		" password=" + env.DatabasePassword() +
		" dbname=" + env.DatabaseName() +
		" port=" + env.DatabasePort() +
		" sslmode=" + env.DatabaseSSLMode()

	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm.DB: %v", err)
	}

	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 30)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &pgConnector{db: gormDB}, nil
}

func (p *pgConnector) GormDB() *gorm.DB { return p.db }

func (p *pgConnector) Migrate() error {
	return p.db.AutoMigrate(
		&models.Politician{}, &models.Occupation{}, &models.Government{},
		&models.User{}, &models.Role{}, &models.UserRole{},
		&models.Article{}, &models.ArticlePolitician{}, &models.ArticleReview{}, &models.ArticleTag{}, &models.ArticleSource{},
	)
}

func (p *pgConnector) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
