package database

import (
	"log"
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

func NewPostgresConnector() Connector {
	dsn := "host=localhost port=5432 user=root password=root dbname=vdm_db sslmode=disable"

	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
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
