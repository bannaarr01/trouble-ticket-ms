package seeds

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func SeedRoles(db *db.DB) {
	const defaultFilter = 1
	roles := []models.Role{
		models.NewRole("initiator", 1, defaultFilter),
		models.NewRole("customer", 2, defaultFilter),
		models.NewRole("admin", 3, defaultFilter),
		models.NewRole("super_admin", 4, defaultFilter),
		models.NewRole("assigner", 5, defaultFilter),
	}

	BulkCreate(db, roles)
}
