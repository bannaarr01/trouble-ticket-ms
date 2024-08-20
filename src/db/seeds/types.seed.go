package seeds

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func SeedTypes(db *db.DB) {
	nTypes := []models.Type{
		models.NewType("incident"),
		models.NewType("complain"),
		models.NewType("request"),
	}

	BulkCreate(db, nTypes)
}
