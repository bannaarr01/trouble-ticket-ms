package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/utils"
)

type ExtIdentifierRepository interface {
	Save(*models.ExternalIdentifier) error
	FindByTicket(*[]models.ExternalIdentifier, uint64) error
	Remove(uint64) error
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

func (e *extIdentifierRepository) FindByTicket(extIdentifier *[]models.ExternalIdentifier, ticketId uint64) error {
	return e.db.Transaction(func(tx *gorm.DB) error {
		// Check if the related record exists
		if err := utils.CheckRelatedRecordExists(tx, &models.TroubleTicket{}, ticketId, "id"); err != nil {
			return err
		}

		// If ticket exists, find extId
		result := tx.Where("trouble_ticket_id = ?", ticketId).
			Preload("Type").
			Find(&extIdentifier)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (e *extIdentifierRepository) Remove(extIdentifierID uint64) error {
	result := e.db.Where("id = ?", extIdentifierID).Delete(&models.ExternalIdentifier{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("external Identifier record with id:%d does not exist", extIdentifierID)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewExtIdentifierRepository(db *db.DB) ExtIdentifierRepository {
	return &extIdentifierRepository{db}
}
