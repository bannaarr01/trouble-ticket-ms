package migrate

import (
	"log"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func migrate(dbConn *db.DB) {
	err := dbConn.AutoMigrate(
		&models.TroubleTicket{}, // all its ref tables
		&models.ExternalIdentifier{},
		&models.RelatedEntity{},
		&models.RelatedParty{},
		&models.StatusChange{},
		&models.Attachment{},
		&models.Note{},
	)
	if err != nil {
		log.Panic(err)
	}
}

func Run(dbConn *db.DB) {
	//defer db.CloseDB(dbConn)
	log.Println("Applying Migration...")
	migrate(dbConn)
	log.Println("Migrated successfully!")
}
