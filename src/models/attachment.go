package models

import "time"

// Attachment File(s) attached to the trouble ticket. e.g. picture of broken device, scanning of a bill or charge etc.
type Attachment struct {
	BaseModel
	Ref             string `json:"ref"`
	Type            string `json:"type"`
	MimeType        string `json:"mime_type"`
	OriginalName    string `json:"original_name"`
	Path            string `json:"path"`
	Size            uint64 `json:"size"`
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	TroubleTicketID uint64 `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}

// AttachmentDTO a data transfer object for an attachment.
type AttachmentDTO struct {
	Ref          string    `json:"ref"`
	Type         string    `json:"type"`
	MimeType     string    `json:"mime_type"`
	OriginalName string    `json:"original_name"`
	Path         string    `json:"path"`
	Size         uint64    `json:"size"`
	Name         string    `json:"name"`
	URL          string    `json:"url"`
	Description  string    `json:"description"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
}

// NewAttachmentDTO creates a new AttachmentDTO instance from an Attachment entity.
func NewAttachmentDTO(attachment Attachment) AttachmentDTO {
	return AttachmentDTO{
		Ref:          attachment.Ref,
		Type:         attachment.Type,
		MimeType:     attachment.MimeType,
		OriginalName: attachment.OriginalName,
		Path:         attachment.Path,
		Size:         attachment.Size,
		Name:         attachment.Name,
		URL:          attachment.URL,
		Description:  attachment.Description,
		CreatedBy:    attachment.CreatedBy,
		CreatedAt:    attachment.CreatedAt,
	}
}
