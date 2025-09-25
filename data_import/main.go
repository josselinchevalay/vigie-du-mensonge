package main

import (
	"fmt"
	"vdm/data_import/database"
	"vdm/data_import/import_governments"
	"vdm/data_import/import_occupations"
	"vdm/data_import/import_presidents"
	"vdm/data_import/models"
)

func main() {
	dbConn := database.NewPostgresConnector()
	defer func(dbConn database.Connector) {
		err := dbConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(dbConn)

	var loaded bool
	if err := dbConn.GormDB().Model(&models.Politician{}).
		Select("count(*) > 0").
		Find(&loaded).Error; err != nil {
		panic(err)
	}

	if loaded {
		fmt.Println("Data already loaded")
		return
	}

	if err := import_presidents.LoadFromCSV(dbConn.GormDB()); err != nil {
		panic(err)
	}

	if err := import_governments.LoadFromCSV(dbConn.GormDB()); err != nil {
		panic(err)
	}

	if err := import_occupations.LoadFromCSV(dbConn.GormDB()); err != nil {
		panic(err)
	}

	fmt.Println("Done")
}
