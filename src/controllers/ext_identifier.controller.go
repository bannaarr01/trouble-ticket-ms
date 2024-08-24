package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/services"
	"trouble-ticket-ms/src/utils"
)

type ExtIdentifierController interface {
	Create(*gin.Context)
	FindOne(*gin.Context)
	FindByTicket(*gin.Context)
	Remove(*gin.Context)
}

type extIdentifierController struct {
	extIdentifierService services.ExtIdentifierService
}

// Create External Identifier
// @Summary create an external Identifier for a trouble ticket
// @Tags External Identifiers
// @Param id path int true "Trouble Ticket ID"
// @Param   request body     models.CreateExternalIdentifierDTO true  "External Identifier Info"
// @Success 200 {object} models.ExternalIdentifierDTO
// @Failure 500 {object} error
// @Router /externalIdentifiers/ticket/{id} [post]
// @Security Bearer
func (e *extIdentifierController) Create(context *gin.Context) {
	troubleTicketID, err := utils.ParseID[uint64](context, "id")
	if err != nil {
		return // Err resp has already been set
	}

	var createExternalIdentifierDTO models.CreateExternalIdentifierDTO

	if !utils.BindJSON(context, &createExternalIdentifierDTO) {
		return
	}

	authUser := context.MustGet("user").(*models.Claims)

	extIdentifier, err := e.extIdentifierService.Create(authUser.PreferredUsername, troubleTicketID, &createExternalIdentifierDTO)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": extIdentifier})
}

func (e *extIdentifierController) FindOne(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e *extIdentifierController) FindByTicket(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e *extIdentifierController) Remove(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewExtIdentifierController(ext services.ExtIdentifierService) ExtIdentifierController {
	return &extIdentifierController{ext}
}
