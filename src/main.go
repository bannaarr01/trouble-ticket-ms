package main

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/db/migrate"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// DB connection
	dbConn := db.Init()

	err := dbConn.CheckMigration()
	if err != nil {
		migrate.Run(dbConn)
	}

	// start app
	bootstrap(dbConn)
}
