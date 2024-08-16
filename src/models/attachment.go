package models

// Attachment File(s) attached to the trouble ticket. e.g. picture of broken device, scanning of a bill or charge etc.
type Attachment struct {
	BaseModel
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
