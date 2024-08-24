package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/services"
	"trouble-ticket-ms/src/utils"
)

type TroubleTicketController interface {
	Create(*gin.Context)
	FindAll(*gin.Context)
	FindOne(*gin.Context)
	Update(*gin.Context)
	Remove(*gin.Context)
	FindAllFilter(*gin.Context)
}

type troubleTicketController struct {
	troubleTicketService services.TroubleTicketService
}

// Create new trouble ticket
// @Summary Create a trouble ticket
// @Tags Trouble Tickets
// @Param  request body  models.CreateTroubleTicketDTO  true  "Create New Ticket info"
// @Success 200 {object} models.TroubleTicketDTO
// @Failure 500 {object} error
// @Router /troubleTickets [post]
// @Security Bearer
func (t *troubleTicketController) Create(context *gin.Context) {
	var createTroubleTicketDTO models.CreateTroubleTicketDTO

	if !utils.BindJSON(context, &createTroubleTicketDTO) {
		return
	}

	authUser := context.MustGet("user").(*models.Claims)

	createdTicket, err := t.troubleTicketService.Create(authUser.PreferredUsername, &createTroubleTicketDTO)

	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": createdTicket})
}

// FindAllFilter related to TroubleTicket
// @Summary fetch all related trouble tickets filters / dropdown
// @Tags Trouble Tickets
// @Success 200 {array} models.FiltersDTO
// @Failure 500 {object} error
// @Router /troubleTickets/filters [get]
// @Security Bearer
func (t *troubleTicketController) FindAllFilter(context *gin.Context) {
	fmt.Println("Controller: FindAllFilter invoked")
	filters, err := t.troubleTicketService.FindAllFilter()
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't fetch trouble ticket filters"})
		return
	}
	context.Set("data", filters)

	context.JSON(http.StatusOK, gin.H{"data": filters})
}

// FindAll TroubleTicket
// @Summary fetch all trouble tickets
// @Tags Trouble Tickets
// @Success 200 {array} models.TroubleTicketDTO
// @Failure 500 {object} error
// @Router /troubleTickets [get]
// @Security Bearer
func (t *troubleTicketController) FindAll(context *gin.Context) {
	allTroubleTickets, err := t.troubleTicketService.FindAll()
	if err != nil {
		log.Printf("error fetching all trouble tickets: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't fetch all trouble tickets"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": allTroubleTickets})
}

func (t *troubleTicketController) FindOne(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketController) Update(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketController) Remove(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewTroubleTicketController(ts services.TroubleTicketService) TroubleTicketController {
	return &troubleTicketController{ts}
}
