package repositories

import (
	"gorm.io/gorm"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/utils"
)

type TroubleTicketRepository interface {
	Create(troubleTicket *models.TroubleTicket) error
	FindAll(troubleTicket *[]models.TroubleTicket) error
	FindOne(troubleTicket *models.TroubleTicket, id string) error
	Remove(troubleTicket *models.TroubleTicket) error
	Update(troubleTicket *models.TroubleTicket, id string) error
}

type troubleTicketRepository struct {
	db *gorm.DB
}

func (t *troubleTicketRepository) Create(troubleTicket *models.TroubleTicket) error {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketRepository) FindAll(troubleTicket *[]models.TroubleTicket) error {
	if err := t.db.
		Preload("Type").
		Preload("Status").
		Preload("Channel").
		Preload("Severity").
		Preload("Priority").
		Preload("RelatedEntities").
		Preload("Attachments").
		Preload("Notes").
		Preload("StatusChanges", utils.NestedPreload("Status")).
		Preload("RelatedParties", utils.NestedPreload("Party", "Role")).
		Preload("ExternalIdentifiers", utils.NestedPreload("Type")).
		Find(&troubleTicket).Error; err != nil {
		return err
	}

	return nil
}

func (t *troubleTicketRepository) FindOne(troubleTicket *models.TroubleTicket, id string) error {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketRepository) Remove(troubleTicket *models.TroubleTicket) error {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketRepository) Update(troubleTicket *models.TroubleTicket, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTroubleTicketRepository(db *gorm.DB) TroubleTicketRepository {
	return &troubleTicketRepository{db}
}
