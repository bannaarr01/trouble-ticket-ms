package seeds

import (
	"log"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

func Run(db *db.DB) {
	log.Println("Seeding init...")

	if NeedsSeeding(db, &models.Role{}) {
		SeedRoles(db)
	}

	if NeedsSeeding(db, &models.Type{}) {
		SeedTypes(db)
	}

	if NeedsSeeding(db, &models.Status{}) {
		SeedStatuses(db)
	}

	if NeedsSeeding(db, &models.Channel{}) {
		SeedChannels(db)
	}

	if NeedsSeeding(db, &models.Severity{}) {
		SeedSeverities(db)
	}

	if NeedsSeeding(db, &models.Priority{}) {
		SeedPriorities(db)
	}

	log.Println("Seeding completed successfully!")
}

func NeedsSeeding[T any](db *db.DB, model *T) bool {
	var count int64
	if err := db.Model(model).Count(&count).Error; err != nil {
		log.Panic(err)
	}
	return count == 0
}

func BulkCreate[T any](db *db.DB, models []T) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	for _, model := range models {
		if err := tx.Create(&model).Error; err != nil {
			tx.Rollback()
			log.Panic(err)
		}
	}
}
