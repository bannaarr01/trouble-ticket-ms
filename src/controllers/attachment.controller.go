package controllers

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/services"
	"trouble-ticket-ms/src/utils"
)

type AttachmentController interface {
	Upload(*gin.Context)
}

type attachmentController struct {
	attachmentService services.AttachmentService
}

// Upload Attachment
// @Summary upload an attachment for a trouble ticket
// @Tags Attachments
// @Param ticketId path int true "Attachment ID"
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
