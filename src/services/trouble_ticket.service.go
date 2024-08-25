package services

import (
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/repositories"
	"trouble-ticket-ms/src/utils"
)

type TroubleTicketService interface {
	Create(string, *models.CreateTroubleTicketDTO) (*models.TroubleTicketDTO, error)
	FindAll(*models.Claims, *models.GetTroubleTicketQuery) (*models.PaginatedTroubleTickets, error)
	FindOne(uint64, *models.Claims) (*models.TroubleTicketDTO, error)
	Remove(uint64, *models.Claims) error
	FindAllFilter() (models.FiltersDTO, error)
	Update(uint64, *models.Claims, *models.UpdateTroubleTicketDTO) (*models.TroubleTicketDTO, error)
}

type troubleTicketService struct {
	troubleTicketRepository repositories.TroubleTicketRepository
}

func (t *troubleTicketService) Update(ticketId uint64, authUser *models.Claims, updateDto *models.UpdateTroubleTicketDTO) (*models.TroubleTicketDTO, error) {
	updatedTroubleTicket, err := t.troubleTicketRepository.Update(ticketId, authUser, updateDto)
	if err != nil {
		return nil, err
	}

	updatedTicketDTO := models.NewTroubleTicketDTO(updatedTroubleTicket)

	return &updatedTicketDTO, nil
}

func (t *troubleTicketService) FindAllFilter() (models.FiltersDTO, error) {
	var filters models.Filters

	if err := t.troubleTicketRepository.FindAllFilter(&filters); err != nil {
		return models.FiltersDTO{}, err
	}

	// Convert Filters to FiltersDTO
	filterDto := models.NewFilterDTO(filters)

	return filterDto, nil
}

// FindAll retrieves all trouble tickets based on query params and return PaginatedTroubleTickets
func (t *troubleTicketService) FindAll(authUser *models.Claims, query *models.GetTroubleTicketQuery) (*models.PaginatedTroubleTickets, error) {
	var troubleTickets []models.TroubleTicket

	totalCount, err := t.troubleTicketRepository.FindAll(authUser, query, &troubleTickets)
	if err != nil {
		return nil, err
	}

	troubleTicketsDTOs := utils.TransformToDTO(troubleTickets,
		func(ticket models.TroubleTicket) models.TroubleTicketDTO {
			return models.NewTroubleTicketDTO(&ticket)
		})

	return &models.PaginatedTroubleTickets{
		TotalCount: totalCount,
		Limit:      query.Limit,
		Offset:     query.Offset,
		Data:       troubleTicketsDTOs,
	}, nil
}

func (t *troubleTicketService) FindOne(ticketId uint64, authUser *models.Claims) (*models.TroubleTicketDTO, error) {
	foundTicket, err := t.troubleTicketRepository.FindOne(ticketId, authUser)
	if err != nil {
		return nil, err
	}

	troubleTicketDTO := models.NewTroubleTicketDTO(foundTicket)
	return &troubleTicketDTO, nil
}

func (t *troubleTicketService) Create(authUserName string, cDto *models.CreateTroubleTicketDTO) (*models.TroubleTicketDTO, error) {
	ticket, err := t.troubleTicketRepository.Create(authUserName, cDto)

	if err != nil {
		return nil, err
	}

	troubleTicketDTO := models.NewTroubleTicketDTO(ticket)

	return &troubleTicketDTO, nil
}

func (t *troubleTicketService) Remove(ticketId uint64, authUser *models.Claims) error {
	err := t.troubleTicketRepository.Remove(ticketId, authUser)

	if err != nil {
		return err
	}

	return nil
}

func NewTroubleTicketService(tRepo repositories.TroubleTicketRepository) TroubleTicketService {
	return &troubleTicketService{tRepo}
}
