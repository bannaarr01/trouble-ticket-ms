package models

// StatusChange The status change history that are associated to the ticket
type StatusChange struct {
	BaseModel
	Reason          string `json:"reason"`
	StatusID        uint64 `gorm:"index;not null" json:"status_id"`
	Status          Status `gorm:"foreignKey:StatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TroubleTicketID uint64 `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}
