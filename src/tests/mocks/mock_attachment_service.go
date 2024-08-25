package mocks

import (
	"mime/multipart"
	"trouble-ticket-ms/src/models"
)

type MockAttachmentService struct {
	RemoveFunc       func(ref string) error
	FindByTicketFunc func(id uint64) ([]models.AttachmentDTO, error)
	FindOneFunc      func(ref string) (*models.AttachmentDTO, error)
	SaveFunc         func(id uint64, user *models.Claims, file *multipart.File, header *multipart.FileHeader) (*models.AttachmentDTO, error)
}

func (m *MockAttachmentService) Remove(ref string) error {
	return m.RemoveFunc(ref)
}

func (m *MockAttachmentService) FindByTicket(id uint64) ([]models.AttachmentDTO, error) {
	return m.FindByTicketFunc(id)
}

func (m *MockAttachmentService) FindOne(ref string) (*models.AttachmentDTO, error) {
	return m.FindOneFunc(ref)
}

func (m *MockAttachmentService) Save(id uint64, user *models.Claims, file *multipart.File, header *multipart.FileHeader) (*models.AttachmentDTO, error) {
	return m.SaveFunc(id, user, file, header)
}
