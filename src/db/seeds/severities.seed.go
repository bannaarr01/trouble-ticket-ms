package seeds

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func SeedSeverities(db *db.DB) {
	severities := []models.Severity{
		models.NewSeverity("critical"),
		models.NewSeverity("major"),
		models.NewSeverity("minor"),
	}

	BulkCreate(db, severities)
}
