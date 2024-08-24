package repositories

import (
	"fmt"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/utils"
)

type ExtIdentifierRepository interface {
	Save(*models.ExternalIdentifier) error
	FindOne(string) (*models.ExternalIdentifier, error)
	FindByTicket(*[]models.ExternalIdentifier, uint64) error
	Remove(string) error
}

type extIdentifierRepository struct {
	db *db.DB
}

func (e *extIdentifierRepository) Save(extId *models.ExternalIdentifier) error {
	tx := e.db.Begin()
	defer tx.Rollback()

	// Check if the related record exists
	if err := utils.CheckRelatedRecordExists(tx, &models.TroubleTicket{}, extId.TroubleTicketID, "id"); err != nil {
		return err
	}

	if err := utils.CheckRelatedRecordExists(tx, &models.Type{}, extId.TypeID, "id"); err != nil {
		return err
	}

	if err := tx.Create(&extId).Error; err != nil {
		return fmt.Errorf("external identifier: %v", err)
	}

	// Preload the Type after creation
	if err := tx.Preload("Type").First(&extId).Error; err != nil {
		return fmt.Errorf("failed to retrieve Type: %v", err)
	}

	// Commit the transaction if everything succeeded
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (e *extIdentifierRepository) FindOne(s string) (*models.ExternalIdentifier, error) {
	//TODO implement me
	panic("implement me")
}

func (e *extIdentifierRepository) FindByTicket(extIdentifier *[]models.ExternalIdentifier, u uint64) error {
	//TODO implement me
	panic("implement me")
}

func (e *extIdentifierRepository) Remove(s string) error {
	//TODO implement me
	panic("implement me")
}

func NewExtIdentifierRepository(db *db.DB) ExtIdentifierRepository {
	return &extIdentifierRepository{db}
}
