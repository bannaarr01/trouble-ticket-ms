package seeds

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func SeedStatuses(db *db.DB) {
	const defaultFilter = 1
	statuses := []models.Status{
		models.NewStatus("acknowledged", 1, defaultFilter),
		models.NewStatus("inProgress", 2, defaultFilter),
		models.NewStatus("pending", 3, defaultFilter),
		models.NewStatus("escalated", 4, defaultFilter),
		models.NewStatus("escalatedProgressed", 5, defaultFilter),
		models.NewStatus("resolved", 5, defaultFilter),
	}

	BulkCreate(db, statuses)
}
