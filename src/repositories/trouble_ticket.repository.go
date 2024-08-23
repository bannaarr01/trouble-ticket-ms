package repositories

import (
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/utils"
)

type TroubleTicketRepository interface {
	Create(*models.TroubleTicket) error
	FindAll(*[]models.TroubleTicket) error
	FindOne(*models.TroubleTicket, string) error
	Remove(*models.TroubleTicket) error
	Update(*models.TroubleTicket, string) error
	FindAllFilter(*models.Filters) error
}

type troubleTicketRepository struct {
	db *db.DB
}

func (t *troubleTicketRepository) Create(troubleTicket *models.TroubleTicket) error {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketRepository) FindAllFilter(allFilter *models.Filters) error {
	if err := t.db.Find(&allFilter.Types).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Statuses).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Severities).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Channels).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Priorities).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Roles).Error; err != nil {
		return err
	}
	return nil
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

func NewTroubleTicketRepository(db *db.DB) TroubleTicketRepository {
	return &troubleTicketRepository{db}
}
