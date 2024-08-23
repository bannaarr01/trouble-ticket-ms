package repositories

import (
	"fmt"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
)

type AttachmentRepository interface {
	Save(*models.Attachment) (*models.Attachment, error)
	FindOne(*models.Attachment, string) error
	Remove(*models.Attachment) error
}

type attachmentRepository struct {
	db *db.DB
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

func (a *attachmentRepository) FindOne(attachment *models.Attachment, s string) error {
	//TODO implement me
	panic("implement me")
}

func (a *attachmentRepository) Remove(attachment *models.Attachment) error {
	//TODO implement me
	panic("implement me")
}

func NewAttachmentRepository(db *db.DB) AttachmentRepository {
	return &attachmentRepository{db}
}
