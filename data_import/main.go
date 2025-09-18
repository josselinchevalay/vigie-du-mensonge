package main

import (
	"fmt"
	"vdm/data_import/database"
	"vdm/data_import/import_governments"
	"vdm/data_import/import_occupations"
	"vdm/data_import/import_presidents"
)

func main() {
	dbConn := database.NewPostgresConnector()
	defer func(dbConn database.Connector) {
		err := dbConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(dbConn)

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
