package services

import (
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/repositories"
	"trouble-ticket-ms/src/utils"
)

type ExtIdentifierService interface {
	Create(string, uint64, *models.CreateExternalIdentifierDTO) (*models.ExternalIdentifierDTO, error)
	FindByTicket(uint64) ([]models.ExternalIdentifierDTO, error)
	Remove(uint64) error
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

func (e *extIdentifierService) FindByTicket(ticketId uint64) ([]models.ExternalIdentifierDTO, error) {
	var extIdentifiers []models.ExternalIdentifier
	err := e.extIdentifierRepository.FindByTicket(&extIdentifiers, ticketId)

	if err != nil {
		return nil, err
	}
	extIdentifiersDTOs := utils.TransformToDTO(extIdentifiers, models.NewExternalIdentifierDTO)
	return extIdentifiersDTOs, nil
}

func (e *extIdentifierService) Remove(extIdentifierID uint64) error {
	err := e.extIdentifierRepository.Remove(extIdentifierID)

	if err != nil {
		return err
	}

	return nil
}

func NewExtIdentifierService(extRepo repositories.ExtIdentifierRepository, deps AppDependencies) ExtIdentifierService {
	return &extIdentifierService{extRepo, deps}
}
