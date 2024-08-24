package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/utils"
)

type AttachmentRepository interface {
	Save(*models.Attachment) (*models.Attachment, error)
	FindOne(string) (*models.Attachment, error)
	FindByTicket(*[]models.Attachment, uint64) error
	Remove(string) error
}

type attachmentRepository struct {
	db *db.DB
}

func (a *attachmentRepository) FindByTicket(attachments *[]models.Attachment, ticketId uint64) error {
	return a.db.Transaction(func(tx *gorm.DB) error {
		// First, check if the ticket exists
		if err := utils.CheckRelatedRecordExists(tx, &models.TroubleTicket{}, ticketId, "id"); err != nil {
			return err
		}

		// If ticket exists, find attachments
		result := tx.Where("trouble_ticket_id = ?", ticketId).Find(&attachments)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (a *attachmentRepository) Save(attachment *models.Attachment) (*models.Attachment, error) {
	tx := a.db.Begin()
	defer tx.Rollback()

	// Check if the related TroubleTicket record exists
	var troubleTicket models.TroubleTicket
	if err := tx.First(&troubleTicket, attachment.TroubleTicketID).Error; err != nil {
		return nil, fmt.Errorf("trouble ticket: %v", err)
	}

	// Create the Attachment
	if err := tx.Create(attachment).Error; err != nil {
		return nil, fmt.Errorf("attachment: %v", err)
	}

	// Commit the transaction if everything succeeded
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return attachment, nil
}

func (a *attachmentRepository) FindOne(ref string) (*models.Attachment, error) {
	var attachment models.Attachment

	err := a.db.First(&attachment, "ref = ?", ref).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("record does not exist")
	}
	return &attachment, err
}

func (a *attachmentRepository) Remove(ref string) error {
	result := a.db.Where("ref = ?", ref).Delete(&models.Attachment{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("attachment record with ref:%v does not exist", ref)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewAttachmentRepository(db *db.DB) AttachmentRepository {
	return &attachmentRepository{db}
}
