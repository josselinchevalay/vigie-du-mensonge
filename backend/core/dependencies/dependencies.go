package dependencies

import (
	"vdm/core/dependencies/database"
	"vdm/core/dependencies/mailer"
	"vdm/core/env"

	"gorm.io/gorm"
)

type Dependencies struct {
	Config      env.Config
	dbConnector database.Connector
	Mailer      mailer.Mailer
}

func (d *Dependencies) GormDB() *gorm.DB {
	return d.dbConnector.GormDB()
}

func New(cfg env.Config, dbConnector database.Connector, mailer mailer.Mailer) *Dependencies {
	return &Dependencies{
		Config:      cfg,
		dbConnector: dbConnector,
		Mailer:      mailer,
	}
}
