package models

import "time"

// Note The note(s) that are associated to the ticket.
type Note struct {
	BaseModel
	Author          string `json:"author"`
	Text            string `gorm:"type:text;not null" json:"text"`
	TroubleTicketID uint64 `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}

type NoteDTO struct {
	ID        uint64    `json:"id"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNoteDTO(note Note) NoteDTO {
	return NoteDTO{
		ID:        note.ID,
		Author:    note.Author,
		Text:      note.Text,
		CreatedAt: note.CreatedAt,
	}
}
