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
)

type AttachmentService interface {
	Save(uint64, *models.Claims, *multipart.File, *multipart.FileHeader) (*models.AttachmentDTO, error)
	// private
	createDirIfNotExists() (*string, error)
	getMimeType(string) string
	writeToFile(string, []byte) error
	generateUniqueFileName() string
	createAttachment(*models.Claims, *multipart.FileHeader, uint64, string, string, string) *models.Attachment
	generateURL(string, string) string
}

type attachmentService struct {
	attachmentRepository repositories.AttachmentRepository
	deps                 AppDependencies
}

func (a *attachmentService) createAttachment(
	user *models.Claims,
	fileHeader *multipart.FileHeader,
	troubleTicketId uint64,
	uniqueFileName, fileExt, dirPath string,
) *models.Attachment {

	size := fileHeader.Size
	ref := a.generateUniqueFileName()
	mimeType := a.getMimeType(fileExt)
	url := a.generateURL(uniqueFileName, fileExt)

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

func (a *attachmentService) generateURL(uniqueFileName, fileExt string) string {
	return a.deps.AppConfig.AttachmentHost + "/static/attachment/file/" + uniqueFileName + fileExt
}

func (a *attachmentService) getMimeType(fileExtension string) string {
	mimeType := mime.TypeByExtension(fileExtension)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	return mimeType
}

func (a *attachmentService) generateUniqueFileName() string {
	// Combine UUID and timestamp to create a unique filename
	return fmt.Sprintf("%d-%s", time.Now().Unix(), uuid.New().String())
}

func (a *attachmentService) writeToFile(filePath string, content []byte) error {
	err := os.WriteFile(filePath, content, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %w", err)
	}
	return nil
}

// createDirIfNotExists checks app root dir if data dir exists, creates if not & returns its path
func (a *attachmentService) createDirIfNotExists() (*string, error) {
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
	dirPath, err := a.createDirIfNotExists()
	if err != nil {
		return nil, err
	}

	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return nil, fmt.Errorf("error reading file content: %w", err)
	}

	fileName := fileHeader.Filename
	uniqueFileName := a.generateUniqueFileName()
	fileExt := filepath.Ext(fileName)
	filePath := filepath.Join(*dirPath, uniqueFileName+fileExt)

	err = a.writeToFile(filePath, fileBytes)
	if err != nil {
		return nil, err
	}

	attachment := a.createAttachment(user, fileHeader, troubleTicketID, uniqueFileName, fileExt, *dirPath)

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
