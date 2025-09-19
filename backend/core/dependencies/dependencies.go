package dependencies

import (
	"vdm/core/dependencies/database"

	"gorm.io/gorm"
)

type Dependencies struct {
	dbConnector database.Connector
}

func (d *Dependencies) GormDB() *gorm.DB {
	return d.dbConnector.GormDB()
}

func New(dbConnector database.Connector) *Dependencies {
	return &Dependencies{
		dbConnector: dbConnector,
	}
}
