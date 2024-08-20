package seeds

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func SeedChannels(db *db.DB) {
	channels := []models.Channel{
		models.NewChannel("operation"),
		models.NewChannel("sales"),
		models.NewChannel("support"),
		models.NewChannel("billing"),
		models.NewChannel("HR"),
		models.NewChannel("finance"),
	}

	BulkCreate(db, channels)
}
