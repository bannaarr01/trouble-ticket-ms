package services

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/repositories"
	"trouble-ticket-ms/src/utils"
)

type AttachmentService interface {
	Save(uint64, *models.Claims, *multipart.File, *multipart.FileHeader) (*models.AttachmentDTO, error)
	FindOne(string) (*models.AttachmentDTO, error)
	FindByTicket(uint64) ([]models.AttachmentDTO, error)
	Remove(string) error
}

type attachmentService struct {
	attachmentRepository repositories.AttachmentRepository
	deps                 AppDependencies
}

func (a *attachmentService) Remove(ref string) error {
	err := a.attachmentRepository.Remove(ref)

	if err != nil {
		return err
	}

	return nil
}

func (a *attachmentService) FindByTicket(ticketId uint64) ([]models.AttachmentDTO, error) {
	var attachments []models.Attachment
	err := a.attachmentRepository.FindByTicket(&attachments, ticketId)

	if err != nil {
		return nil, err
	}
	attachmentDTOs := utils.TransformToDTO(attachments, models.NewAttachmentDTO)
	return attachmentDTOs, nil
}

func (a *attachmentService) FindOne(ref string) (*models.AttachmentDTO, error) {
	foundAttachment, err := a.attachmentRepository.FindOne(ref)
	if err != nil {
		return nil, fmt.Errorf("error retrieving attachment with ref %s: %w", ref, err)
	}

	attachmentDTO := models.NewAttachmentDTO(*foundAttachment)
	return &attachmentDTO, nil
}

func createAttachment(
	user *models.Claims,
	fileHeader *multipart.FileHeader,
	troubleTicketId uint64,
	attachmentHost, uniqueFileName, fileExt, dirPath string,
) *models.Attachment {

	size := fileHeader.Size
	ref := generateUniqueFileName()
	mimeType := getMimeType(fileExt)
	url := generateURL(attachmentHost, uniqueFileName, fileExt)

	return &models.Attachment{
		BaseModel: models.BaseModel{
			CreatedBy: user.PreferredUsername,
		},
		Ref:             ref,
		Type:            fileExt,
		MimeType:        mimeType,
		OriginalName:    fileHeader.Filename,
		Path:            dirPath,
		Size:            uint64(size),
		Name:            uniqueFileName,
		URL:             url,
		Description:     "Trouble Ticket Attachment",
		TroubleTicketID: troubleTicketId,
	}
}

func generateURL(attachmentHost, uniqueFileName, fileExt string) string {
	return attachmentHost + "/static/attachment/file/" + uniqueFileName + fileExt
}

func getMimeType(fileExtension string) string {
	mimeType := mime.TypeByExtension(fileExtension)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	return mimeType
}

func generateUniqueFileName() string {
	// Combine UUID and timestamp to create a unique filename
	return fmt.Sprintf("%d-%s", time.Now().Unix(), uuid.New().String())
}

func writeFile(filePath string, content []byte) error {
	err := os.WriteFile(filePath, content, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %w", err)
	}
	return nil
}

// createDirIfNotExists checks app root dir if data dir exists, creates if not & returns its path
func createDirIfNotExists() (*string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("error preparing attachment location") //get current working directory
	}

	// at project root
	dirPath := filepath.Join(cwd, "data")

	if _, err = os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return nil, fmt.Errorf("error making attachment location") // creating directory error
		}
	}
	return &dirPath, nil
}

func (a *attachmentService) Save(troubleTicketID uint64, user *models.Claims, file *multipart.File, fileHeader *multipart.FileHeader) (*models.AttachmentDTO, error) {
	dirPath, err := createDirIfNotExists()
	if err != nil {
		return nil, err
	}

	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return nil, fmt.Errorf("error reading file content: %w", err)
	}

	fileName := fileHeader.Filename
	uniqueFileName := generateUniqueFileName()
	fileExt := filepath.Ext(fileName)
	filePath := filepath.Join(*dirPath, uniqueFileName+fileExt)
	attachmentHost := a.deps.AppConfig.AttachmentHost

	err = writeFile(filePath, fileBytes)
	if err != nil {
		return nil, err
	}

	attachment := createAttachment(user, fileHeader, troubleTicketID, attachmentHost, uniqueFileName, fileExt, *dirPath)

	savedAttachment, err := a.attachmentRepository.Save(attachment)
	if err != nil {
		return nil, fmt.Errorf("error saving attachment: %w", err)
	}

	attachmentDTO := models.NewAttachmentDTO(*savedAttachment)
	return &attachmentDTO, nil
}

func NewAttachmentService(atRepo repositories.AttachmentRepository, deps AppDependencies) AttachmentService {
	return &attachmentService{atRepo, deps}
}
