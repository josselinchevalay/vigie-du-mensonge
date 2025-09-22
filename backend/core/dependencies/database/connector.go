package database

import (
	"fmt"
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

type PostgresConnector struct {
	DB *gorm.DB
}

func NewConnector(config env.DatabaseConfig) (Connector, error) {
	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}

	dsn := config.DSN()

	gormDB, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm.DB: %v", err)
	}

	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &PostgresConnector{DB: gormDB}, nil
}

func (p *PostgresConnector) GormDB() *gorm.DB { return p.DB }

func (p *PostgresConnector) Migrate() error {
	return p.DB.AutoMigrate(
		&models.Politician{}, &models.Occupation{}, &models.Government{},
		&models.User{}, &models.Role{}, &models.UserRole{}, &models.UserToken{},
		&models.Article{}, &models.ArticlePolitician{}, &models.ArticleReview{}, &models.ArticleTag{}, &models.ArticleSource{},
	)
}

func (p *PostgresConnector) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
