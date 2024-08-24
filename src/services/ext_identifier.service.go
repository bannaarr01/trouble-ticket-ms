package services

import (
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/repositories"
)

type ExtIdentifierService interface {
	Create(string, uint64, *models.CreateExternalIdentifierDTO) (*models.ExternalIdentifierDTO, error)
	FindOne(string) (*models.AttachmentDTO, error)
	FindByTicket(uint64) ([]models.AttachmentDTO, error)
	Remove(string) error
}

type extIdentifierService struct {
	extIdentifierRepository repositories.ExtIdentifierRepository
	deps                    AppDependencies
}

func (e *extIdentifierService) Create(authUserName string, troubleTicketId uint64, cExtDto *models.CreateExternalIdentifierDTO) (*models.ExternalIdentifierDTO, error) {
	externalIdentifier := models.NewExternalIdentifier(troubleTicketId, cExtDto, models.SetField("CreatedBy", authUserName))

	err := e.extIdentifierRepository.Save(&externalIdentifier)
	if err != nil {
		return nil, err
	}
	externalIdentifierDto := models.NewExternalIdentifierDTO(externalIdentifier)

	return &externalIdentifierDto, nil
}

func (e *extIdentifierService) FindOne(s string) (*models.AttachmentDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (e *extIdentifierService) FindByTicket(u uint64) ([]models.AttachmentDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (e *extIdentifierService) Remove(s string) error {
	//TODO implement me
	panic("implement me")
}

func NewExtIdentifierService(extRepo repositories.ExtIdentifierRepository, deps AppDependencies) ExtIdentifierService {
	return &extIdentifierService{extRepo, deps}
}
