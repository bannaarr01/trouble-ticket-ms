package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"trouble-ticket-ms/src/services"
)

type TroubleTicketController interface {
	Create(context *gin.Context)
	FindAll(context *gin.Context)
	FindOne(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type troubleTicketController struct {
	troubleTicketService services.TroubleTicketService
}

func (t *troubleTicketController) Create(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// FindAll TroubleTicket
// @Summary fetch all trouble tickets
// @Tags Trouble Tickets
// @Success 200 {array} models.TroubleTicketDTO
// @Failure 500 {object} error
// @Router /troubleTickets [get]
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
