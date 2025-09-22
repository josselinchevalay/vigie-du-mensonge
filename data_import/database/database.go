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

func buildDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))
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
