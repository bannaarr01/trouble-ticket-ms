package models

// Note The note(s) that are associated to the ticket.
type Note struct {
	BaseModel
	Author          string `json:"author"`
	Text            string `gorm:"type:text;not null" json:"text"`
	TroubleTicketID uint64 `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}
