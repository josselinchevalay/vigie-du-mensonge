package dependencies

import (
	"vdm/core/dependencies/database"
	"vdm/core/dependencies/mailer"

	"gorm.io/gorm"
)

type Dependencies struct {
	dbConnector database.Connector
	Mailer      mailer.Mailer
}

func (d *Dependencies) GormDB() *gorm.DB {
	return d.dbConnector.GormDB()
}

func New(dbConnector database.Connector, mailer mailer.Mailer) *Dependencies {
	return &Dependencies{
		dbConnector: dbConnector,
		Mailer:      mailer,
	}
}
