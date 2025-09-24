package database

import (
	"fmt"
	"log"
	"os"
	"vdm/data_import/models"

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
	db *gorm.DB
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func buildDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "vdm_db"), getEnv("DB_PORT", "5432"), getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "root"), getEnv("DB_NAME", "vdm_db"), getEnv("DB_SSL_MODE", "disable"))
}

func NewPostgresConnector() Connector {
	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}

	db, err := gorm.Open(postgres.Open(buildDSN()), gormConfig)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresConnector{db: db}
}

func (p *PostgresConnector) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (p *PostgresConnector) Migrate() error {
	return p.db.AutoMigrate(
		&models.Politician{},
		&models.Occupation{},
		&models.Government{},
	)
}

func (p *PostgresConnector) GormDB() *gorm.DB { return p.db }
