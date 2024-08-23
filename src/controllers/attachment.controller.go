package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/services"
	"trouble-ticket-ms/src/utils"
)

type AttachmentController interface {
	Upload(*gin.Context)
	FindOne(*gin.Context)
	FindByTicket(*gin.Context)
	Remove(*gin.Context)
}

type attachmentController struct {
	attachmentService services.AttachmentService
}

// Remove an Attachment
// @Summary remove an attachment by its ref
// @Tags Attachments
// @Param ref path string true "Attachment Ref"
// @Success 200 {object} any
// @Failure 500 {object} error
// @Router /attachments/ref/{ref} [delete]
// @Security Bearer
func (a *attachmentController) Remove(context *gin.Context) {
	attachmentRef, err := utils.ParseString(context, "ref")
	if err != nil {
		return // Err resp has already been set
	}

	err = a.attachmentService.Remove(attachmentRef)

	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// FindByTicket Attachments
// @Summary find attachments by a trouble ticket ID
// @Tags Attachments
// @Param ticketId path int true "Trouble Ticket ID"
// @Success 200 {array} []models.AttachmentDTO
// @Failure 500 {object} error
// @Router /attachments/{ticketId} [get]
// @Security Bearer
func (a *attachmentController) FindByTicket(context *gin.Context) {
	troubleTicketID, err := utils.ParseID[uint64](context, "ticketId")
	if err != nil {
		return // Err resp has already been set
	}

	attachments, err := a.attachmentService.FindByTicket(troubleTicketID)

	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Errorf("couldn't fetch attachments: %v", err).Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": attachments})
}

// FindOne Attachment
// @Summary find an attachment by its ref
// @Tags Attachments
// @Param ref path string true "Attachment Ref"
// @Success 200 {object} models.AttachmentDTO
// @Failure 500 {object} error
// @Router /attachments/ref/{ref} [get]
// @Security Bearer
func (a *attachmentController) FindOne(context *gin.Context) {
	attachmentRef, err := utils.ParseString(context, "ref")
	if err != nil {
		return // Err resp has already been set
	}

	foundAttachment, err := a.attachmentService.FindOne(attachmentRef)

	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": foundAttachment})
}

// Upload Attachment
// @Summary upload an attachment for a trouble ticket
// @Tags Attachments
// @Param ticketId path int true "Trouble Ticket ID"
// @Accept multipart/form-data
// @Param file formData file true "Attachment file"
// @Success 200 {object} models.AttachmentDTO
// @Failure 500 {object} error
// @Router /attachments/{ticketId} [post]
// @Security Bearer
func (a *attachmentController) Upload(context *gin.Context) {
	troubleTicketID, err := utils.ParseID[uint64](context, "ticketId")
	if err != nil {
		return // Err resp has already been set
	}

	// available after file middleware
	file, _ := context.MustGet("file").(multipart.File)
	fileHeader, _ := context.MustGet("fileHeader").(*multipart.FileHeader)
	defer file.Close()

	// user will be available after passing through auth guard middleware
	authUser := context.MustGet("user").(*models.Claims)
	// get saved attachment as attachment DTO
	savedAttachment, err := a.attachmentService.Save(troubleTicketID, authUser, &file, fileHeader)

	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": savedAttachment})
}

func NewAttachmentController(at services.AttachmentService) AttachmentController {
	return &attachmentController{at}
}
