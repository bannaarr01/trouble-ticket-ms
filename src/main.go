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

// Same template as external identifier
//TODO: Endpoint for related parties for ticket, tighten each endpoint more
//TODO: Add Notes to ticket endpoints n logic
//TODO: Add RelatedEntities to ticket endpoints n logic

// More Todo Internal ....
// TODO: Ticket Assignment By system Auto by cronjob task n Assigner. Send email to the assigned n related parties update status
// TODO: Notify also when resolved
