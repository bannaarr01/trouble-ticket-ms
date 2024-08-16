package migrate

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func Migrate(dbConn *db.DB) {
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
		return
	}
}

//func main() {
//	dbConn := db.Init()
//
//	defer db.CloseDB(dbConn)
//	//migrate
//	migrate(dbConn)
//}
