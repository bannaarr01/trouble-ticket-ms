package seeds

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func SeedPriorities(db *db.DB) {
	priorities := []models.Priority{
		models.NewPriority("critical", 1),
		models.NewPriority("high", 2),
		models.NewPriority("medium", 3),
		models.NewPriority("low", 4),
	}

	BulkCreate(db, priorities)
}
