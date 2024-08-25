package mocks

import (
	"trouble-ticket-ms/src/models"
)

type MockExtIdentifierRepository struct {
	SaveFunc         func(*models.ExternalIdentifier) error
	FindByTicketFunc func(*[]models.ExternalIdentifier, uint64) error
	RemoveFunc       func(uint64) error
}

func (m *MockExtIdentifierRepository) Save(extId *models.ExternalIdentifier) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(extId)
	}
	return nil
}

func (m *MockExtIdentifierRepository) FindByTicket(extIdentifier *[]models.ExternalIdentifier, ticketId uint64) error {
	if m.FindByTicketFunc != nil {
		return m.FindByTicketFunc(extIdentifier, ticketId)
	}
	return nil
}

func (m *MockExtIdentifierRepository) Remove(extIdentifierID uint64) error {
	if m.RemoveFunc != nil {
		return m.RemoveFunc(extIdentifierID)
	}
	return nil
}
