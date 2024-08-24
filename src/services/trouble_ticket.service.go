package services

import (
	"golang.org/x/sync/errgroup"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/repositories"
)

type TroubleTicketService interface {
	Create(string, *models.CreateTroubleTicketDTO) (*models.TroubleTicketDTO, error)
	FindAll() ([]models.TroubleTicketDTO, error)
	FindOne()
	Remove()
	FindAllFilter() (models.FiltersDTO, error)
}

type troubleTicketService struct {
	troubleTicketRepository repositories.TroubleTicketRepository
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

// FindAll retrieves all trouble tickets from the trouble ticket repo and returns them as a slice of TroubleTicketDTOs.
func (t *troubleTicketService) FindAll() ([]models.TroubleTicketDTO, error) {
	var troubleTickets []models.TroubleTicket

	if err := t.troubleTicketRepository.FindAll(&troubleTickets); err != nil {
		return nil, err
	}

	// Create an errgroup.Group to manage the concurrent conversion of trouble tickets to DTOs.
	var g errgroup.Group
	var troubleTicketDTOs []models.TroubleTicketDTO

	// Iterate over the retrieved trouble tickets.
	for _, trbTicket := range troubleTickets {
		// Capture the loop variable to avoid issues with concurrent access.
		ticket := trbTicket

		// Add a new goroutine to the errgroup that converts the tickets to a DTO and appends it to the result slice.
		g.Go(func() error {
			dto := models.NewTroubleTicketDTO(&ticket)
			troubleTicketDTOs = append(troubleTicketDTOs, dto)
			return nil
		})
	}

	// Wait for all goroutines in the errgroup to complete.
	// If any goroutine returns an error, return immediately with the error.
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Return the slice of trouble ticket DTOs.
	return troubleTicketDTOs, nil
}

func (t *troubleTicketService) FindOne() {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketService) Create(authUserName string, cDto *models.CreateTroubleTicketDTO) (*models.TroubleTicketDTO, error) {
	ticket, err := t.troubleTicketRepository.Create(authUserName, cDto)

	if err != nil {
		return nil, err
	}

	troubleTicketDTO := models.NewTroubleTicketDTO(ticket)

	return &troubleTicketDTO, nil
}

func (t *troubleTicketService) Remove() {
	//TODO implement me
	panic("implement me")
}

func NewTroubleTicketService(tRepo repositories.TroubleTicketRepository) TroubleTicketService {
	return &troubleTicketService{tRepo}
}
