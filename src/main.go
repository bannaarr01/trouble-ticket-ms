package main

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/db/migrate"
	"trouble-ticket-ms/src/db/seeds"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// DB connection
	dbConn := db.Init()

	if !dbConn.MigrationUpToDate() {
		// migrate
		migrate.Run(dbConn)
		// seed
		seeds.Run(dbConn)
	}

	// start app
	bootstrap(dbConn)
}
